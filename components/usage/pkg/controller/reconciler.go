// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/usage/pkg/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reconciler interface {
	Reconcile() error
}

type ReconcilerFunc func() error

func (f ReconcilerFunc) Reconcile() error {
	return f()
}

type UsageReconciler struct {
	nowFunc           func() time.Time
	conn              *gorm.DB
	pricer            *WorkspacePricer
	billingController BillingController
}

func NewUsageReconciler(conn *gorm.DB, pricer *WorkspacePricer, billingController BillingController) *UsageReconciler {
	return &UsageReconciler{
		conn:              conn,
		pricer:            pricer,
		billingController: billingController,
		nowFunc:           time.Now,
	}
}

type UsageReconcileStatus struct {
	StartTime time.Time
	EndTime   time.Time

	WorkspaceInstances        int
	InvalidWorkspaceInstances int
}

func (u *UsageReconciler) Reconcile() (err error) {
	ctx := context.Background()
	now := time.Now().UTC()

	reportUsageReconcileStarted()
	defer func() {
		reportUsageReconcileFinished(time.Since(now), err)
	}()

	startOfCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	startOfNextMonth := startOfCurrentMonth.AddDate(0, 1, 0)

	status, report, err := u.ReconcileTimeRange(ctx, startOfCurrentMonth, startOfNextMonth)
	if err != nil {
		return err
	}
	log.WithField("usage_reconcile_status", status).Info("Reconcile completed.")

	// For now, write the report to a temp directory so we can manually retrieve it
	dir := os.TempDir()
	f, err := ioutil.TempFile(dir, fmt.Sprintf("%s-*", now.Format(time.RFC3339)))
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	err = enc.Encode(report)
	if err != nil {
		return fmt.Errorf("failed to marshal report to JSON: %w", err)
	}

	stat, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %w", err)
	}
	log.Infof("Wrote usage report into %s", filepath.Join(dir, stat.Name()))

	err = db.CreateUsageRecords(ctx, u.conn, usageReportToUsageRecords(report, u.pricer, u.nowFunc().UTC()))
	if err != nil {
		return fmt.Errorf("failed to write usage records to database: %s", err)
	}

	return nil
}

func (u *UsageReconciler) ReconcileTimeRange(ctx context.Context, from, to time.Time) (*UsageReconcileStatus, UsageReport, error) {
	now := u.nowFunc().UTC()
	log.Infof("Gathering usage data from %s to %s", from, to)
	status := &UsageReconcileStatus{
		StartTime: from,
		EndTime:   to,
	}
	instances, invalidInstances, err := u.loadWorkspaceInstances(ctx, from, to)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load workspace instances: %w", err)
	}
	status.WorkspaceInstances = len(instances)
	status.InvalidWorkspaceInstances = len(invalidInstances)

	if len(invalidInstances) > 0 {
		log.WithField("invalid_workspace_instances", invalidInstances).Errorf("Detected %d invalid instances. These will be skipped in the current run.", len(invalidInstances))
	}
	log.WithField("workspace_instances", instances).Debug("Successfully loaded workspace instances.")

	instancesByAttributionID := groupInstancesByAttributionID(instances)

	err = u.billingController.Reconcile(ctx, now, instancesByAttributionID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to reconcile billing: %w", err)
	}

	return status, instancesByAttributionID, nil
}

type UsageReport map[db.AttributionID][]db.WorkspaceInstanceForUsage

func (u UsageReport) CreditSummaryForTeams(pricer *WorkspacePricer, maxStopTime time.Time) map[string]int64 {
	creditsPerTeamID := map[string]int64{}

	for attribution, instances := range u {
		entity, id := attribution.Values()
		if entity != db.AttributionEntity_Team {
			continue
		}

		var credits float64
		for _, instance := range instances {
			credits += pricer.CreditsUsedByInstance(&instance, maxStopTime)
		}

		creditsPerTeamID[id] = int64(math.Ceil(credits))
	}

	return creditsPerTeamID
}

type invalidWorkspaceInstance struct {
	reason              string
	workspaceInstanceID uuid.UUID
}

func (u *UsageReconciler) loadWorkspaceInstances(ctx context.Context, from, to time.Time) ([]db.WorkspaceInstanceForUsage, []invalidWorkspaceInstance, error) {
	log.Infof("Gathering usage data from %s to %s", from, to)
	instances, err := db.ListWorkspaceInstancesInRange(ctx, u.conn, from, to)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list instances from db: %w", err)
	}
	log.Infof("Identified %d instances between %s and %s", len(instances), from, to)

	valid, invalid := validateInstances(instances)
	trimmed := trimStartStopTime(valid, from, to)
	return trimmed, invalid, nil
}

func validateInstances(instances []db.WorkspaceInstanceForUsage) (valid []db.WorkspaceInstanceForUsage, invalid []invalidWorkspaceInstance) {
	for _, i := range instances {
		// i is a pointer to the current element, we need to assign it to ensure we're copying the value, not the current pointer.
		instance := i

		// Each instance must have a start time, without it, we do not have a baseline for usage computation.
		if !instance.CreationTime.IsSet() {
			invalid = append(invalid, invalidWorkspaceInstance{
				reason:              "missing creation time",
				workspaceInstanceID: instance.ID,
			})
			continue
		}

		start := instance.CreationTime.Time()

		// Currently running instances do not have a stopped time set, so we ignore these.
		if instance.StoppedTime.IsSet() {
			stop := instance.StoppedTime.Time()
			if stop.Before(start) {
				invalid = append(invalid, invalidWorkspaceInstance{
					reason:              "stop time is before start time",
					workspaceInstanceID: instance.ID,
				})
				continue
			}
		}

		valid = append(valid, instance)
	}
	return valid, invalid
}

// trimStartStopTime ensures that start time or stop time of an instance is never outside of specified start or stop time range.
func trimStartStopTime(instances []db.WorkspaceInstanceForUsage, maximumStart, minimumStop time.Time) []db.WorkspaceInstanceForUsage {
	var updated []db.WorkspaceInstanceForUsage

	for _, instance := range instances {
		if instance.CreationTime.Time().Before(maximumStart) {
			instance.CreationTime = db.NewVarcharTime(maximumStart)
		}

		if instance.StoppedTime.Time().After(minimumStop) {
			instance.StoppedTime = db.NewVarcharTime(minimumStop)
		}

		updated = append(updated, instance)
	}
	return updated
}

func groupInstancesByAttributionID(instances []db.WorkspaceInstanceForUsage) UsageReport {
	result := map[db.AttributionID][]db.WorkspaceInstanceForUsage{}
	for _, instance := range instances {
		if _, ok := result[instance.UsageAttributionID]; !ok {
			result[instance.UsageAttributionID] = []db.WorkspaceInstanceForUsage{}
		}

		result[instance.UsageAttributionID] = append(result[instance.UsageAttributionID], instance)
	}

	return result
}

func usageReportToUsageRecords(report UsageReport, pricer *WorkspacePricer, now time.Time) []db.WorkspaceInstanceUsage {
	var usageRecords []db.WorkspaceInstanceUsage

	for attributionId, instances := range report {
		for _, instance := range instances {
			var stoppedAt sql.NullTime
			if instance.StoppedTime.IsSet() {
				stoppedAt = sql.NullTime{Time: instance.StoppedTime.Time(), Valid: true}
			}

			projectID := ""
			if instance.ProjectID.Valid {
				projectID = instance.ProjectID.String
			}

			usageRecords = append(usageRecords, db.WorkspaceInstanceUsage{
				InstanceID:     instance.ID,
				AttributionID:  attributionId,
				WorkspaceID:    instance.WorkspaceID,
				ProjectID:      projectID,
				UserID:         instance.OwnerID,
				WorkspaceType:  instance.Type,
				WorkspaceClass: instance.WorkspaceClass,
				StartedAt:      instance.CreationTime.Time(),
				StoppedAt:      stoppedAt,
				CreditsUsed:    pricer.CreditsUsedByInstance(&instance, now),
				GenerationID:   0,
			})
		}
	}

	return usageRecords
}

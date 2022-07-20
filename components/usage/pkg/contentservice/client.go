// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package contentservice

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/content-service/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Interface interface {
	UploadFile(ctx context.Context, filePath string) error
}

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{url: url}
}

func (c *Client) UploadFile(ctx context.Context, filePath string) error {
	url, err := c.getSignedUploadUrl(ctx, filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("failed to obtain signed upload URL: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, file)
	if err != nil {
		return fmt.Errorf("failed to construct http request: %w", err)
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	log.Infof("Uploading %q to cloud storage...", filePath)
	req.ContentLength = fileInfo.Size()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make http request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http response code: %s", resp.Status)
	}
	log.Info("Upload complete")

	return nil
}

func (c *Client) getSignedUploadUrl(ctx context.Context, key string) (string, error) {
	conn, err := grpc.Dial(c.url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("failed to dial content-service gRPC server: %w", err)
	}
	defer conn.Close()

	uc := api.NewUsageReportServiceClient(conn)

	resp, err := uc.UploadURL(ctx, &api.UsageReportUploadURLRequest{Name: key})
	if err != nil {
		return "", fmt.Errorf("failed RPC to content service: %w", err)
	}

	return resp.Url, nil
}

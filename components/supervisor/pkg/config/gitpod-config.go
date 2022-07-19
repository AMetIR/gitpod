// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package config

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/gitpod-io/gitpod/common-go/log"
	gitpod "github.com/gitpod-io/gitpod/gitpod-protocol"

	"github.com/gitpod-io/gitpod/supervisor/api"
)

// ConfigInterface provides access to the gitpod config file.
type ConfigInterface interface {
	// Watch starts the config watching
	Watch(ctx context.Context)
	// Observe provides channels triggered whenever the config is changed
	Observe(ctx context.Context) <-chan *gitpod.GitpodConfig
}

// ConfigService provides access to the gitpod config file.
type ConfigService struct {
	location      string
	locationReady <-chan struct{}

	cond   *sync.Cond
	config *gitpod.GitpodConfig

	pollTimer *time.Timer

	log *logrus.Entry
}

// NewConfigService creates a new instance of ConfigService.
func NewConfigService(configLocation string, locationReady <-chan struct{}, log *logrus.Entry) *ConfigService {
	return &ConfigService{
		location:      configLocation,
		locationReady: locationReady,
		cond:          sync.NewCond(&sync.Mutex{}),
		log:           log.WithField("location", configLocation),
	}
}

// Observe provides channels triggered whenever the config is changed.
func (service *ConfigService) Observe(ctx context.Context) <-chan *gitpod.GitpodConfig {
	configs := make(chan *gitpod.GitpodConfig)
	go func() {
		defer close(configs)

		service.cond.L.Lock()
		defer service.cond.L.Unlock()
		for {
			configs <- service.config

			service.cond.Wait()
			if ctx.Err() != nil {
				return
			}
		}
	}()
	return configs
}

// Watch starts the config watching.
func (service *ConfigService) Watch(ctx context.Context) {
	service.log.Info("gitpod config watcher: starting...")

	select {
	case <-service.locationReady:
	case <-ctx.Done():
		return
	}

	_, err := os.Stat(service.location)
	if os.IsNotExist(err) {
		service.poll(ctx)
	}
	service.watch(ctx)
}

func (service *ConfigService) watch(ctx context.Context) {
	watcher, err := fsnotify.NewWatcher()
	defer func() {
		if err != nil {
			service.log.WithError(err).Error("gitpod config watcher: failed to start")
			return
		}

		service.log.Info("gitpod config watcher: started")
	}()
	if err != nil {
		return
	}

	err = watcher.Add(service.location)
	if err != nil {
		watcher.Close()
		return
	}

	go func() {
		defer service.log.Info("gitpod config watcher: stopped")
		defer watcher.Close()

		polling := make(chan struct{}, 1)
		service.scheduleUpdateConfig(ctx, polling)
		for {
			select {
			case <-polling:
				return
			case <-ctx.Done():
				return
			case err := <-watcher.Errors:
				service.log.WithError(err).Error("gitpod config watcher: failed to watch")
			case <-watcher.Events:
				service.scheduleUpdateConfig(ctx, polling)
			}
		}
	}()
}

func (service *ConfigService) scheduleUpdateConfig(ctx context.Context, polling chan<- struct{}) {
	service.cond.L.Lock()
	defer service.cond.L.Unlock()
	if service.pollTimer != nil {
		service.pollTimer.Stop()
	}
	service.pollTimer = time.AfterFunc(100*time.Millisecond, func() {
		err := service.updateConfig()
		if os.IsNotExist(err) {
			polling <- struct{}{}
			go service.poll(ctx)
		} else if err != nil {
			service.log.WithError(err).Error("gitpod config watcher: failed to parse")
		}
	})
}

func (service *ConfigService) poll(ctx context.Context) {
	timer := time.NewTicker(2 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
		}

		if _, err := os.Stat(service.location); !os.IsNotExist(err) {
			service.watch(ctx)
			return
		}
	}
}

func runAsGitpodUser(cmd *exec.Cmd) *exec.Cmd {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	if cmd.SysProcAttr.Credential == nil {
		cmd.SysProcAttr.Credential = &syscall.Credential{}
	}
	cmd.SysProcAttr.Credential.Uid = 33333
	cmd.SysProcAttr.Credential.Gid = 33333
	return cmd
}

type subscription struct {
	// id      uint64
	// channel chan *api.SubscribeResponse
	// once    sync.Once
	// closed  bool
	// cancel  context.CancelFunc
}

type pendingNotification struct {
	// message *api.SubscribeResponse
	// responseChannel chan *api.NotifyResponse
	// once            sync.Once
	// closed          bool
}

type NotificationService struct {
	// mutex                sync.Mutex
	nextSubscriptionID   uint64
	subscriptions        map[uint64]*subscription
	nextNotificationID   uint64
	pendingNotifications map[uint64]*pendingNotification

	api.UnimplementedNotificationServiceServer
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		subscriptions:        make(map[uint64]*subscription),
		pendingNotifications: make(map[uint64]*pendingNotification),
	}
}

func (service *ConfigService) updateConfig() error {
	service.cond.L.Lock()
	defer service.cond.L.Unlock()

	config, err := service.parse()
	if err != nil {
		panic(err)
	}
	service.config = config
	service.cond.Broadcast()

	// send notification

	ctx := context.Background()
	notificationService := NewNotificationService()

	log.Info(notificationService.pendingNotifications)
	log.Info(notificationService.subscriptions)
	fmt.Println(notificationService.nextNotificationID)
	fmt.Println(notificationService.nextSubscriptionID)

	message := "The .gitpod.yml configuration file changed. Would you like to rebuild the workspace to test it?"
	result, err := notificationService.Notify(ctx, &api.NotifyRequest{
		Level:   api.NotifyRequest_INFO,
		Message: message,
		Actions: []string{"Rebuild Now"},
	})
	if err != nil {
		log.WithError(err).Fatal("failed to send test notification")
	}

	if result.Action == "Rebuild Now" {
		gpPath, err := exec.LookPath("gp")
		if err != nil {
			log.WithError(err).Fatal("failed to find gp cli")
		}

		repoUrl := "https://github.com/akosyakov/parcel-demo"
		additionalContent := "eyJpbmRleC5qcyI6ImxhbGFsYSJ9"
		rebuildUrl := fmt.Sprintf("%s/#additionalcontent/%s/%s", os.Getenv("GITPOD_WORKSPACE_URL"), additionalContent, repoUrl)

		// type AddititionalContent struct {
		// 	gitpodYml string
		// }

		// var ad AddititionalContent

		gpCmd := exec.Command(gpPath, "preview", "--external", rebuildUrl)
		gpCmd = runAsGitpodUser(gpCmd)
		err = gpCmd.Start()
		if err != nil {
			log.WithError(err).Fatal("todo: handle error")
		}
		err = gpCmd.Process.Release()
		if err != nil {
			log.WithError(err).Fatal("todo: handle error")
		}
	}

	service.log.WithField("config", service.config).Warn("ANDREA: gitpod config watcher: updated")

	service.log.WithField("config", service.config).Debug("gitpod config watcher: updated")

	return err
}

func (service *ConfigService) parse() (*gitpod.GitpodConfig, error) {
	data, err := os.ReadFile(service.location)
	if err != nil {
		return nil, err
	}
	var config *gitpod.GitpodConfig
	err = yaml.Unmarshal(data, &config)
	return config, err
}

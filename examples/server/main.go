package main

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"gitlab.com/gobl/gobl/pkg/bootstrap"
	"gitlab.com/gobl/gobl/pkg/service"

	"gitlab.com/gobl/gobl/pkg/cmd"
	"gitlab.com/gobl/gobl/pkg/logger"
)

type serverState struct {
	service.State
	notifyCh chan struct{}
	cancel   context.CancelFunc
}

var (
	st                    *serverState
	ErrInvalidServerState = errors.New("state is not a server state object")
)

func initApp(ctx context.Context, store service.State) error {
	log := logger.Get(logger.ApplicationLogLevel(), logger.ConfiguredLumberjackLogger())

	ss, ok := store.(*serverState)
	if !ok {
		return errors.New("store is not a server state object")
	}

	log.Info("initialising application")

	childCtx, cancel := context.WithCancel(ctx)
	notifyCh := make(chan struct{})

	timeout := 5 * time.Second

	go func(c context.Context, notifyCh chan struct{}) {
		log.Info("Starting long running process...")

		count := 0

		// long running process here
		for {
			select {
			case <-c.Done():
				log.Info("Stopping long running process")
				notifyCh <- struct{}{}

				return
			case <-time.After(timeout):
				count++
				log.Info("Current count", zap.Int("count", count))
			}
		}
	}(childCtx, notifyCh)

	ss.cancel = cancel
	ss.notifyCh = notifyCh

	return nil
}

func cleanupApp(state service.State) error {
	ss, ok := state.(*serverState)
	if !ok {
		return ErrInvalidServerState
	}

	ss.cancel()

	// we wait for notification that our long running process has cleanly exited
	<-ss.notifyCh

	return nil
}

type Service struct {
	service.Service
}

func NewService() *Service {
	return &Service{
		Service: bootstrap.New(),
	}
}

func main() {
	app := NewService().
		AddInitFunc(initApp).
		AddCleanupFunc(cleanupApp)

	cmd.SetCliProperties("usage", "a short description", "full help text for the cli")

	st = &serverState{}

	cmd.Execute(context.Background(), app, st)
}

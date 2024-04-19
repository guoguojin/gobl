package main

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"gitlab.com/gobl/gobl/pkg/bootstrap"

	"gitlab.com/gobl/gobl/examples/turnstile/events"
	"gitlab.com/gobl/gobl/examples/turnstile/states"
	"gitlab.com/gobl/gobl/examples/turnstile/transitions"
	"gitlab.com/gobl/gobl/pkg/cmd"
	"gitlab.com/gobl/gobl/pkg/fsm"
	"gitlab.com/gobl/gobl/pkg/logger"
	"gitlab.com/gobl/gobl/pkg/service"
)

const (
	machineName = "Turnstile Service"
)

type TurnstileState struct {
	service.State
	machine            fsm.FSM
	errCh              chan error
	incoming           chan fsm.Event
	cancelMachine      context.CancelFunc
	cancelErrorHandler context.CancelFunc
}

func initStateMachine(ctx context.Context, state service.State) error {
	s, ok := state.(*TurnstileState)
	if !ok {
		return errors.New("unexpected state")
	}
	s.incoming = make(chan fsm.Event)
	s.machine, s.errCh = fsm.New(uuid.New(), machineName, states.Locked(uuid.New()).
		WithTransitions(
			fsm.Transition{
				Checks: []fsm.CheckFn{transitions.HasCoin},
				Next:   transitions.ToUnlocked,
			},
		))

	machineCtx, cancel := context.WithCancel(ctx)

	s.cancelMachine = cancel
	cleanupFunc := func() error { return nil }

	go func(c context.Context, events <-chan fsm.Event) {
		logger.Logger().Info("Starting Turnstile State Machine")
		if err := s.machine.Run(c, events, cleanupFunc); err != nil {
			logger.Logger().Error("state machine error", zap.Error(err))
			s.cancelMachine()
		}
	}(machineCtx, s.incoming)

	return nil
}

func initErrorHandler(ctx context.Context, state service.State) error {
	s, ok := state.(*TurnstileState)
	if !ok {
		return errors.New("unexpected state")
	}

	errCtx, cancel := context.WithCancel(ctx)

	s.cancelErrorHandler = cancel

	go func(c context.Context, errCh <-chan error) {
		for {
			select {
			case <-c.Done():
				logger.Logger().Warn("Stopping turnstile service error handler")
				return
			case err := <-errCh:
				logger.Logger().Error(err.Error())
			}
		}
	}(errCtx, s.errCh)

	return nil
}

func runTurnstile(_ context.Context, state service.State) error {
	s, ok := state.(*TurnstileState)
	if !ok {
		return errors.New("unexpected state")
	}
	s.incoming <- events.Push(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.InsertCoin(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.InsertCoin(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.Push(uuid.New(), "TEST", time.Now().UnixNano())
	s.incoming <- events.Push(uuid.New(), "TEST", time.Now().UnixNano())
	return nil
}

func cleanupStateMachine(state service.State) error {
	s, ok := state.(*TurnstileState)
	if !ok {
		return errors.New("unexpected state")
	}
	s.cancelMachine()
	return nil
}

func cleanupErrorHandler(state service.State) error {
	s, ok := state.(*TurnstileState)
	if !ok {
		return errors.New("unexpected state")
	}
	s.cancelErrorHandler()
	return nil
}

func main() {
	app := bootstrap.New().
		AddInitFunc(initStateMachine, initErrorHandler).
		AddCleanupFunc(cleanupStateMachine, cleanupErrorHandler).
		WithRunFunc(runTurnstile)

	st := new(TurnstileState)

	cmd.Execute(context.Background(), app, st)
}

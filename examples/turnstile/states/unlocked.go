package states

import (
	"errors"

	"github.com/google/uuid"

	"gitlab.com/gobl/gobl/examples/turnstile/events"
	"gitlab.com/gobl/gobl/pkg/fsm"
)

type UnlockedState struct {
	id          uuid.UUID
	Pushed      bool
	transitions []fsm.Transition
}

func Unlocked(id uuid.UUID) *UnlockedState {
	unlocked := UnlockedState{
		id:     id,
		Pushed: false,
	}

	return &unlocked
}

func (u *UnlockedState) ID() uuid.UUID {
	return u.id
}

func (u *UnlockedState) Description() string {
	return "Turnstile Unlocked"
}

func (u *UnlockedState) Execute(evt fsm.Event) error {
	switch e := evt.(type) {
	case events.PushEvent:
		u.Pushed = true
	case events.InsertCoinEvent:
		return errors.New("coin already inserted, returning coin")
	default:
		return fsm.ErrUnexpectedEvent(e)
	}
	return nil
}

func (u *UnlockedState) Next() fsm.State {
	return fsm.Next(u, u.transitions...)
}

func (u *UnlockedState) WithTransitions(transitions ...fsm.Transition) fsm.State {
	u.transitions = append(u.transitions, transitions...)
	return u
}

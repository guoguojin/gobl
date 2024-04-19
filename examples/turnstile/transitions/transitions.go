package transitions

import (
	"github.com/google/uuid"

	"gitlab.com/gobl/gobl/examples/turnstile/states"
	"gitlab.com/gobl/gobl/pkg/fsm"
)

func HasCoin(state fsm.State) bool {
	switch s := state.(type) {
	case *states.LockedState:
		return s.HasCoin
	default:
		return false
	}
}

func Pushed(state fsm.State) bool {
	switch s := state.(type) {
	case *states.UnlockedState:
		return s.Pushed
	default:
		return false
	}
}

func ToLocked(_ fsm.State) fsm.State {
	return states.Locked(uuid.New()).WithTransitions(
		fsm.Transition{
			Checks: []fsm.CheckFn{HasCoin},
			Next:   ToUnlocked,
		},
	)
}

func ToUnlocked(_ fsm.State) fsm.State {
	return states.Unlocked(uuid.New()).WithTransitions(
		fsm.Transition{
			Checks: []fsm.CheckFn{Pushed},
			Next:   ToLocked,
		},
	)
}

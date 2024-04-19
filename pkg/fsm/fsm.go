package fsm

import (
	"fmt"
	"time"
)

const (
	// TimestampFormat is the default format for events that are passing through the state machine (RFC3339Nano)
	TimestampFormat = time.RFC3339Nano
)

// TimestampToTime takes the timestamp represented as nanoseconds past epoch and converts it into a
// time.Time struct
func TimestampToTime(t int64) time.Time {
	return time.Unix(0, t)
}

// TimestampToString take the timestamp represented as nanoseconds past epoch and converts it to the
// a string using the Timestamp format
func TimestampToString(t int64) string {
	return TimestampToTime(t).Format(TimestampFormat)
}

// ErrUnexpectedEvent returns an error describing the event received
func ErrUnexpectedEvent(event Event) error {
	return fmt.Errorf("an unexpected event %s was received from %s", event.Name(), event.Source())
}

func ErrUnexpectedState(state State) error {
	return fmt.Errorf("an unexpected state was received: %s", state.Description())
}

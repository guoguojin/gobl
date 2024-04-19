package notification

import (
	"gitlab.com/gobl/gobl/pkg/property"
	"gitlab.com/gobl/gobl/pkg/tags"
	"time"
)

// Notifier is an interface that defines a method for sending notifications.
// This interface should be implemented by any type that can send notifications.
type Notifier interface {
	Notify(Notification) error
}

type Code int32

type Level uint8

const (
	None Level = iota
	Debug
	Info
	Warn
	Error

	none = "NONE"
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	default:
		return none
	}
}

type Notification interface {
	// Code returns a code representing the type of notification being sent.
	// It is up to the user to define the codes and their meanings.
	Code() Code
	// Level returns the severity level of the notification.
	Level() Level
	// Message returns information regarding why the notification is being sent.
	Message() string
	// Error returns an error associated with the notification, if any.
	Error() error
	// Timestamp returns the time the notification was created.
	Timestamp() time.Time
	// Subsystem returns the name of the subsystem that generated the notification.
	Subsystem() string
	// TracerID returns that can be used to filter for an event as it passes through the system.
	TracerID() string
	// Tags returns a set of tags that can be used to filter for an event as it passes through the system.
	Tags() tags.Tags
	// Properties returns a set of properties associated with the notification.
	Properties() property.Properties
}

// Handler is an interface that defines a method for handling notifications.
// This interface should be implemented by any type that can handle notifications.
type Handler interface {
	Handle(Notification) error
}

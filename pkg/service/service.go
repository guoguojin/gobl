package service

import (
	"context"
	"math"

	"gitlab.com/gobl/gobl/pkg/property"
)

// Service is the interface you need to implement for creating a service with Gobl
// The application package is a concrete implementation of this interface.
type Service interface {
	// Init executes the initialisation functions that were provided to the service in the order they were provided
	Init(ctx context.Context, state State) error
	// AddInitFunc adds the given initialisation functions to the service
	AddInitFunc(fns ...InitFunc) Service
	// Cleanup executes the cleanup functions that were provided to the service in the order they were provided
	Cleanup(state State) error
	// AddCleanupFunc adds the given cleanup functions to the service
	AddCleanupFunc(fns ...CleanupFunc) Service
	// WithRunFunc passes an alternative run function to the service
	WithRunFunc(fn RunFunc) Service
	// RunFunction returns the current main function to run for the service
	RunFunction() RunFunc
	// SetProperties sets the properties for the service
	SetProperties(properties property.Properties) error
	// AddProperty adds a property to the service
	AddProperty(name string, p property.Property) Service
	// GetProperty returns the requested property
	GetProperty(name string) (property.Property, error)
}

// InitFunc is a function that can be called to perform an initialisation task for a service
type InitFunc func(ctx context.Context, state State) error

// CleanupFunc is a function that can be called to perform a cleanup task for a service on graceful exit
type CleanupFunc func(state State) error

// RunFunc is a function that can be run in place of the long running service
type RunFunc func(ctx context.Context, state State) error

// MaxPort returns the highest port number that a service can run on
func MaxPort() int {
	return math.MaxUint16
}

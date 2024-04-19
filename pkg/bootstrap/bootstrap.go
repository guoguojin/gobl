package bootstrap

import (
	"context"

	"gitlab.com/gobl/gobl/pkg/cmd"
	"gitlab.com/gobl/gobl/pkg/property"
	"gitlab.com/gobl/gobl/pkg/service"
)

type bootstrap struct {
	initFunctions    []service.InitFunc
	cleanupFunctions []service.CleanupFunc
	runFunc          service.RunFunc
	properties       property.Properties
}

// New creates a new instance of the bootstrap server application
// The bootstrap server application is a concrete implementation of the
// service.Service interface
func New() service.Service {
	return &bootstrap{
		initFunctions:    make([]service.InitFunc, 0),
		cleanupFunctions: make([]service.CleanupFunc, 0),
		runFunc:          nil,
		properties:       make(property.Properties),
	}
}

// Init is called when the application starts and executes the initialisation functions
// that have been added to the application
func (a *bootstrap) Init(ctx context.Context, state service.State) error {
	for _, f := range a.initFunctions {
		if err := f(ctx, state); err != nil {
			return err
		}
	}

	return nil
}

// AddInitFunc adds initialisation function that are needed to initialise the application
func (a *bootstrap) AddInitFunc(initFuncs ...service.InitFunc) service.Service {
	a.initFunctions = append(a.initFunctions, initFuncs...)
	return a
}

// Cleanup is called when the application terminates. It executes all cleanup functions
// that have been added to the application attempting to ensure the application can
// terminate gracefully.
func (a *bootstrap) Cleanup(state service.State) error {
	for _, f := range a.cleanupFunctions {
		if err := f(state); err != nil {
			return err
		}
	}

	return nil
}

// AddCleanupFunc adds cleanup function that will be run when the application attempts a graceful shutdown
func (a *bootstrap) AddCleanupFunc(fns ...service.CleanupFunc) service.Service {
	a.cleanupFunctions = append(a.cleanupFunctions, fns...)
	return a
}

// SetProperties adds properties that may be needed by the application
func (a *bootstrap) SetProperties(properties property.Properties) error {
	a.properties = properties

	usage := a.properties.GetOrElse(property.KeyUsage, property.Empty()).NonEmptyString(cmd.Usage)
	shortDesc := a.properties.GetOrElse(property.KeyShortDesc, property.Empty()).NonEmptyString(cmd.ShortDescription)
	longDesc := a.properties.GetOrElse(property.KeyLongDesc, property.Empty()).NonEmptyString(cmd.LongDescription)

	cmd.SetCliProperties(usage, shortDesc, longDesc)

	return nil
}

// WithRunFunc allows you to set the function that should be executed instead of starting the server
// application when the application starts.
func (a *bootstrap) WithRunFunc(f service.RunFunc) service.Service {
	a.runFunc = f
	return a
}

// RunFunction is the function that will be executed when the application starts instead of the
// main server process.
func (a *bootstrap) RunFunction() service.RunFunc {
	return a.runFunc
}

// AddProperty allows you to add a property that may be needed by the server
func (a *bootstrap) AddProperty(name string, p property.Property) service.Service {
	a.properties.Set(name, p)
	return a
}

// GetProperty allows you to retrieve a property that has been set for the server
func (a *bootstrap) GetProperty(name string) (property.Property, error) {
	p, ok := a.properties.Get(name)
	if !ok {
		return property.Empty(), property.ErrDoesNotExist
	}
	return p, nil
}

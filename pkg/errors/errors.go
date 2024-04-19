package errors

import (
	ge "errors"
)

// ErrNotImplemented is a useful error to return when you are starting out and have
// not implemented your functions or logic yet
var ErrNotImplemented = ge.New("NotImplemented")

// ErrNoServiceInitialisers is returned when you have no service initialised
var ErrNoServiceInitialisers = ge.New("no service initialisers have been defined")

// ErrServicePropertyNotDefined is returned when you have no service property defined for a
// given key
var ErrServicePropertyNotDefined = ge.New("service property has not been defined")

package testing

import "errors"

var (
	ErrInvalidEnvironmentVariable = errors.New("invalid environment variable, expected KEY=value")
)

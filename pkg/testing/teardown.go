package testing

import "testing"

type TeardownFunc func(*testing.T) error

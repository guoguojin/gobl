package utils

import (
	"strings"
)

// IsTrue takes a string representing a boolean value and converts it to a boolean, if the string is true, or yes (regardless of case)
// it will be evaluated as true, anything else will be regarded as false
func IsTrue(value string) bool {
	if len(value) < 1 || len(value) > 5 {
		return false
	} // false is the longest

	switch strings.ToUpper(value)[0:1] {
	case "T", "Y", "1":
		return true
	default:
		return false
	}
}

// Ptr returns a pointer to the given value
func Ptr[T any](value T) *T {
	return &value
}

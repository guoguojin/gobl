package property

import "errors"

const (
	// KeyUsage is the property key for the usage property of an application
	KeyUsage = "usage"
	// KeyShortDesc is the property key for the short description property of an application
	KeyShortDesc = "shortDesc"
	// KeyLongDesc is the property key for the long description property of an application
	KeyLongDesc = "longDesc"
)

// ErrDoesNotExist is an error that is returned when a property does not exist
var ErrDoesNotExist = errors.New("property does not exist")

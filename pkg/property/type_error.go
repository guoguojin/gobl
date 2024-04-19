package property

import (
	"fmt"
)

type TypeError struct {
	propertyType Type
	msg          string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("%s - %s", e.propertyType, e.msg)
}

func NewTypeError(propertyType Type, msg string) TypeError {
	return TypeError{
		propertyType: propertyType,
		msg:          msg,
	}
}

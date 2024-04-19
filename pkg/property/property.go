package property

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const (
	dtFormat             = time.RFC3339Nano
	dps                  = 6
	emptyArrayString     = "[]"
	trueString           = "true"
	falseString          = "false"
	floatZeroString      = "0.000000"
	emptyInterfaceString = "{}"
)

// Property of an order holds information that is relevant
type Property struct {
	Name  string      `json:"name"`
	Type  Type        `json:"type"`
	Value interface{} `json:"value"`
}

// Empty creates an empty property whose type is undefined and value is nil
func Empty() Property {
	return Property{
		Type:  Undefined,
		Value: nil,
	}
}

// StringProperty creates a Property that contains unicode text based information
func StringProperty(name, value string) Property {
	return Property{
		Name:  name,
		Type:  String,
		Value: value,
	}
}

// StringArrayProperty creates a Property that contains an array of unicode text based information
func StringArrayProperty(name string, value ...string) Property {
	return Property{
		Name:  name,
		Type:  StringArray,
		Value: value,
	}
}

// BoolProperty creates a Property that contains a boolean value
func BoolProperty(name string, value bool) Property {
	return Property{
		Name:  name,
		Type:  Bool,
		Value: value,
	}
}

// IntProperty creates a Property that contains an integer (32 or 64 bit based on architecture) value
func IntProperty(name string, value int) Property {
	return Property{
		Name:  name,
		Type:  Int,
		Value: value,
	}
}

// Int8Property creates a Property that contains an 8-bit integer value
func Int8Property(name string, value int8) Property {
	return Property{
		Name:  name,
		Type:  Int8,
		Value: value,
	}
}

// Int16Property creates a Property that contains an 16-bit integer value
func Int16Property(name string, value int16) Property {
	return Property{
		Name:  name,
		Type:  Int16,
		Value: value,
	}
}

// Int32Property creates a Property that contains an 32-bit integer value
func Int32Property(name string, value int32) Property {
	return Property{
		Name:  name,
		Type:  Int32,
		Value: value,
	}
}

// Int64Property creates a Property that contains an 64-bit integer value
func Int64Property(name string, value int64) Property {
	return Property{
		Name:  name,
		Type:  Int64,
		Value: value,
	}
}

// UintProperty creates a Property that contains an unsigned integer value (32 or 64-bit depending on architecture)
func UintProperty(name string, value uint) Property {
	return Property{
		Name:  name,
		Type:  Uint,
		Value: value,
	}
}

// Uint8Property creates a Property that contains an 8-bit unsigned integer value
func Uint8Property(name string, value uint8) Property {
	return Property{
		Name:  name,
		Type:  Uint8,
		Value: value,
	}
}

// Uint16Property creates a Property that contains a 16-bit unsigned integer value
func Uint16Property(name string, value uint16) Property {
	return Property{
		Name:  name,
		Type:  Uint16,
		Value: value,
	}
}

// Uint32Property creates a Property that contains a 32-bit unsigned integer value
func Uint32Property(name string, value uint32) Property {
	return Property{
		Name:  name,
		Type:  Uint32,
		Value: value,
	}
}

// Uint64Property creates a Property that contains a 64-bit integer value
func Uint64Property(name string, value uint64) Property {
	return Property{
		Name:  name,
		Type:  Uint64,
		Value: value,
	}
}

// Float32Property creates a Property that contains a 32-bit floating point value
func Float32Property(name string, value float32) Property {
	return Property{
		Name:  name,
		Type:  Float32,
		Value: value,
	}
}

// Float64Property creates a Property that contains a 64-bit floating point value
func Float64Property(name string, value float64) Property {
	return Property{
		Name:  name,
		Type:  Float64,
		Value: value,
	}
}

// ByteArrayProperty creates a Property that contains a byte array value
func ByteArrayProperty(name string, value []byte) Property {
	return Property{
		Name:  name,
		Type:  ByteArray,
		Value: value,
	}
}

// DateTimeProperty creates a Property that contains a date/time value
func DateTimeProperty(name string, value time.Time) Property {
	return Property{
		Name:  name,
		Type:  DateTime,
		Value: value,
	}
}

// DurationProperty creates a Property that contains a duration value
func DurationProperty(name string, value time.Duration) Property {
	return Property{
		Name:  name,
		Type:  Duration,
		Value: value,
	}
}

// DecimalProperty creates a Property that contains a decimal value
func DecimalProperty(name string, value decimal.Decimal) Property {
	return Property{
		Name:  name,
		Type:  Decimal,
		Value: value,
	}
}

// InterfaceProperty creates a Property that contains an interface{} value
func InterfaceProperty(name string, value interface{}) Property {
	return Property{
		Name:  name,
		Type:  Interface,
		Value: value,
	}
}

// BoolArrayProperty creates a Property that contains an array of boolean values
func BoolArrayProperty(name string, value ...bool) Property {
	return Property{
		Name:  name,
		Type:  BoolArray,
		Value: value,
	}
}

// IntArrayProperty creates a Property that contains an array of integer values
func IntArrayProperty(name string, value ...int) Property {
	return Property{
		Name:  name,
		Type:  IntArray,
		Value: value,
	}
}

// Int8ArrayProperty creates a Property that contains an array of 8-bit integer values
func Int8ArrayProperty(name string, value ...int8) Property {
	return Property{
		Name:  name,
		Type:  Int8Array,
		Value: value,
	}
}

// Int16ArrayProperty creates a Property that contains an array of 16-bit integer values
func Int16ArrayProperty(name string, value ...int16) Property {
	return Property{
		Name:  name,
		Type:  Int16Array,
		Value: value,
	}
}

// Int32ArrayProperty creates a Property that contains an array of 32-bit integer values
func Int32ArrayProperty(name string, value ...int32) Property {
	return Property{
		Name:  name,
		Type:  Int32Array,
		Value: value,
	}
}

// Int64ArrayProperty creates a Property that contains an array of 64-bit integer values
func Int64ArrayProperty(name string, value ...int64) Property {
	return Property{
		Name:  name,
		Type:  Int64Array,
		Value: value,
	}
}

// UintArrayProperty creates a Property that contains an array of unsigned integer values
func UintArrayProperty(name string, value ...uint) Property {
	return Property{
		Name:  name,
		Type:  UintArray,
		Value: value,
	}
}

// Uint8ArrayProperty creates a Property that contains an array of 8-bit unsigned integer values
func Uint8ArrayProperty(name string, value ...uint8) Property {
	return Property{
		Name:  name,
		Type:  Uint8Array,
		Value: value,
	}
}

// Uint16ArrayProperty creates a Property that contains an array of 16-bit unsigned integer values
func Uint16ArrayProperty(name string, value ...uint16) Property {
	return Property{
		Name:  name,
		Type:  Uint16Array,
		Value: value,
	}
}

// Uint32ArrayProperty creates a Property that contains an array of 32-bit unsigned integer values
func Uint32ArrayProperty(name string, value ...uint32) Property {
	return Property{
		Name:  name,
		Type:  Uint32Array,
		Value: value,
	}
}

// Uint64ArrayProperty creates a Property that contains an array of 64-bit unsigned integer values
func Uint64ArrayProperty(name string, value ...uint64) Property {
	return Property{
		Name:  name,
		Type:  Uint64Array,
		Value: value,
	}
}

// Float32ArrayProperty creates a Property that contains an array of 32 bit floating point values
func Float32ArrayProperty(name string, value ...float32) Property {
	return Property{
		Name:  name,
		Type:  Float32Array,
		Value: value,
	}
}

// Float64ArrayProperty creates a Property that contains an array of 64 bit floating point values
func Float64ArrayProperty(name string, value ...float64) Property {
	return Property{
		Name:  name,
		Type:  Float64Array,
		Value: value,
	}
}

// DateTimeArrayProperty creates a Property that contains an array of Date/Time values
func DateTimeArrayProperty(name string, value ...time.Time) Property {
	return Property{
		Name:  name,
		Type:  DateTimeArray,
		Value: value,
	}
}

// DurationArrayProperty creates a Property that contains an array of Duration values
func DurationArrayProperty(name string, value ...time.Duration) Property {
	return Property{
		Name:  name,
		Type:  DurationArray,
		Value: value,
	}
}

// DecimalArrayProperty creates a Property that contains an array of decimal values
func DecimalArrayProperty(name string, value ...decimal.Decimal) Property {
	return Property{
		Name:  name,
		Type:  DecimalArray,
		Value: value,
	}
}

// InterfaceArrayProperty creates a Property that contains an array of interface{} values
func InterfaceArrayProperty(name string, value ...interface{}) Property {
	return Property{
		Name:  name,
		Type:  InterfaceArray,
		Value: value,
	}
}

// NonEmptyString returns the string value of the property if possible,
// or it will return the default value it is given. If the string value
// of the property is empty, the default value is returned
func (p Property) NonEmptyString(d string) string {
	v := p.String()
	if v == "" {
		return d
	}
	return v
}

// String returns the string representation of the value stored in the property
//
//nolint:funlen
func (p Property) String() string {
	switch p.Type {
	case Bool:
		if p.Value == nil {
			return falseString
		}
		return fmt.Sprintf("%t", p.Value)
	case Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64:
		if p.Value == nil {
			return "0"
		}
		return fmt.Sprintf("%d", p.Value)
	case Float32, Float64:
		if p.Value == nil {
			return floatZeroString
		}
		return fmt.Sprintf("%.6f", p.Value)
	case DateTime:
		if p.Value == nil {
			p.Value = time.Unix(0, 0).UTC()
		}
		t, ok := p.Value.(time.Time)
		if !ok {
			return time.Unix(0, 0).UTC().Format(dtFormat)
		}

		return t.Format(dtFormat)
	case Duration:
		if p.Value == nil {
			p.Value = time.Duration(0)
		}
		d, ok := p.Value.(time.Duration)
		if !ok {
			return time.Duration(0).String()
		}
		return d.String()
	case Decimal:
		if p.Value == nil {
			return floatZeroString
		}
		d, ok := p.Value.(decimal.Decimal)
		if !ok {
			return floatZeroString
		}

		return d.StringFixed(dps)
	case Interface:
		if p.Value == nil {
			return emptyInterfaceString
		}
		return fmt.Sprintf("%v", p.Value)
	case ByteArray:
		if p.Value == nil {
			return emptyArrayString
		}

		bs, ok := p.Value.([]byte)
		if !ok {
			return emptyArrayString
		}
		return string(bs)
	case BoolArray:
		if p.Value == nil {
			p.Value = []bool{}
		}

		bs, ok := p.Value.([]bool)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range bs {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%t", v))
		}
		b.WriteString("]")
		return b.String()
	case IntArray:
		if p.Value == nil {
			p.Value = []int{}
		}

		is, ok := p.Value.([]int)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range is {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}

		b.WriteString("]")
		return b.String()
	case Int8Array:
		if p.Value == nil {
			p.Value = []int8{}
		}
		i8s, ok := p.Value.([]int8)
		if !ok {
			return emptyArrayString
		}
		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range i8s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}

		b.WriteString("]")
		return b.String()
	case Int16Array:
		if p.Value == nil {
			p.Value = []int16{}
		}

		i16s, ok := p.Value.([]int16)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range i16s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}

		b.WriteString("]")
		return b.String()
	case Int32Array:
		if p.Value == nil {
			p.Value = []int32{}
		}

		i32s, ok := p.Value.([]int32)
		if !ok {
			return emptyArrayString
		}
		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range i32s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}

		b.WriteString("]")
		return b.String()
	case Int64Array:
		if p.Value == nil {
			p.Value = []int64{}
		}

		i64s, ok := p.Value.([]int64)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range i64s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}

		b.WriteString("]")
		return b.String()
	case UintArray:
		if p.Value == nil {
			p.Value = []uint{}
		}

		us, ok := p.Value.([]uint)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")

		for i, v := range us {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}
		b.WriteString("]")
		return b.String()
	case Uint8Array:
		if p.Value == nil {
			p.Value = []uint8{}
		}

		u8s, ok := p.Value.([]uint8)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")

		for i, v := range u8s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}
		b.WriteString("]")
		return b.String()
	case Uint16Array:
		if p.Value == nil {
			p.Value = []uint16{}
		}

		u16s, ok := p.Value.([]uint16)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")

		for i, v := range u16s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}
		b.WriteString("]")
		return b.String()
	case Uint32Array:
		if p.Value == nil {
			p.Value = []uint32{}
		}

		u32s, ok := p.Value.([]uint32)
		if !ok {
			return emptyArrayString
		}
		b := new(strings.Builder)
		b.WriteString("[")

		for i, v := range u32s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}
		b.WriteString("]")
		return b.String()
	case Uint64Array:
		if p.Value == nil {
			p.Value = []uint64{}
		}

		u64s, ok := p.Value.([]uint64)
		if !ok {
			return emptyArrayString
		}
		b := new(strings.Builder)
		b.WriteString("[")

		for i, v := range u64s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%d", v))
		}
		b.WriteString("]")
		return b.String()
	case Float32Array:
		if p.Value == nil {
			p.Value = []float32{}
		}

		f32s, ok := p.Value.([]float32)
		if !ok {
			return emptyArrayString
		}
		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range f32s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%.6f", v))
		}
		b.WriteString("]")
		return b.String()
	case Float64Array:
		if p.Value == nil {
			p.Value = []float64{}
		}

		f64s, ok := p.Value.([]float64)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range f64s {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%.6f", v))
		}
		b.WriteString("]")
		return b.String()
	case DateTimeArray:
		if p.Value == nil {
			p.Value = []time.Time{}
		}

		ts, ok := p.Value.([]time.Time)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range ts {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(v.Format(dtFormat))
		}
		b.WriteString("]")
		return b.String()
	case DurationArray:
		if p.Value == nil {
			p.Value = []time.Duration{}
		}

		ds, ok := p.Value.([]time.Duration)
		if !ok {
			return emptyArrayString
		}
		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range ds {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(v.String())
		}
		b.WriteString("]")
		return b.String()
	case DecimalArray:
		if p.Value == nil {
			p.Value = []decimal.Decimal{}
		}

		ds, ok := p.Value.([]decimal.Decimal)
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range ds {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(v.StringFixed(dps))
		}
		b.WriteString("]")
		return b.String()
	case String:
		if p.Value == nil {
			return ""
		}

		v, ok := p.Value.(string)
		if !ok {
			return ""
		}

		return v
	case StringArray:
		if p.Value == nil {
			return emptyArrayString
		}

		ss, ok := p.Value.([]string)
		if !ok {
			return emptyArrayString
		}
		b := new(strings.Builder)
		b.WriteString("[\"")
		for i, v := range ss {
			if i > 0 {
				b.WriteString("\",\"")
			}
			b.WriteString(v)
		}
		b.WriteString("\"]")
		return b.String()
	case InterfaceArray:
		if p.Value == nil {
			return emptyArrayString
		}

		is, ok := p.Value.([]interface{})
		if !ok {
			return emptyArrayString
		}

		b := new(strings.Builder)
		b.WriteString("[")
		for i, v := range is {
			if i > 0 {
				b.WriteRune(',')
			}
			b.WriteString(fmt.Sprintf("%v", v))
		}
		b.WriteString("]")
		return b.String()
	case Undefined:
		return ""
	default:
		return ""
	}
}

func getPropertyError(field, propertyType string) error {
	return fmt.Errorf("%s is not an %s value", field, propertyType)
}

// Int returns the value of the property as an int if possible and an error if it is not
func (p Property) Int() (int, error) {
	if p.Type != Int {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(int)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// IntOrDefault returns the value of the property as an int if possible or the provided default value if it is not
func (p Property) IntOrDefault(d int) int {
	v, err := p.Int()
	if err != nil {
		return d
	}
	return v
}

// Uint returns the value of the property as an uint if possible and an error if it is not
func (p Property) Uint() (uint, error) {
	if p.Type != Uint {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(uint)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// UintOrDefault returns the value of the property as an uint if possible or the provided default value if it is not
func (p Property) UintOrDefault(d uint) uint {
	v, err := p.Uint()
	if err != nil {
		return d
	}
	return v
}

// Float64 returns the value of the property as a float64 if possible and an error if it is not
func (p Property) Float64() (float64, error) {
	if p.Type != Float64 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(float64)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Float64OrDefault returns the value of the property as a float64 if possible or the provided default value if it is not
func (p Property) Float64OrDefault(d float64) float64 {
	v, err := p.Float64()
	if err != nil {
		return d
	}
	return v
}

// Float32 returns the value of the property as a float32 if possible and an error if it is not
func (p Property) Float32() (float32, error) {
	if p.Type != Float32 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(float32)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Float32OrDefault returns the value of the property as a float32 if possible or the provided default value if it is not
func (p Property) Float32OrDefault(d float32) float32 {
	v, err := p.Float32()
	if err != nil {
		return d
	}
	return v
}

// Int8 returns the value of the property as an int8 if possible and an error if it is not
func (p Property) Int8() (int8, error) {
	if p.Type != Int8 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(int8)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int8OrDefault returns the value of the property as an int8 if possible or the provided default value if it is not
func (p Property) Int8OrDefault(d int8) int8 {
	v, err := p.Int8()
	if err != nil {
		return d
	}
	return v
}

// Int16 returns the value of the property as an int8 if possible and an error if it is not
func (p Property) Int16() (int16, error) {
	if p.Type != Int16 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(int16)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int16OrDefault returns the value of the property as an int8 if possible or the provided default value if it is not
func (p Property) Int16OrDefault(d int16) int16 {
	v, err := p.Int16()
	if err != nil {
		return d
	}
	return v
}

// Int32 returns the value of the property as an int8 if possible and an error if it is not
func (p Property) Int32() (int32, error) {
	if p.Type != Int32 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(int32)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int32OrDefault returns the value of the property as an int32 if possible or the provided default value if it is not
func (p Property) Int32OrDefault(d int32) int32 {
	v, err := p.Int32()
	if err != nil {
		return d
	}
	return v
}

// Int64 returns the value of the property as an int64 if possible and an error if it is not
func (p Property) Int64() (int64, error) {
	if p.Type != Int64 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(int64)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int64OrDefault returns the value of the property as an int64 if possible or the provided default value if it is not
func (p Property) Int64OrDefault(d int64) int64 {
	v, err := p.Int64()
	if err != nil {
		return d
	}
	return v
}

// Uint8 returns the value of the property as an uint8 if possible and an error if it is not
func (p Property) Uint8() (uint8, error) {
	if p.Type != Uint8 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(uint8)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint8OrDefault returns the value of the property as an uint8 if possible or the provided default value if it is not
func (p Property) Uint8OrDefault(d uint8) uint8 {
	v, err := p.Uint8()
	if err != nil {
		return d
	}
	return v
}

// Uint16 returns the value of the property as an uint16 if possible and an error if it is not
func (p Property) Uint16() (uint16, error) {
	if p.Type != Uint16 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(uint16)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint16OrDefault returns the value of the property as an uint16 if possible or the provided default value if it is not
func (p Property) Uint16OrDefault(d uint16) uint16 {
	v, err := p.Uint16()
	if err != nil {
		return d
	}
	return v
}

// Uint32 returns the value of the property as an uint32 if possible and an error if it is not
func (p Property) Uint32() (uint32, error) {
	if p.Type != Uint32 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(uint32)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint32OrDefault returns the value of the property as an uint32 if possible or the provided default value if it is not
func (p Property) Uint32OrDefault(d uint32) uint32 {
	v, err := p.Uint32()
	if err != nil {
		return d
	}
	return v
}

// Uint64 returns the value of the property as an uint64 if possible and an error if it is not
func (p Property) Uint64() (uint64, error) {
	if p.Type != Uint64 {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(uint64)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint64OrDefault returns the value of the property as an uint64 if possible or the provided default value if it is not
func (p Property) Uint64OrDefault(d uint64) uint64 {
	v, err := p.Uint64()
	if err != nil {
		return d
	}
	return v
}

// DateTime returns the value of the property as a date/time if possible and an error if it is not
func (p Property) DateTime() (dt time.Time, err error) {
	if p.Type != DateTime {
		return time.Unix(0, 0), getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(time.Time)

	if !ok {
		return time.Unix(0, 0), getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// DateTimeOrDefault returns the value of the property as a date/time if possible or the provided default value if it is not
func (p Property) DateTimeOrDefault(d time.Time) time.Time {
	v, err := p.DateTime()
	if err != nil {
		return d
	}
	return v
}

// Duration returns the value of the property as a duration if possible and an error if it is not
func (p Property) Duration() (d time.Duration, err error) {
	if p.Type != Duration {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(time.Duration)

	if !ok {
		return 0, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// DurationOrDefault returns the value of the property as a duration if possible or the provided default value if it is not
func (p Property) DurationOrDefault(d time.Duration) time.Duration {
	v, err := p.Duration()
	if err != nil {
		return d
	}
	return v
}

// Decimal returns the value of the property as a decimal if possible and an error if it is not
func (p Property) Decimal() (decimal.Decimal, error) {
	if p.Type != Decimal {
		return decimal.Zero, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(decimal.Decimal)

	if !ok {
		return decimal.Zero, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// DecimalOrDefault returns the value of the property as a decimal if possible or the provided default value if it is not
func (p Property) DecimalOrDefault(d decimal.Decimal) decimal.Decimal {
	v, err := p.Decimal()
	if err != nil {
		return d
	}
	return v
}

func stringToBool(v string) (b bool, err error) {
	switch strings.ToUpper(v) {
	case "Y", "YES":
		return true, nil
	case "N", "NO":
		return false, nil
	default:
		return strconv.ParseBool(v)
	}
}

// Bool returns the value of the property as a boolean value if possible and an error if it is not
func (p Property) Bool() (bool, error) {
	if p.Type != Bool {
		return false, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.(bool)

	if !ok {
		return false, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// BoolOrDefault returns the value of the property as a boolean value if possible or the provided default value if it is not
func (p Property) BoolOrDefault(b bool) bool {
	v, err := p.Bool()
	if err != nil {
		return b
	}
	return v
}

// BoolArray returns the value of the property as an array of boolean values if possible and an error if it is not
func (p Property) BoolArray() ([]bool, error) {
	if p.Type != BoolArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]bool)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// BoolArrayOrDefault returns the value of the property as an array of boolean values if possible or the provided default value if it is not
func (p Property) BoolArrayOrDefault(d []bool) []bool {
	v, err := p.BoolArray()
	if err != nil {
		return d
	}
	return v
}

// IntArray returns the value of the property as an array of int values if possible and an error if it is not
func (p Property) IntArray() ([]int, error) {
	if p.Type != IntArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]int)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// IntArrayOrDefault returns the value of the property as an array of int values if possible or the provided default value if it is not
func (p Property) IntArrayOrDefault(d []int) []int {
	v, err := p.IntArray()
	if err != nil {
		return d
	}
	return v
}

// Int8Array returns the value of the property as an array of int8 values if possible and an error if it is not
func (p Property) Int8Array() ([]int8, error) {
	if p.Type != Int8Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]int8)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int8ArrayOrDefault returns the value of the property as an array of int8 values if possible or the provided default value if it is not
func (p Property) Int8ArrayOrDefault(d []int8) []int8 {
	v, err := p.Int8Array()
	if err != nil {
		return d
	}
	return v
}

// Int16Array returns the value of the property as an array of int16 values if possible and an error if it is not
func (p Property) Int16Array() ([]int16, error) {
	if p.Type != Int16Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]int16)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int16ArrayOrDefault returns the value of the property as an array of int16 values if possible or the provided default value if it is not
func (p Property) Int16ArrayOrDefault(d []int16) []int16 {
	v, err := p.Int16Array()
	if err != nil {
		return d
	}
	return v
}

// Int32Array returns the value of the property as an array of int32 values if possible and an error if it is not
func (p Property) Int32Array() ([]int32, error) {
	if p.Type != Int32Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]int32)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int32ArrayOrDefault returns the value of the property as an array of int32 values if possible or the provided default value if it is not
func (p Property) Int32ArrayOrDefault(d []int32) []int32 {
	v, err := p.Int32Array()
	if err != nil {
		return d
	}
	return v
}

// Int64Array returns the value of the property as an array of int64 values if possible and an error if it is not
func (p Property) Int64Array() ([]int64, error) {
	if p.Type != Int64Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]int64)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Int64ArrayOrDefault returns the value of the property as an array of int64 values if possible or the provided default value if it is not
func (p Property) Int64ArrayOrDefault(d []int64) []int64 {
	v, err := p.Int64Array()
	if err != nil {
		return d
	}
	return v
}

// UintArray returns the value of the property as an array of uint values if possible and an error if it is not
func (p Property) UintArray() ([]uint, error) {
	if p.Type != UintArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]uint)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// UintArrayOrDefault returns the value of the property as an array of uint values if possible or the provided default value if it is not
func (p Property) UintArrayOrDefault(d []uint) []uint {
	v, err := p.UintArray()
	if err != nil {
		return d
	}
	return v
}

// Uint8Array returns the value of the property as an array of uint8 values if possible and an error if it is not
func (p Property) Uint8Array() ([]uint8, error) {
	if p.Type != Uint8Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]uint8)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint8ArrayOrDefault returns the value of the property as an array of uint8 values if possible or the provided default value if it is not
func (p Property) Uint8ArrayOrDefault(d []uint8) []uint8 {
	v, err := p.Uint8Array()
	if err != nil {
		return d
	}
	return v
}

// Uint16Array returns the value of the property as an array of uint16 values if possible and an error if it is not
func (p Property) Uint16Array() ([]uint16, error) {
	if p.Type != Uint16Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]uint16)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint16ArrayOrDefault returns the value of the property as an array of uint16 values if possible
// or the provided default value if it is not
func (p Property) Uint16ArrayOrDefault(d []uint16) []uint16 {
	v, err := p.Uint16Array()
	if err != nil {
		return d
	}
	return v
}

// Uint32Array returns the value of the property as an array of uint32 values if possible and an error if it is not
func (p Property) Uint32Array() ([]uint32, error) {
	if p.Type != Uint32Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]uint32)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint32ArrayOrDefault returns the value of the property as an array of uint32 values if possible
// or the provided default value if it is not
func (p Property) Uint32ArrayOrDefault(d []uint32) []uint32 {
	v, err := p.Uint32Array()
	if err != nil {
		return d
	}
	return v
}

// Uint64Array returns the value of the property as an array of uint64 values if possible and an error if it is not
func (p Property) Uint64Array() ([]uint64, error) {
	if p.Type != Uint64Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]uint64)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Uint64ArrayOrDefault returns the value of the property as an array of uint64 values if possible
// or the provided default value if it is not
func (p Property) Uint64ArrayOrDefault(d []uint64) []uint64 {
	v, err := p.Uint64Array()
	if err != nil {
		return d
	}
	return v
}

// Float32Array returns the value of the property as an array of float32 values if possible and an error if it is not
func (p Property) Float32Array() ([]float32, error) {
	if p.Type != Float32Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]float32)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Float32ArrayOrDefault returns the value of the property as an array of float32 values if possible
// or the provided default value if it is not
func (p Property) Float32ArrayOrDefault(d []float32) []float32 {
	v, err := p.Float32Array()
	if err != nil {
		return d
	}
	return v
}

// Float64Array returns the value of the property as a float64 of boolean values if possible and an error if it is not
func (p Property) Float64Array() ([]float64, error) {
	if p.Type != Float64Array {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]float64)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// Float64ArrayOrDefault returns the value of the property as an array of float64 values if possible
// or the provided default value if it is not
func (p Property) Float64ArrayOrDefault(d []float64) []float64 {
	v, err := p.Float64Array()
	if err != nil {
		return d
	}
	return v
}

// StringArray returns the value of the property as an array of string values if possible and an error if it is not
//
//nolint:funlen
func (p Property) StringArray() ([]string, error) {
	switch p.Type {
	case Int:
		v, err := p.Int()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{strconv.Itoa(v)}, nil
	case Int8:
		v, err := p.Int8()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Int16:
		v, err := p.Int16()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Int32:
		v, err := p.Int32()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Int64:
		v, err := p.Int64()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Uint:
		v, err := p.Uint()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Uint8:
		v, err := p.Uint8()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Uint16:
		v, err := p.Uint16()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Uint32:
		v, err := p.Uint32()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case Uint64:
		v, err := p.Uint64()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{fmt.Sprintf("%d", v)}, nil
	case IntArray:
		vs, err := p.IntArray()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = strconv.Itoa(v)
		}

		return rs, nil
	case Int8Array:
		vs, err := p.Int8Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case Int16Array:
		vs, err := p.Int16Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case Int32Array:
		vs, err := p.Int32Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case Int64Array:
		vs, err := p.Int64Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case UintArray:
		vs, err := p.UintArray()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case Uint8Array:
		vs, err := p.Uint8Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case Uint16Array:
		vs, err := p.Uint16Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case Uint32Array:
		vs, err := p.Uint32Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}
		return rs, nil
	case Uint64Array:
		vs, err := p.Uint64Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%d", v)
		}

		return rs, nil
	case Bool:
		v, err := p.Bool()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		b := falseString

		if v {
			b = trueString
		}

		return []string{b}, nil
	case BoolArray:
		vs, err := p.BoolArray()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			if v {
				rs[i] = trueString
			} else {
				rs[i] = falseString
			}
		}

		return rs, nil
	case Float32:
		v, err := p.Float32()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		return []string{fmt.Sprintf("%.6f", v)}, nil
	case Float32Array:
		vs, err := p.Float32Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%.6f", v)
		}

		return rs, nil
	case Float64:
		v, err := p.Float64()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		return []string{fmt.Sprintf("%.6f", v)}, nil
	case Float64Array:
		vs, err := p.Float64Array()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = fmt.Sprintf("%.6f", v)
		}

		return rs, nil
	case String:
		v, ok := p.Value.(string)
		if !ok {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{v}, nil
	case StringArray:
		v, ok := p.Value.([]string)
		if !ok {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return v, nil
	case DateTime:
		v, ok := p.Value.(time.Time)
		if !ok {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{v.Format(dtFormat)}, nil
	case DateTimeArray:
		vs, err := p.DateTimeArray()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = v.Format(dtFormat)
		}

		return rs, nil
	case Duration:
		v, ok := p.Value.(time.Duration)
		if !ok {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{v.String()}, nil
	case DurationArray:
		vs, err := p.DurationArray()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = v.String()
		}

		return rs, nil
	case Decimal:
		v, ok := p.Value.(decimal.Decimal)
		if !ok {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		return []string{v.StringFixed(dps)}, nil
	case DecimalArray:
		vs, err := p.DecimalArray()
		if err != nil {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}

		rs := make([]string, len(vs))

		for i, v := range vs {
			rs[i] = v.StringFixed(dps)
		}

		return rs, nil
	case Interface:
		return []string{fmt.Sprintf("%v", p.Value)}, nil
	case InterfaceArray:
		v, ok := p.Value.([]interface{})
		if !ok {
			return nil, NewTypeError(p.Type, "cannot return value for data type")
		}
		vs := make([]string, len(v))
		for i, vv := range v {
			vs[i] = fmt.Sprintf("%v", vv)
		}

		return vs, nil
	case ByteArray, Undefined:
		return nil, NewTypeError(p.Type, "cannot return value for data type")
	default:
		return nil, NewTypeError(p.Type, "cannot return value for data type")
	}
}

// StringArrayOrDefault returns the value of the property as an array of string values if possible
// or the provided default value if it is not
func (p Property) StringArrayOrDefault(d []string) []string {
	v, err := p.StringArray()
	if err != nil {
		return d
	}
	return v
}

// DateTimeArray returns the value of the property as an array of date/time values if possible and an error if it is not
func (p Property) DateTimeArray() ([]time.Time, error) {
	if p.Type != DateTimeArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]time.Time)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// DateTimeArrayOrDefault returns the value of the property as an array of date/time values if possible
// or the provided default value if it is not
func (p Property) DateTimeArrayOrDefault(d []time.Time) []time.Time {
	v, err := p.DateTimeArray()
	if err != nil {
		return d
	}
	return v
}

// DurationArray returns the value of the property as an array of duration values if possible and an error if it is not
func (p Property) DurationArray() ([]time.Duration, error) {
	if p.Type != DurationArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]time.Duration)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// DurationArrayOrDefault returns the value of the property as an array of duration values if possible
// or the provided default value if it is not
func (p Property) DurationArrayOrDefault(d []time.Duration) []time.Duration {
	v, err := p.DurationArray()
	if err != nil {
		return d
	}
	return v
}

// DecimalArray returns the value of the property as an array of decimal values if possible and an error if it is not
func (p Property) DecimalArray() ([]decimal.Decimal, error) {
	if p.Type != DecimalArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]decimal.Decimal)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// DecimalArrayOrDefault returns the value of the property as an array of decimal values if possible
// or the provided default value if it is not
func (p Property) DecimalArrayOrDefault(d []decimal.Decimal) []decimal.Decimal {
	v, err := p.DecimalArray()
	if err != nil {
		return d
	}
	return v
}

// ByteArray returns the value of the property as a byte array if possible and an error if it is not
func (p Property) ByteArray() ([]byte, error) {
	if p.Type != ByteArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	v, ok := p.Value.([]byte)

	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return v, nil
}

// ByteArrayOrDefault returns the value of the property as a byte array if possible
// or the provided default value if it is not
func (p Property) ByteArrayOrDefault(d []byte) []byte {
	v, err := p.ByteArray()
	if err != nil {
		return d
	}
	return v
}

// Interface returns the value of the property as an interface{} if possible
// and an error if the property type is undefined or the value is nil
func (p Property) Interface() (interface{}, error) {
	if p.Type == Undefined || p.Value == nil {
		return nil, getPropertyError(p.Name, p.Type.String())
	}

	return p.Value, nil
}

// InterfaceOrDefault returns the value of the property as an interface{} if possible
// or the provided default value if the data type is undefined or the value is nil
func (p Property) InterfaceOrDefault(d interface{}) interface{} {
	if p.Type == Undefined || p.Value == nil {
		return d
	}
	return p.Value
}

// InterfaceArray returns the value of the property as an array of interface{} if possible and an error
func (p Property) InterfaceArray() ([]interface{}, error) {
	if p.Type != InterfaceArray {
		return nil, getPropertyError(p.Name, p.Type.String())
	}
	v, ok := p.Value.([]interface{})
	if !ok {
		return nil, getPropertyError(p.Name, p.Type.String())
	}
	return v, nil
}

// InterfaceArrayOrDefault returns the value of the property as an array of interface{} if possible
func (p Property) InterfaceArrayOrDefault(d []interface{}) []interface{} {
	v, err := p.InterfaceArray()
	if err != nil {
		return d
	}
	return v
}

// New creates a new property with the provided name and value
// The data type of the property is determined by the type of the value
func New(name string, value interface{}) Property {
	return Property{
		Name:  name,
		Type:  getType(value),
		Value: value,
	}
}

func getType(value interface{}) Type {
	switch value.(type) {
	case bool:
		return Bool
	case int:
		return Int
	case int8:
		return Int8
	case int16:
		return Int16
	case int32:
		return Int32
	case int64:
		return Int64
	case uint:
		return Uint
	case uint8:
		return Uint8
	case uint16:
		return Uint16
	case uint32:
		return Uint32
	case uint64:
		return Uint64
	case float32:
		return Float32
	case float64:
		return Float64
	case []byte:
		return ByteArray
	case string:
		return String
	case time.Time:
		return DateTime
	case time.Duration:
		return Duration
	case decimal.Decimal:
		return Decimal
	case []bool:
		return BoolArray
	case []int:
		return IntArray
	case []int8:
		return Int8Array
	case []int16:
		return Int16Array
	case []int32:
		return Int32Array
	case []int64:
		return Int64Array
	case []uint:
		return UintArray
	case []uint16:
		return Uint16Array
	case []uint32:
		return Uint32Array
	case []uint64:
		return Uint64Array
	case []float32:
		return Float32Array
	case []float64:
		return Float64Array
	case []string:
		return StringArray
	case []time.Time:
		return DateTimeArray
	case []time.Duration:
		return DurationArray
	case []decimal.Decimal:
		return DecimalArray
	case []interface{}:
		return InterfaceArray
	default:
		return Interface
	}
}

func (p Property) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Property) UnmarshalBinary(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err := d.Decode(p); err != nil {
		return err
	}
	switch p.Value.(type) {
	case json.Number:
		num, ok := p.Value.(json.Number)
		if !ok {
			return fmt.Errorf("value %v is not a json.Number", p.Value)
		}
		if err := convertJSONNumber(p, num); err != nil {
			return err
		}
	case []json.Number:
		nums, ok := p.Value.([]json.Number)
		if !ok {
			return fmt.Errorf("value %v is not a []json.Number", p.Value)
		}
		if err := convertJSONNumberSlice(p, nums); err != nil {
			return err
		}
	case []interface{}:
		vs, ok := p.Value.([]interface{})
		if !ok {
			return fmt.Errorf("value %v is not a []interface{}", p.Value)
		}

		if len(vs) == 0 {
			return nil
		}
		switch vs[0].(type) {
		case json.Number:
			nums := make([]json.Number, len(vs))
			for i, v := range vs {
				nums[i], ok = v.(json.Number)
				if !ok {
					return fmt.Errorf("value %v is not a json.Number", p.Value)
				}
			}
			return convertJSONNumberSlice(p, nums)
		case string:
			return convertInterfaceSliceOfStrings(p, vs)
		case bool:
			return convertInterfaceSliceOfBools(p, vs)
		default:
			return nil
		}
	case interface{}:
		if err := convertInterface(p); err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}

//nolint:funlen
func convertJSONNumber(p *Property, num json.Number) error {
	switch p.Type {
	case Int:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		if v > math.MaxInt || v < math.MinInt {
			return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
		}
		p.Value = int(v)
	case Int8:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		if v > math.MaxInt8 || v < math.MinInt8 {
			return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
		}
		p.Value = int8(v)
	case Int16:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		if v > math.MaxInt16 || v < math.MinInt16 {
			return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
		}
		p.Value = int16(v)
	case Int32:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		if v > math.MaxInt32 || v < math.MinInt32 {
			return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
		}
		p.Value = int32(v)
	case Int64:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		p.Value = v
	case Uint:
		v, err := num.Float64()
		if err != nil {
			return err
		}
		if v > float64(math.MaxUint) || v < 0 {
			return fmt.Errorf("value %f out of range for data type %s", v, p.Type.String())
		}
		p.Value = uint(v)
	case Uint8:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		if v > math.MaxUint8 || v < 0 {
			return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
		}
		p.Value = uint8(v)
	case Uint16:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		if v > math.MaxUint16 || v < 0 {
			return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
		}
		p.Value = uint16(v)
	case Uint32:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		if v > math.MaxUint32 || v < 0 {
			return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
		}
		p.Value = uint32(v)
	case Uint64:
		v, err := num.Float64()
		if err != nil {
			return err
		}
		if v > float64(math.MaxUint64) || v < 0 {
			return fmt.Errorf("value %f out of range for data type %s", v, p.Type.String())
		}
		p.Value = uint64(v)
	case Float32:
		v, err := num.Float64()
		if err != nil {
			return err
		}
		if v > math.MaxFloat32 || v < -math.MaxFloat32 {
			return fmt.Errorf("value %f out of range for data type %s", v, p.Type.String())
		}
		p.Value = float32(v)
	case Float64:
		v, err := num.Float64()
		if err != nil {
			return err
		}
		p.Value = v
	case Decimal:
		v, err := num.Float64()
		if err != nil {
			return err
		}
		p.Value = decimal.NewFromFloat(v)
	case Duration:
		v, err := num.Int64()
		if err != nil {
			return err
		}
		p.Value = time.Duration(v)
	default:
		return nil
	}
	return nil
}

//nolint:funlen
func convertJSONNumberSlice(p *Property, nums []json.Number) error {
	switch p.Type {
	case IntArray:
		vs := make([]int, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			if v > math.MaxInt || v < math.MinInt {
				return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
			}
			vs[i] = int(v)
		}
		p.Value = vs
	case Int8Array:
		vs := make([]int8, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			if v > math.MaxInt8 || v < math.MinInt8 {
				return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
			}
			vs[i] = int8(v)
		}
		p.Value = vs
	case Int16Array:
		vs := make([]int16, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			if v > math.MaxInt16 || v < math.MinInt16 {
				return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
			}
			vs[i] = int16(v)
		}
		p.Value = vs
	case Int32Array:
		vs := make([]int32, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			if v > math.MaxInt32 || v < math.MinInt32 {
				return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
			}
			vs[i] = int32(v)
		}
		p.Value = vs
	case Int64Array:
		vs := make([]int64, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			vs[i] = v
		}
		p.Value = vs
	case UintArray:
		vs := make([]uint, len(nums))
		for i, num := range nums {
			v, err := num.Float64()
			if err != nil {
				return err
			}
			if v > float64(math.MaxUint) || v < 0 {
				return fmt.Errorf("value %f out of range for data type %s", v, p.Type.String())
			}
			vs[i] = uint(v)
		}
		p.Value = vs
	case Uint8Array:
		vs := make([]uint8, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			if v > math.MaxUint8 || v < 0 {
				return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
			}
			vs[i] = uint8(v)
		}
		p.Value = vs
	case Uint16Array:
		vs := make([]uint16, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			if v > math.MaxUint16 || v < 0 {
				return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
			}
			vs[i] = uint16(v)
		}
		p.Value = vs
	case Uint32Array:
		vs := make([]uint32, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			if v > math.MaxUint32 || v < 0 {
				return fmt.Errorf("value %d out of range for data type %s", v, p.Type.String())
			}
			vs[i] = uint32(v)
		}
		p.Value = vs
	case Uint64Array:
		vs := make([]uint64, len(nums))
		for i, num := range nums {
			v, err := num.Float64()
			if err != nil {
				return err
			}
			if v > float64(math.MaxUint64) || v < 0 {
				return fmt.Errorf("value %f out of range for data type %s", v, p.Type.String())
			}
			vs[i] = uint64(v)
		}
		p.Value = vs
	case Float32Array:
		vs := make([]float32, len(nums))
		for i, num := range nums {
			v, err := num.Float64()
			if err != nil {
				return err
			}
			if v > math.MaxFloat32 || v < -math.MaxFloat32 {
				return fmt.Errorf("value %f out of range for data type %s", v, p.Type.String())
			}
			vs[i] = float32(v)
		}
		p.Value = vs
	case Float64Array:
		vs := make([]float64, len(nums))
		for i, num := range nums {
			v, err := num.Float64()
			if err != nil {
				return err
			}
			vs[i] = v
		}
		p.Value = vs
	case DecimalArray:
		vs := make([]decimal.Decimal, len(nums))
		for i, num := range nums {
			v, err := num.Float64()
			if err != nil {
				return err
			}
			vs[i] = decimal.NewFromFloat(v)
		}
		p.Value = vs
	case DurationArray:
		vs := make([]time.Duration, len(nums))
		for i, num := range nums {
			v, err := num.Int64()
			if err != nil {
				return err
			}
			vs[i] = time.Duration(v)
		}
		p.Value = vs
	default:
		return nil
	}
	return nil
}

func convertInterface(p *Property) error {
	if p.Type == Interface {
		return nil
	}

	s, ok := p.Value.(string)
	if ok {
		return convertFromString(p, s)
	}

	fs := fmt.Sprintf("%v", p.Value)
	return convertFromString(p, fs)
}

func convertFromString(p *Property, s string) error {
	switch p.Type {
	case Bool:
		v, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		p.Value = v
	case ByteArray:
		p.Value = []byte(s)
	case String:
		p.Value = s
	case Decimal:
		v, err := decimal.NewFromString(s)
		if err != nil {
			return err
		}
		p.Value = v
	case DateTime:
		v, err := time.Parse(dtFormat, s)
		if err != nil {
			return err
		}
		p.Value = v
	case Duration:
		v, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		p.Value = v
	default:
		return nil
	}
	return nil
}

func convertInterfaceSliceOfStrings(p *Property, vs []interface{}) error {
	switch p.Type {
	case BoolArray:
		bs := make([]bool, len(vs))
		for i, v := range vs {
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("value %v is not a string", p.Value)
			}
			b, err := strconv.ParseBool(s)
			if err != nil {
				return err
			}
			bs[i] = b
		}
		p.Value = bs
	case StringArray:
		ss := make([]string, len(vs))
		for i, v := range vs {
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("value %v is not a string", p.Value)
			}
			ss[i] = s
		}
		p.Value = ss
	case DateTimeArray:
		dts := make([]time.Time, len(vs))
		for i, v := range vs {
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("value %v is not a string", p.Value)
			}
			d, err := time.Parse(dtFormat, s)
			if err != nil {
				return err
			}
			dts[i] = d
		}
		p.Value = dts
	case DurationArray:
		ds := make([]time.Duration, len(vs))
		for i, v := range vs {
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("value %v is not a string", p.Value)
			}
			d, err := time.ParseDuration(s)
			if err != nil {
				return err
			}
			ds[i] = d
		}
		p.Value = ds
	case DecimalArray:
		ds := make([]decimal.Decimal, len(vs))
		for i, v := range vs {
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("value %v is not a string", p.Value)
			}
			d, err := decimal.NewFromString(s)
			if err != nil {
				return err
			}
			ds[i] = d
		}
		p.Value = ds
	default:
		return nil
	}
	return nil
}

func convertInterfaceSliceOfBools(p *Property, vs []interface{}) error {
	bs := make([]bool, len(vs))
	for i, v := range vs {
		b, ok := v.(bool)
		if !ok {
			return fmt.Errorf("value %v is not a bool", p.Value)
		}
		bs[i] = b
	}
	p.Value = bs
	return nil
}

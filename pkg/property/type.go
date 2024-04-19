package property

type Type uint8

const (
	Undefined Type = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	ByteArray
	String
	DateTime
	Duration
	Decimal
	Interface
	BoolArray
	IntArray
	Int8Array
	Int16Array
	Int32Array
	Int64Array
	UintArray
	Uint8Array
	Uint16Array
	Uint32Array
	Uint64Array
	Float32Array
	Float64Array
	StringArray
	DateTimeArray
	DurationArray
	DecimalArray
	InterfaceArray
)

// String returns the name of the Type as a string
// E.g. Bool, Int, String, Float32 etc.
func (p Type) String() string {
	switch p {
	case Bool:
		return TypeLabelBool
	case Int:
		return TypeLabelInt
	case Int8:
		return TypeLabelInt8
	case Int16:
		return TypeLabelInt16
	case Int32:
		return TypeLabelInt32
	case Int64:
		return TypeLabelInt64
	case Uint:
		return TypeLabelUint
	case Uint8:
		return TypeLabelUint8
	case Uint16:
		return TypeLabelUint16
	case Uint32:
		return TypeLabelUint32
	case Uint64:
		return TypeLabelUint64
	case Float32:
		return TypeLabelFloat32
	case Float64:
		return TypeLabelFloat64
	case ByteArray:
		return TypeLabelByteArray
	case String:
		return TypeLabelString
	case DateTime:
		return TypeLabelDateTime
	case Duration:
		return TypeLabelDuration
	case Decimal:
		return TypeLabelDecimal
	case Interface:
		return TypeLabelInterface
	case BoolArray:
		return TypeLabelBoolArray
	case IntArray:
		return TypeLabelIntArray
	case Int8Array:
		return TypeLabelInt8Array
	case Int16Array:
		return TypeLabelInt16Array
	case Int32Array:
		return TypeLabelInt32Array
	case Int64Array:
		return TypeLabelInt64Array
	case UintArray:
		return TypeLabelUintArray
	case Uint8Array:
		return TypeLabelUint8Array
	case Uint16Array:
		return TypeLabelUint16Array
	case Uint32Array:
		return TypeLabelUint32Array
	case Uint64Array:
		return TypeLabelUint64Array
	case Float32Array:
		return TypeLabelFloat32Array
	case Float64Array:
		return TypeLabelFloat64Array
	case StringArray:
		return TypeLabelStringArray
	case DateTimeArray:
		return TypeLabelDateTimeArray
	case DurationArray:
		return TypeLabelDurationArray
	case DecimalArray:
		return TypeLabelDecimalArray
	case InterfaceArray:
		return TypeLabelInterfaceArray
	case Undefined:
		return TypeLabelUndefined
	default:
		return TypeLabelUndefined
	}
}

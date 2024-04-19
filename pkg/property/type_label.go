package property

type TypeLabel = string

const (
	TypeLabelUndefined      TypeLabel = "Undefined"
	TypeLabelBool           TypeLabel = "Bool"
	TypeLabelInt            TypeLabel = "Int"
	TypeLabelInt8           TypeLabel = "Int8"
	TypeLabelInt16          TypeLabel = "Int16"
	TypeLabelInt32          TypeLabel = "Int32"
	TypeLabelInt64          TypeLabel = "Int64"
	TypeLabelUint           TypeLabel = "Uint"
	TypeLabelUint8          TypeLabel = "Uint8"
	TypeLabelUint16         TypeLabel = "Uint16"
	TypeLabelUint32         TypeLabel = "Uint32"
	TypeLabelUint64         TypeLabel = "Uint64"
	TypeLabelFloat32        TypeLabel = "Float32"
	TypeLabelFloat64        TypeLabel = "Float64"
	TypeLabelByteArray      TypeLabel = "ByteArray"
	TypeLabelString         TypeLabel = "String"
	TypeLabelDateTime       TypeLabel = "DateTime"
	TypeLabelDecimal        TypeLabel = "Decimal"
	TypeLabelInterface      TypeLabel = "Interface"
	TypeLabelDuration       TypeLabel = "Duration"
	TypeLabelBoolArray      TypeLabel = "BoolArray"
	TypeLabelIntArray       TypeLabel = "IntArray"
	TypeLabelInt8Array      TypeLabel = "Int8Array"
	TypeLabelInt16Array     TypeLabel = "Int16Array"
	TypeLabelInt32Array     TypeLabel = "Int32Array"
	TypeLabelInt64Array     TypeLabel = "Int64Array"
	TypeLabelUintArray      TypeLabel = "UintArray"
	TypeLabelUint8Array     TypeLabel = "Uint8Array"
	TypeLabelUint16Array    TypeLabel = "Uint16Array"
	TypeLabelUint32Array    TypeLabel = "Uint32Array"
	TypeLabelUint64Array    TypeLabel = "Uint64Array"
	TypeLabelFloat32Array   TypeLabel = "Float32Array"
	TypeLabelFloat64Array   TypeLabel = "Float64Array"
	TypeLabelStringArray    TypeLabel = "StringArray"
	TypeLabelDateTimeArray  TypeLabel = "DateTimeArray"
	TypeLabelDurationArray  TypeLabel = "DurationArray"
	TypeLabelDecimalArray   TypeLabel = "DecimalArray"
	TypeLabelInterfaceArray TypeLabel = "InterfaceArray"
)

// GetPropertyType returns the Type associated with the given label description
func GetPropertyType(label TypeLabel) Type {
	switch label {
	case TypeLabelBool:
		return Bool
	case TypeLabelInt:
		return Int
	case TypeLabelInt8:
		return Int8
	case TypeLabelInt16:
		return Int16
	case TypeLabelInt32:
		return Int32
	case TypeLabelInt64:
		return Int64
	case TypeLabelUint:
		return Uint
	case TypeLabelUint8:
		return Uint8
	case TypeLabelUint16:
		return Uint16
	case TypeLabelUint32:
		return Uint32
	case TypeLabelUint64:
		return Uint64
	case TypeLabelFloat32:
		return Float32
	case TypeLabelFloat64:
		return Float64
	case TypeLabelByteArray:
		return ByteArray
	case TypeLabelString:
		return String
	case TypeLabelDateTime:
		return DateTime
	case TypeLabelDuration:
		return Duration
	case TypeLabelDecimal:
		return Decimal
	case TypeLabelInterface:
		return Interface
	case TypeLabelBoolArray:
		return BoolArray
	case TypeLabelIntArray:
		return IntArray
	case TypeLabelInt8Array:
		return Int8Array
	case TypeLabelInt16Array:
		return Int16Array
	case TypeLabelInt32Array:
		return Int32Array
	case TypeLabelInt64Array:
		return Int64Array
	case TypeLabelUintArray:
		return UintArray
	case TypeLabelUint8Array:
		return Uint8Array
	case TypeLabelUint16Array:
		return Uint16Array
	case TypeLabelUint32Array:
		return Uint32Array
	case TypeLabelUint64Array:
		return Uint64Array
	case TypeLabelFloat32Array:
		return Float32Array
	case TypeLabelFloat64Array:
		return Float64Array
	case TypeLabelStringArray:
		return StringArray
	case TypeLabelDateTimeArray:
		return DateTimeArray
	case TypeLabelDurationArray:
		return DurationArray
	case TypeLabelDecimalArray:
		return DecimalArray
	case TypeLabelInterfaceArray:
		return InterfaceArray
	default:
		return Undefined
	}
}

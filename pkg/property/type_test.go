package property_test

import (
	"testing"

	"gitlab.com/gobl/gobl/pkg/property"
)

func TestType_String(t *testing.T) {
	tests := []struct {
		name string
		p    property.Type
		want string
	}{
		{
			name: "Undefined",
			p:    property.Undefined,
			want: property.TypeLabelUndefined,
		},
		{
			name: "Bool",
			p:    property.Bool,
			want: property.TypeLabelBool,
		},
		{
			name: "Int",
			p:    property.Int,
			want: property.TypeLabelInt,
		},
		{
			name: "Int8",
			p:    property.Int8,
			want: property.TypeLabelInt8,
		},
		{
			name: "Int16",
			p:    property.Int16,
			want: property.TypeLabelInt16,
		},
		{
			name: "Int32",
			p:    property.Int32,
			want: property.TypeLabelInt32,
		},
		{
			name: "Int64",
			p:    property.Int64,
			want: property.TypeLabelInt64,
		},
		{
			name: "Uint",
			p:    property.Uint,
			want: property.TypeLabelUint,
		},
		{
			name: "Uint8",
			p:    property.Uint8,
			want: property.TypeLabelUint8,
		},
		{
			name: "Uint16",
			p:    property.Uint16,
			want: property.TypeLabelUint16,
		},
		{
			name: "Uint32",
			p:    property.Uint32,
			want: property.TypeLabelUint32,
		},
		{
			name: "Uint64",
			p:    property.Uint64,
			want: property.TypeLabelUint64,
		},
		{
			name: "Float32",
			p:    property.Float32,
			want: property.TypeLabelFloat32,
		},
		{
			name: "Float64",
			p:    property.Float64,
			want: property.TypeLabelFloat64,
		},
		{
			name: "ByteArray",
			p:    property.ByteArray,
			want: property.TypeLabelByteArray,
		},
		{
			name: "String",
			p:    property.String,
			want: property.TypeLabelString,
		},
		{
			name: "DateTime",
			p:    property.DateTime,
			want: property.TypeLabelDateTime,
		},
		{
			name: "Decimal",
			p:    property.Decimal,
			want: property.TypeLabelDecimal,
		},
		{
			name: "Duration",
			p:    property.Duration,
			want: property.TypeLabelDuration,
		},
		{
			name: "BoolArray",
			p:    property.BoolArray,
			want: property.TypeLabelBoolArray,
		},
		{
			name: "IntArray",
			p:    property.IntArray,
			want: property.TypeLabelIntArray,
		},
		{
			name: "Int8Array",
			p:    property.Int8Array,
			want: property.TypeLabelInt8Array,
		},
		{
			name: "Int16Array",
			p:    property.Int16Array,
			want: property.TypeLabelInt16Array,
		},
		{
			name: "Int32Array",
			p:    property.Int32Array,
			want: property.TypeLabelInt32Array,
		},
		{
			name: "Int64Array",
			p:    property.Int64Array,
			want: property.TypeLabelInt64Array,
		},
		{
			name: "UintArray",
			p:    property.UintArray,
			want: property.TypeLabelUintArray,
		},
		{
			name: "Uint8Array",
			p:    property.Uint8Array,
			want: property.TypeLabelUint8Array,
		},
		{
			name: "Uint16Array",
			p:    property.Uint16Array,
			want: property.TypeLabelUint16Array,
		},
		{
			name: "Uint32Array",
			p:    property.Uint32Array,
			want: property.TypeLabelUint32Array,
		},
		{
			name: "Uint64Array",
			p:    property.Uint64Array,
			want: property.TypeLabelUint64Array,
		},
		{
			name: "Float32Array",
			p:    property.Float32Array,
			want: property.TypeLabelFloat32Array,
		},
		{
			name: "Float64Array",
			p:    property.Float64Array,
			want: property.TypeLabelFloat64Array,
		},
		{
			name: "StringArray",
			p:    property.StringArray,
			want: property.TypeLabelStringArray,
		},
		{
			name: "DateTimeArray",
			p:    property.DateTimeArray,
			want: property.TypeLabelDateTimeArray,
		},
		{
			name: "DurationArray",
			p:    property.DurationArray,
			want: property.TypeLabelDurationArray,
		},
		{
			name: "DecimalArray",
			p:    property.DecimalArray,
			want: property.TypeLabelDecimalArray,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Type.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

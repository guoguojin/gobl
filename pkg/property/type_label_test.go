package property_test

import (
	"testing"

	"gitlab.com/gobl/gobl/pkg/property"
)

func TestGetPropertyType(t *testing.T) {
	type args struct {
		label property.TypeLabel
	}
	tests := []struct {
		name string
		args args
		want property.Type
	}{
		{
			name: "Undefined",
			args: args{
				label: property.TypeLabelUndefined,
			},
			want: property.Undefined,
		},
		{
			name: "Bool",
			args: args{
				label: property.TypeLabelBool,
			},
			want: property.Bool,
		},
		{
			name: "Int",
			args: args{
				label: property.TypeLabelInt,
			},
			want: property.Int,
		},
		{
			name: "Int8",
			args: args{
				label: property.TypeLabelInt8,
			},
			want: property.Int8,
		},
		{
			name: "Int16",
			args: args{
				label: property.TypeLabelInt16,
			},
			want: property.Int16,
		},
		{
			name: "Int32",
			args: args{
				label: property.TypeLabelInt32,
			},
			want: property.Int32,
		},
		{
			name: "Int64",
			args: args{
				label: property.TypeLabelInt64,
			},
			want: property.Int64,
		},
		{
			name: "Uint",
			args: args{
				label: property.TypeLabelUint,
			},
			want: property.Uint,
		},
		{
			name: "Uint8",
			args: args{
				label: property.TypeLabelUint8,
			},
			want: property.Uint8,
		},
		{
			name: "Uint16",
			args: args{
				label: property.TypeLabelUint16,
			},
			want: property.Uint16,
		},
		{
			name: "Uint32",
			args: args{
				label: property.TypeLabelUint32,
			},
			want: property.Uint32,
		},
		{
			name: "Uint64",
			args: args{
				label: property.TypeLabelUint64,
			},
			want: property.Uint64,
		},
		{
			name: "Float32",
			args: args{
				label: property.TypeLabelFloat32,
			},
			want: property.Float32,
		},
		{
			name: "Float64",
			args: args{
				label: property.TypeLabelFloat64,
			},
			want: property.Float64,
		},
		{
			name: "ByteArray",
			args: args{
				label: property.TypeLabelByteArray,
			},
			want: property.ByteArray,
		},
		{
			name: "String",
			args: args{
				label: property.TypeLabelString,
			},
			want: property.String,
		},
		{
			name: "DateTime",
			args: args{
				label: property.TypeLabelDateTime,
			},
			want: property.DateTime,
		},
		{
			name: "Duration",
			args: args{
				label: property.TypeLabelDuration,
			},
			want: property.Duration,
		},
		{
			name: "Decimal",
			args: args{
				label: property.TypeLabelDecimal,
			},
			want: property.Decimal,
		},
		{
			name: "BoolArray",
			args: args{
				label: property.TypeLabelBoolArray,
			},
			want: property.BoolArray,
		},
		{
			name: "IntArray",
			args: args{
				label: property.TypeLabelIntArray,
			},
			want: property.IntArray,
		},
		{
			name: "Int8Array",
			args: args{
				label: property.TypeLabelInt8Array,
			},
			want: property.Int8Array,
		},
		{
			name: "Int16Array",
			args: args{
				label: property.TypeLabelInt16Array,
			},
			want: property.Int16Array,
		},
		{
			name: "Int32Array",
			args: args{
				label: property.TypeLabelInt32Array,
			},
			want: property.Int32Array,
		},
		{
			name: "Int64Array",
			args: args{
				label: property.TypeLabelInt64Array,
			},
			want: property.Int64Array,
		},
		{
			name: "UintArray",
			args: args{
				label: property.TypeLabelUintArray,
			},
			want: property.UintArray,
		},
		{
			name: "Uint8Array",
			args: args{
				label: property.TypeLabelUint8Array,
			},
			want: property.Uint8Array,
		},
		{
			name: "Uint16Array",
			args: args{
				label: property.TypeLabelUint16Array,
			},
			want: property.Uint16Array,
		},
		{
			name: "Uint32Array",
			args: args{
				label: property.TypeLabelUint32Array,
			},
			want: property.Uint32Array,
		},
		{
			name: "Uint64Array",
			args: args{
				label: property.TypeLabelUint64Array,
			},
			want: property.Uint64Array,
		},
		{
			name: "Float32Array",
			args: args{
				label: property.TypeLabelFloat32Array,
			},
			want: property.Float32Array,
		},
		{
			name: "Float64Array",
			args: args{
				label: property.TypeLabelFloat64Array,
			},
			want: property.Float64Array,
		},
		{
			name: "StringArray",
			args: args{
				label: property.TypeLabelStringArray,
			},
			want: property.StringArray,
		},
		{
			name: "DateTimeArray",
			args: args{
				label: property.TypeLabelDateTimeArray,
			},
			want: property.DateTimeArray,
		},
		{
			name: "DurationArray",
			args: args{
				label: property.TypeLabelDurationArray,
			},
			want: property.DurationArray,
		},
		{
			name: "DecimalArray",
			args: args{
				label: property.TypeLabelDecimalArray,
			},
			want: property.DecimalArray,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := property.GetPropertyType(tt.args.label); got != tt.want {
				t.Errorf("GetPropertyType() = %v, want %v", got, tt.want)
			}
		})
	}
}

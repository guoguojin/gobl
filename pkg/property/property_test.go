package property

import (
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var dt1 = time.Date(2020, 1, 1, 9, 0, 0, 0, time.UTC)
var dt2 = time.Date(2020, 6, 30, 8, 0, 0, 0, time.UTC)
var dt3 = time.Date(2019, 12, 25, 12, 37, 0, 0, time.UTC)

var dc1 = decimal.NewFromFloat(1.23456789)
var dc2 = decimal.NewFromFloat(12345.67890)
var dc3 = decimal.NewFromFloat(9876543210.123456789)
var dc4 = decimal.NewFromInt(math.MaxInt8)
var dc5 = decimal.NewFromInt(math.MaxInt16)
var dc6 = decimal.NewFromInt(math.MaxInt32)
var dc7 = decimal.NewFromInt(math.MaxInt64)
var dc8 = decimal.NewFromInt(math.MinInt8)
var dc9 = decimal.NewFromInt(math.MinInt16)
var dc10 = decimal.NewFromInt(math.MinInt32)
var dc11 = decimal.NewFromInt(math.MinInt64)

var f32v01 float32 = 1
var f32v02 float32 = -1
var f32v03 float32 = math.MaxInt8
var f32v04 float32 = math.MinInt8
var f32v05 float32 = math.MaxInt16
var f32v06 float32 = math.MinInt16
var f32v07 float32 = math.MaxInt32
var f32v08 float32 = math.MinInt32
var f32v09 float32 = math.MaxInt64
var f32v10 float32 = math.MinInt64
var f32v11 float32 = math.MaxUint8
var f32v12 float32 = math.MaxUint16
var f32v13 float32 = math.MaxUint32
var f32v14 float32 = math.MaxUint64

var f64v01 float64 = 1
var f64v02 float64 = -1
var f64v03 float64 = math.MaxInt8
var f64v04 float64 = math.MinInt8
var f64v05 float64 = math.MaxInt16
var f64v06 float64 = math.MinInt16
var f64v07 float64 = math.MaxInt32
var f64v08 float64 = math.MinInt32
var f64v09 float64 = math.MaxInt64
var f64v10 float64 = math.MinInt64
var f64v11 float64 = math.MaxUint8
var f64v12 float64 = math.MaxUint16
var f64v13 float64 = math.MaxUint32
var f64v14 float64 = math.MaxUint64
var f64v15 float64 = math.MaxUint8 * -1
var f64v16 float64 = math.MaxUint16 * -1
var f64v17 float64 = math.MaxUint32 * -1
var f64v18 float64 = math.MaxUint64 * -1

var i8v01 int8 = 1
var i8v02 int8 = -1
var i8v03 int8 = math.MaxInt8
var i8v04 int8 = math.MinInt8

var u8v01 uint8 = 1
var u8v02 uint8 = math.MaxInt8
var u8v03 uint8 = math.MaxUint8

var i16v01 int16 = 1
var i16v02 int16 = -1
var i16v03 int16 = math.MaxInt8
var i16v04 int16 = math.MinInt8
var i16v05 int16 = math.MaxUint8
var i16v06 int16 = math.MaxInt16
var i16v07 int16 = math.MinInt16

var u16v01 uint16 = 1
var u16v02 uint16 = math.MaxInt8
var u16v03 uint16 = math.MaxUint8
var u16v04 uint16 = math.MaxInt16
var u16v05 uint16 = math.MaxUint16

var i32v01 int32 = 1
var i32v02 int32 = -1
var i32v03 int32 = math.MaxInt8
var i32v04 int32 = math.MinInt8
var i32v05 int32 = math.MaxUint8
var i32v06 int32 = math.MaxInt16
var i32v07 int32 = math.MinInt16
var i32v08 int32 = math.MaxUint16
var i32v09 int32 = math.MaxInt32
var i32v10 int32 = math.MinInt32

var u32v01 uint32 = 1
var u32v02 uint32 = math.MaxInt8
var u32v03 uint32 = math.MaxUint8
var u32v04 uint32 = math.MaxInt16
var u32v05 uint32 = math.MaxUint16
var u32v06 uint32 = math.MaxInt32
var u32v07 uint32 = math.MaxUint32

var i64v01 int64 = 1
var i64v02 int64 = -1
var i64v03 int64 = math.MaxInt8
var i64v04 int64 = math.MinInt8
var i64v05 int64 = math.MaxUint8
var i64v06 int64 = math.MaxInt16
var i64v07 int64 = math.MinInt16
var i64v08 int64 = math.MaxUint16
var i64v09 int64 = math.MaxInt32
var i64v10 int64 = math.MinInt32
var i64v11 int64 = math.MaxUint32
var i64v12 int64 = math.MaxInt64
var i64v13 int64 = math.MinInt64

var u64v01 uint64 = 1
var u64v02 uint64 = math.MaxInt8
var u64v03 uint64 = math.MaxUint8
var u64v04 uint64 = math.MaxInt16
var u64v05 uint64 = math.MaxUint16
var u64v06 uint64 = math.MaxInt32
var u64v07 uint64 = math.MaxUint32
var u64v08 uint64 = math.MaxInt64
var u64v09 uint64 = math.MaxUint64

var i01 = 1
var i02 = -1
var i03 = math.MaxInt8
var i04 = math.MinInt8
var i05 = math.MaxUint8
var i06 = math.MaxInt16
var i07 = math.MinInt16
var i08 = math.MaxUint16
var i09 = math.MaxInt32
var i10 = math.MinInt32

var u01 uint = 1
var u02 uint = math.MaxInt8
var u03 uint = math.MaxUint8
var u04 uint = math.MaxInt16
var u05 uint = math.MaxUint16
var u06 uint = math.MaxInt32
var u07 uint = math.MaxUint32

func TestBoolArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []bool
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "all true",
			args: args{"test", []bool{true, true, true}},
			want: Property{
				Name:  "test",
				Type:  BoolArray,
				Value: []bool{true, true, true},
			},
		},
		{
			name: "all false",
			args: args{"test", []bool{false, false, false}},
			want: Property{
				Name:  "test",
				Type:  BoolArray,
				Value: []bool{false, false, false},
			},
		},
		{
			name: "some true, some false",
			args: args{"test", []bool{true, false, false, true, false}},
			want: Property{
				Name:  "test",
				Type:  BoolArray,
				Value: []bool{true, false, false, true, false},
			},
		},
		{
			name: "nil",
			args: args{"test", nil},
			want: Property{
				Name:  "test",
				Type:  BoolArray,
				Value: []bool(nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BoolArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoolProperty(t *testing.T) {
	type args struct {
		name  string
		value bool
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "true",
			args: args{
				name:  "test",
				value: true,
			},
			want: Property{
				Name:  "test",
				Type:  Bool,
				Value: true,
			},
		},
		{
			name: "false",
			args: args{
				name:  "test",
				value: false,
			},
			want: Property{
				Name:  "test",
				Type:  Bool,
				Value: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BoolProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestByteArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []byte
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "test",
			args: args{
				name:  "test",
				value: []byte("test"),
			},
			want: Property{
				Name:  "test",
				Type:  ByteArray,
				Value: []byte("test"),
			},
		},
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  ByteArray,
				Value: []byte(nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteArrayProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDateTimeArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []time.Time
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  DateTimeArray,
				Value: []time.Time(nil),
			},
		},
		{
			name: "Empty array",
			args: args{
				name:  "test",
				value: []time.Time{},
			},
			want: Property{
				Name:  "test",
				Type:  DateTimeArray,
				Value: []time.Time{},
			},
		},
		{
			name: "one date",
			args: args{
				name:  "test",
				value: []time.Time{dt1},
			},
			want: Property{
				Name:  "test",
				Type:  DateTimeArray,
				Value: []time.Time{dt1},
			},
		},
		{
			name: "multiple dates",
			args: args{
				name:  "test",
				value: []time.Time{dt1, dt2, dt3},
			},
			want: Property{
				Name:  "test",
				Type:  DateTimeArray,
				Value: []time.Time{dt1, dt2, dt3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DateTimeArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDateTimeProperty(t *testing.T) {
	type args struct {
		name  string
		value time.Time
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "date 1",
			args: args{
				name:  "dt1",
				value: dt1,
			},
			want: Property{
				Name:  "dt1",
				Type:  DateTime,
				Value: dt1,
			},
		},
		{
			name: "date 2",
			args: args{
				name:  "dt2",
				value: dt2,
			},
			want: Property{
				Name:  "dt2",
				Type:  DateTime,
				Value: dt2,
			},
		},
		{
			name: "date 3",
			args: args{
				name:  "dt3",
				value: dt3,
			},
			want: Property{
				Name:  "dt3",
				Type:  DateTime,
				Value: dt3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DateTimeProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDurationProperty(t *testing.T) {
	type args struct {
		name  string
		value time.Duration
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "1 minute",
			args: args{
				name:  "dur1",
				value: time.Minute,
			},
			want: Property{
				Name:  "dur1",
				Type:  Duration,
				Value: time.Minute,
			},
		},
		{
			name: "1 hour",
			args: args{
				name:  "dur2",
				value: time.Hour,
			},
			want: Property{
				Name:  "dur2",
				Type:  Duration,
				Value: time.Hour,
			},
		},
		{
			name: "1 day",
			args: args{
				name:  "dur3",
				value: time.Hour * 24,
			},
			want: Property{
				Name:  "dur3",
				Type:  Duration,
				Value: time.Hour * 24,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DurationProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDurationArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []time.Duration
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "dur1",
				value: nil,
			},
			want: Property{
				Name:  "dur1",
				Type:  DurationArray,
				Value: []time.Duration(nil),
			},
		},
		{
			name: "Empty array",
			args: args{
				name:  "dur2",
				value: []time.Duration{},
			},
			want: Property{
				Name:  "dur2",
				Type:  DurationArray,
				Value: []time.Duration{},
			},
		},
		{
			name: "1 duration",
			args: args{
				name:  "dur3",
				value: []time.Duration{time.Hour},
			},
			want: Property{
				Name:  "dur3",
				Type:  DurationArray,
				Value: []time.Duration{time.Hour},
			},
		},
		{
			name: "Multiple durations",
			args: args{
				name:  "dur4",
				value: []time.Duration{time.Minute, time.Hour},
			},
			want: Property{
				Name:  "dur4",
				Type:  DurationArray,
				Value: []time.Duration{time.Minute, time.Hour},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DurationArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDecimalArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  DecimalArray,
				Value: []decimal.Decimal(nil),
			},
		},
		{
			name: "no items",
			args: args{
				name:  "test",
				value: []decimal.Decimal{},
			},
			want: Property{
				Name:  "test",
				Type:  DecimalArray,
				Value: []decimal.Decimal{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []decimal.Decimal{dc1},
			},
			want: Property{
				Name:  "test",
				Type:  DecimalArray,
				Value: []decimal.Decimal{dc1},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []decimal.Decimal{dc1, dc2, dc3, dc4, dc5, dc6, dc7, dc8, dc9, dc10, dc11},
			},
			want: Property{
				Name:  "test",
				Type:  DecimalArray,
				Value: []decimal.Decimal{dc1, dc2, dc3, dc4, dc5, dc6, dc7, dc8, dc9, dc10, dc11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecimalArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDecimalProperty(t *testing.T) {
	type args struct {
		name  string
		value decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "decimal 1",
			args: args{
				name:  "test",
				value: dc1,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc1,
			},
		},
		{
			name: "decimal 2",
			args: args{
				name:  "test",
				value: dc2,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc2,
			},
		},
		{
			name: "decimal 3",
			args: args{
				name:  "test",
				value: dc3,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc3,
			},
		},
		{
			name: "decimal 4",
			args: args{
				name:  "test",
				value: dc4,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc4,
			},
		},
		{
			name: "decimal 5",
			args: args{
				name:  "test",
				value: dc5,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc5,
			},
		},
		{
			name: "decimal 6",
			args: args{
				name:  "test",
				value: dc6,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc6,
			},
		},
		{
			name: "decimal 7",
			args: args{
				name:  "test",
				value: dc7,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc7,
			},
		},
		{
			name: "decimal 8",
			args: args{
				name:  "test",
				value: dc8,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc8,
			},
		},
		{
			name: "decimal 9",
			args: args{
				name:  "test",
				value: dc9,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc9,
			},
		},
		{
			name: "decimal 10",
			args: args{
				name:  "test",
				value: dc10,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc10,
			},
		},
		{
			name: "decimal 11",
			args: args{
				name:  "test",
				value: dc11,
			},
			want: Property{
				Name:  "test",
				Type:  Decimal,
				Value: dc11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecimalProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFloat32ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []float32
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Float32Array,
				Value: []float32(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []float32{},
			},
			want: Property{
				Name:  "test",
				Type:  Float32Array,
				Value: []float32{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []float32{f32v01},
			},
			want: Property{
				Name:  "test",
				Type:  Float32Array,
				Value: []float32{f32v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []float32{f32v01, f32v02, f32v03, f32v04, f32v05, f32v06, f32v07, f32v08, f32v09, f32v10, f32v11, f32v12, f32v13, f32v14},
			},
			want: Property{
				Name:  "test",
				Type:  Float32Array,
				Value: []float32{f32v01, f32v02, f32v03, f32v04, f32v05, f32v06, f32v07, f32v08, f32v09, f32v10, f32v11, f32v12, f32v13, f32v14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFloat32Property(t *testing.T) {
	type args struct {
		name  string
		value float32
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "f32v01",
			args: args{
				name:  "test",
				value: f32v01,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v01,
			},
		},
		{
			name: "f32v02",
			args: args{
				name:  "test",
				value: f32v02,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v02,
			},
		},
		{
			name: "f32v03",
			args: args{
				name:  "test",
				value: f32v03,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v03,
			},
		},
		{
			name: "f32v04",
			args: args{
				name:  "test",
				value: f32v04,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v04,
			},
		},
		{
			name: "f32v05",
			args: args{
				name:  "test",
				value: f32v05,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v05,
			},
		},
		{
			name: "f32v06",
			args: args{
				name:  "test",
				value: f32v06,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v06,
			},
		},
		{
			name: "f32v07",
			args: args{
				name:  "test",
				value: f32v07,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v07,
			},
		},
		{
			name: "f32v08",
			args: args{
				name:  "test",
				value: f32v08,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v08,
			},
		},
		{
			name: "f32v09",
			args: args{
				name:  "test",
				value: f32v09,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v09,
			},
		},
		{
			name: "f32v10",
			args: args{
				name:  "test",
				value: f32v10,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v10,
			},
		},
		{
			name: "f32v11",
			args: args{
				name:  "test",
				value: f32v11,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v11,
			},
		},
		{
			name: "f32v12",
			args: args{
				name:  "test",
				value: f32v12,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v12,
			},
		},
		{
			name: "f32v13",
			args: args{
				name:  "test",
				value: f32v13,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v13,
			},
		},
		{
			name: "f32v14",
			args: args{
				name:  "test",
				value: f32v14,
			},
			want: Property{
				Name:  "test",
				Type:  Float32,
				Value: f32v14,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFloat64ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []float64
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Float64Array,
				Value: []float64(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []float64{},
			},
			want: Property{
				Name:  "test",
				Type:  Float64Array,
				Value: []float64{},
			},
		},
		{
			name: "one value",
			args: args{
				name:  "test",
				value: []float64{f64v01},
			},
			want: Property{
				Name:  "test",
				Type:  Float64Array,
				Value: []float64{f64v01},
			},
		},
		{
			name: "multiple values",
			args: args{
				name:  "test",
				value: []float64{f64v01, f64v02, f64v03, f64v04, f64v05, f64v06, f64v07, f64v08, f64v09, f64v10, f64v11, f64v12, f64v13, f64v14, f64v15, f64v16, f64v17, f64v18},
			},
			want: Property{
				Name:  "test",
				Type:  Float64Array,
				Value: []float64{f64v01, f64v02, f64v03, f64v04, f64v05, f64v06, f64v07, f64v08, f64v09, f64v10, f64v11, f64v12, f64v13, f64v14, f64v15, f64v16, f64v17, f64v18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFloat64Property(t *testing.T) {
	type args struct {
		name  string
		value float64
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "f64v01",
			args: args{
				name:  "test",
				value: f64v01,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v01,
			},
		},
		{
			name: "f64v02",
			args: args{
				name:  "test",
				value: f64v02,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v02,
			},
		},
		{
			name: "f64v03",
			args: args{
				name:  "test",
				value: f64v03,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v03,
			},
		},
		{
			name: "f64v04",
			args: args{
				name:  "test",
				value: f64v04,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v04,
			},
		},
		{
			name: "f64v05",
			args: args{
				name:  "test",
				value: f64v05,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v05,
			},
		},
		{
			name: "f64v06",
			args: args{
				name:  "test",
				value: f64v06,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v06,
			},
		},
		{
			name: "f64v07",
			args: args{
				name:  "test",
				value: f64v07,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v07,
			},
		},
		{
			name: "f64v08",
			args: args{
				name:  "test",
				value: f64v08,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v08,
			},
		},
		{
			name: "f64v09",
			args: args{
				name:  "test",
				value: f64v09,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v09,
			},
		},
		{
			name: "f64v10",
			args: args{
				name:  "test",
				value: f64v10,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v10,
			},
		},
		{
			name: "f64v11",
			args: args{
				name:  "test",
				value: f64v11,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v11,
			},
		},
		{
			name: "f64v12",
			args: args{
				name:  "test",
				value: f64v12,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v12,
			},
		},
		{
			name: "f64v13",
			args: args{
				name:  "test",
				value: f64v13,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v13,
			},
		},
		{
			name: "f64v14",
			args: args{
				name:  "test",
				value: f64v14,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v14,
			},
		},
		{
			name: "f64v15",
			args: args{
				name:  "test",
				value: f64v15,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v15,
			},
		},
		{
			name: "f64v16",
			args: args{
				name:  "test",
				value: f64v16,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v16,
			},
		},
		{
			name: "f64v15",
			args: args{
				name:  "test",
				value: f64v15,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v15,
			},
		},
		{
			name: "f64v16",
			args: args{
				name:  "test",
				value: f64v16,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v16,
			},
		},
		{
			name: "f64v17",
			args: args{
				name:  "test",
				value: f64v17,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v17,
			},
		},
		{
			name: "f64v18",
			args: args{
				name:  "test",
				value: f64v18,
			},
			want: Property{
				Name:  "test",
				Type:  Float64,
				Value: f64v18,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt16ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []int16
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Int16Array,
				Value: []int16(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []int16{},
			},
			want: Property{
				Name:  "test",
				Type:  Int16Array,
				Value: []int16{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []int16{i16v01},
			},
			want: Property{
				Name:  "test",
				Type:  Int16Array,
				Value: []int16{i16v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []int16{i16v01, i16v02, i16v03, i16v04, i16v05, i16v06, i16v07},
			},
			want: Property{
				Name:  "test",
				Type:  Int16Array,
				Value: []int16{i16v01, i16v02, i16v03, i16v04, i16v05, i16v06, i16v07},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt16Property(t *testing.T) {
	type args struct {
		name  string
		value int16
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "i16v01",
			args: args{
				name:  "test",
				value: i16v01,
			},
			want: Property{
				Name:  "test",
				Type:  Int16,
				Value: i16v01,
			},
		},
		{
			name: "i16v02",
			args: args{
				name:  "test",
				value: i16v02,
			},
			want: Property{
				Name:  "test",
				Type:  Int16,
				Value: i16v02,
			},
		},
		{
			name: "i16v03",
			args: args{
				name:  "test",
				value: i16v03,
			},
			want: Property{
				Name:  "test",
				Type:  Int16,
				Value: i16v03,
			},
		},
		{
			name: "i16v04",
			args: args{
				name:  "test",
				value: i16v04,
			},
			want: Property{
				Name:  "test",
				Type:  Int16,
				Value: i16v04,
			},
		},
		{
			name: "i16v05",
			args: args{
				name:  "test",
				value: i16v05,
			},
			want: Property{
				Name:  "test",
				Type:  Int16,
				Value: i16v05,
			},
		},
		{
			name: "i16v06",
			args: args{
				name:  "test",
				value: i16v06,
			},
			want: Property{
				Name:  "test",
				Type:  Int16,
				Value: i16v06,
			},
		},
		{
			name: "i16v07",
			args: args{
				name:  "test",
				value: i16v07,
			},
			want: Property{
				Name:  "test",
				Type:  Int16,
				Value: i16v07,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt32ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []int32
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Int32Array,
				Value: []int32(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []int32{},
			},
			want: Property{
				Name:  "test",
				Type:  Int32Array,
				Value: []int32{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []int32{i32v01},
			},
			want: Property{
				Name:  "test",
				Type:  Int32Array,
				Value: []int32{i32v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []int32{i32v01, i32v02, i32v03, i32v04, i32v05, i32v06, i32v07, i32v08, i32v09, i32v10},
			},
			want: Property{
				Name:  "test",
				Type:  Int32Array,
				Value: []int32{i32v01, i32v02, i32v03, i32v04, i32v05, i32v06, i32v07, i32v08, i32v09, i32v10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt32Property(t *testing.T) {
	type args struct {
		name  string
		value int32
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "i32v01",
			args: args{
				name:  "test",
				value: i32v01,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v01,
			},
		},
		{
			name: "i32v02",
			args: args{
				name:  "test",
				value: i32v02,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v02,
			},
		},
		{
			name: "i32v03",
			args: args{
				name:  "test",
				value: i32v03,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v03,
			},
		},
		{
			name: "i32v04",
			args: args{
				name:  "test",
				value: i32v04,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v04,
			},
		},
		{
			name: "i32v05",
			args: args{
				name:  "test",
				value: i32v05,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v05,
			},
		},
		{
			name: "i32v06",
			args: args{
				name:  "test",
				value: i32v06,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v06,
			},
		},
		{
			name: "i32v07",
			args: args{
				name:  "test",
				value: i32v07,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v07,
			},
		},
		{
			name: "i32v08",
			args: args{
				name:  "test",
				value: i32v08,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v08,
			},
		},
		{
			name: "i32v09",
			args: args{
				name:  "test",
				value: i32v09,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v09,
			},
		},
		{
			name: "i32v10",
			args: args{
				name:  "test",
				value: i32v10,
			},
			want: Property{
				Name:  "test",
				Type:  Int32,
				Value: i32v10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt64ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []int64
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Int64Array,
				Value: []int64(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []int64{},
			},
			want: Property{
				Name:  "test",
				Type:  Int64Array,
				Value: []int64{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []int64{i64v01},
			},
			want: Property{
				Name:  "test",
				Type:  Int64Array,
				Value: []int64{i64v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []int64{i64v01, i64v02, i64v03, i64v04, i64v05, i64v06, i64v07, i64v08, i64v09, i64v10, i64v11, i64v12, i64v13},
			},
			want: Property{
				Name:  "test",
				Type:  Int64Array,
				Value: []int64{i64v01, i64v02, i64v03, i64v04, i64v05, i64v06, i64v07, i64v08, i64v09, i64v10, i64v11, i64v12, i64v13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt64Property(t *testing.T) {
	type args struct {
		name  string
		value int64
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "i64v01",
			args: args{
				name:  "test",
				value: i64v01,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v01,
			},
		},
		{
			name: "i64v02",
			args: args{
				name:  "test",
				value: i64v02,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v02,
			},
		},
		{
			name: "i64v03",
			args: args{
				name:  "test",
				value: i64v03,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v03,
			},
		},
		{
			name: "i64v04",
			args: args{
				name:  "test",
				value: i64v04,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v04,
			},
		},
		{
			name: "i64v05",
			args: args{
				name:  "test",
				value: i64v05,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v05,
			},
		},
		{
			name: "i64v06",
			args: args{
				name:  "test",
				value: i64v06,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v06,
			},
		},
		{
			name: "i64v07",
			args: args{
				name:  "test",
				value: i64v07,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v07,
			},
		},
		{
			name: "i64v08",
			args: args{
				name:  "test",
				value: i64v08,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v08,
			},
		},
		{
			name: "i64v09",
			args: args{
				name:  "test",
				value: i64v09,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v09,
			},
		},
		{
			name: "i64v10",
			args: args{
				name:  "test",
				value: i64v10,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v10,
			},
		},
		{
			name: "i64v11",
			args: args{
				name:  "test",
				value: i64v11,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v11,
			},
		},
		{
			name: "i64v12",
			args: args{
				name:  "test",
				value: i64v12,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v12,
			},
		},
		{
			name: "i64v13",
			args: args{
				name:  "test",
				value: i64v13,
			},
			want: Property{
				Name:  "test",
				Type:  Int64,
				Value: i64v13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt8ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []int8
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Int8Array,
				Value: []int8(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []int8{},
			},
			want: Property{
				Name:  "test",
				Type:  Int8Array,
				Value: []int8{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []int8{i8v01},
			},
			want: Property{
				Name:  "test",
				Type:  Int8Array,
				Value: []int8{i8v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []int8{i8v01, i8v02, i8v03, i8v04},
			},
			want: Property{
				Name:  "test",
				Type:  Int8Array,
				Value: []int8{i8v01, i8v02, i8v03, i8v04},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestInt8Property(t *testing.T) {
	type args struct {
		name  string
		value int8
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "i8v01",
			args: args{
				name:  "test",
				value: i8v01,
			},
			want: Property{
				Name:  "test",
				Type:  Int8,
				Value: i8v01,
			},
		},
		{
			name: "i8v02",
			args: args{
				name:  "test",
				value: i8v02,
			},
			want: Property{
				Name:  "test",
				Type:  Int8,
				Value: i8v02,
			},
		},
		{
			name: "i8v03",
			args: args{
				name:  "test",
				value: i8v03,
			},
			want: Property{
				Name:  "test",
				Type:  Int8,
				Value: i8v03,
			},
		},
		{
			name: "i8v04",
			args: args{
				name:  "test",
				value: i8v04,
			},
			want: Property{
				Name:  "test",
				Type:  Int8,
				Value: i8v04,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestIntArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []int
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  IntArray,
				Value: []int(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []int{},
			},
			want: Property{
				Name:  "test",
				Type:  IntArray,
				Value: []int{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []int{i01},
			},
			want: Property{
				Name:  "test",
				Type:  IntArray,
				Value: []int{i01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []int{i01, i02, i03, i04, i05, i06, i07, i08, i09, i10},
			},
			want: Property{
				Name:  "test",
				Type:  IntArray,
				Value: []int{i01, i02, i03, i04, i05, i06, i07, i08, i09, i10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestIntProperty(t *testing.T) {
	type args struct {
		name  string
		value int
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "i01",
			args: args{
				name:  "test",
				value: i01,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i01,
			},
		},
		{
			name: "i02",
			args: args{
				name:  "test",
				value: i02,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i02,
			},
		},
		{
			name: "i03",
			args: args{
				name:  "test",
				value: i03,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i03,
			},
		},
		{
			name: "i04",
			args: args{
				name:  "test",
				value: i04,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i04,
			},
		},
		{
			name: "i05",
			args: args{
				name:  "test",
				value: i05,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i05,
			},
		},
		{
			name: "i06",
			args: args{
				name:  "test",
				value: i06,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i06,
			},
		},
		{
			name: "i07",
			args: args{
				name:  "test",
				value: i07,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i07,
			},
		},
		{
			name: "i08",
			args: args{
				name:  "test",
				value: i08,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i08,
			},
		},
		{
			name: "i09",
			args: args{
				name:  "test",
				value: i09,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i09,
			},
		},
		{
			name: "i10",
			args: args{
				name:  "test",
				value: i10,
			},
			want: Property{
				Name:  "test",
				Type:  Int,
				Value: i10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_BoolArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: BoolArray,
				value:        []bool{true, true, true},
			},
			args: args{
				d: []bool{true, false, true},
			},
			want: []bool{true, true, true},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []bool{true, false, true},
			},
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.BoolArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_BoolArray(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []bool
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: BoolArray,
				value:        []bool{true},
			},
			want:    []bool{true},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        100,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.BoolArray()
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Bool(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        100,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Bool()
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_BoolOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		b bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				b: false,
			},
			want: true,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        100,
			},
			args: args{
				b: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.BoolOrDefault(tt.args.b)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_DateTime(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    time.Time
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        time.Unix(100, 0),
			},
			want:    time.Unix(100, 0),
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    time.Unix(0, 0),
			wantErr: true,
		},
		{
			name: "invalid property value",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        "not a date",
			},
			want:    time.Unix(0, 0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.DateTime()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Duration(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    time.Duration
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Duration,
				value:        time.Minute,
			},
			want:    time.Minute,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    time.Duration(0),
			wantErr: true,
		},
		{
			name: "invalid property value",
			fields: fields{
				name:         "test",
				propertyType: Duration,
				value:        "not a date",
			},
			want:    time.Duration(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Duration()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_DateTimeArray(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []time.Time
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DateTimeArray,
				value:        []time.Time{time.Unix(1, 0), time.Unix(2, 0)},
			},
			want:    []time.Time{time.Unix(1, 0), time.Unix(2, 0)},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        100,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.DateTimeArray()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_DurationArray(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []time.Duration
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DurationArray,
				value:        []time.Duration{time.Minute, time.Hour},
			},
			want:    []time.Duration{time.Minute, time.Hour},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid property value",
			fields: fields{
				name:         "test",
				propertyType: DurationArray,
				value:        []string{"not a duration"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.DurationArray()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_DateTimeArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []time.Time
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DateTimeArray,
				value:        []time.Time{time.Unix(1, 0), time.Unix(2, 0)},
			},
			args: args{
				d: []time.Time{time.Unix(100, 0), time.Unix(101, 0)},
			},
			want: []time.Time{time.Unix(1, 0), time.Unix(2, 0)},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []time.Time{time.Unix(100, 0), time.Unix(101, 0)},
			},
			want: []time.Time{time.Unix(100, 0), time.Unix(101, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.DateTimeArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_DurationArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []time.Duration
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DurationArray,
				value:        []time.Duration{time.Minute, time.Second},
			},
			args: args{
				d: []time.Duration{time.Hour, time.Minute},
			},
			want: []time.Duration{time.Minute, time.Second},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []time.Duration{time.Hour, time.Minute},
			},
			want: []time.Duration{time.Hour, time.Minute},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.DurationArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_DateTimeOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        time.Unix(1, 0),
			},
			args: args{
				d: time.Unix(100, 0),
			},
			want: time.Unix(1, 0),
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: time.Unix(100, 0),
			},
			want: time.Unix(100, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.DateTimeOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_DurationOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Duration,
				value:        time.Second,
			},
			args: args{
				d: time.Minute,
			},
			want: time.Second,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: time.Minute,
			},
			want: time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.DurationOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Decimal(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    decimal.Decimal
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        decimal.New(1, 1),
			},
			want:    decimal.New(1, 1),
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        false,
			},
			want:    decimal.Zero,
			wantErr: true,
		},
		{
			name: "invalid property value",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        "not a decimal",
			},
			want:    decimal.Zero,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Decimal()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_DecimalArray(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []decimal.Decimal
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DecimalArray,
				value:        []decimal.Decimal{decimal.New(1, 1), decimal.New(2, 1), decimal.New(3, 1)},
			},
			want:    []decimal.Decimal{decimal.New(1, 1), decimal.New(2, 1), decimal.New(3, 1)},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.DecimalArray()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_DecimalArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []decimal.Decimal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []decimal.Decimal
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: DecimalArray,
				value:        []decimal.Decimal{decimal.New(1, 1), decimal.New(2, 1), decimal.New(3, 1)},
			},
			args: args{
				d: []decimal.Decimal{decimal.New(1, 2)},
			},
			want: []decimal.Decimal{decimal.New(1, 1), decimal.New(2, 1), decimal.New(3, 1)},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []decimal.Decimal{decimal.New(1, 2), decimal.New(2, 2), decimal.New(3, 2)},
			},
			want: []decimal.Decimal{decimal.New(1, 2), decimal.New(2, 2), decimal.New(3, 2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.DecimalArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_DecimalOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d decimal.Decimal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   decimal.Decimal
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        decimal.New(1, 1),
			},
			args: args{
				d: decimal.New(1, 2),
			},
			want: decimal.New(1, 1),
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: decimal.New(1, 2),
			},
			want: decimal.New(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.DecimalOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Float32(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    float32
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float32,
				value:        float32(100),
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        false,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Float32()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Float32Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []float32
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float32Array,
				value:        []float32{10, 20, 30},
			},
			want:    []float32{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        false,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Float32Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Float32ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []float32
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float32Array,
				value:        []float32{10, 20, 30},
			},
			args: args{
				d: []float32{100, 200, 300},
			},
			want: []float32{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []float32{100, 200, 300},
			},
			want: []float32{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Float32ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Float32OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float32,
				value:        float32(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Float32OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Float64(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float64,
				value:        10.0,
			},
			want:    10.0,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Float64()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Float64Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []float64
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float64Array,
				value:        []float64{10, 20, 30},
			},
			want:    []float64{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Float64Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Float64ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []float64
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float64Array,
				value:        []float64{10, 20, 30},
			},
			args: args{
				d: []float64{100, 200, 300},
			},
			want: []float64{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []float64{100, 200, 300},
			},
			want: []float64{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Float64ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Float64OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Float64,
				value:        10.0,
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Float64OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        100,
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int16(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    int16
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int16,
				value:        int16(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int16()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int16Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []int16
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int16Array,
				value:        []int16{10, 20, 30},
			},
			want:    []int16{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int16Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int16ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []int16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int16
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int16Array,
				value:        []int16{10, 20, 30},
			},
			args: args{
				d: []int16{100, 200, 300},
			},
			want: []int16{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []int16{100, 200, 300},
			},
			want: []int16{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int16ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int16OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d int16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int16
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int16,
				value:        int16(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int16OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int32(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    int32
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int32,
				value:        int32(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int32()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int32Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []int32
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int32Array,
				value:        []int32{10, 20, 30},
			},
			want:    []int32{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int32Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int32ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int32
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int32Array,
				value:        []int32{10, 20, 30},
			},
			args: args{
				d: []int32{100, 200, 300},
			},
			want: []int32{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int32ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int32OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int32,
				value:        int32(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int32OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int64(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int64,
				value:        int64(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int64()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int64Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []int64
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int64Array,
				value:        []int64{10, 20, 30},
			},
			want:    []int64{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int64Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int64ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int64
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int64Array,
				value:        []int64{10, 20, 30},
			},
			args: args{
				d: []int64{100, 200, 300},
			},
			want: []int64{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []int64{100, 200, 300},
			},
			want: []int64{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int64ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int64OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int64,
				value:        int64(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int64OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int8(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    int8
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int8,
				value:        int8(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int8()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int8Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []int8
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int8Array,
				value:        []int8{10, 20, 30},
			},
			want:    []int8{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Int8Array()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int8Array() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8Array() got = %v, want %v", got, tt.want)
			}
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Int8ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int8
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int8Array,
				value:        []int8{10, 20, 30},
			},
			args: args{
				d: []int8{50, 60, 70},
			},
			want: []int8{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []int8{10, 20, 30},
			},
			want: []int8{10, 20, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int8ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Int8OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int8
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int8,
				value:        int8(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Int8OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_IntArray(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []int
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: IntArray,
				value:        []int{10, 20, 30},
			},
			want:    []int{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.IntArray()
			if (err != nil) != tt.wantErr {
				t.Errorf("IntArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntArray() got = %v, want %v", got, tt.want)
			}
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_IntArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: IntArray,
				value:        []int{10, 20, 30},
			},
			args: args{
				d: []int{100, 200, 300},
			},
			want: []int{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []int{100, 200, 300},
			},
			want: []int{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.IntArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_IntOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        10,
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.IntOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_String(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Bool",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want: "true",
		},
		{
			name: "Bool Nil",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        nil,
			},
			want: "false",
		},
		{
			name: "Int",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        10,
			},
			want: "10",
		},
		{
			name: "Int Nil",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Int8",
			fields: fields{
				name:         "test",
				propertyType: Int8,
				value:        int8(10),
			},
			want: "10",
		},
		{
			name: "Int8 Nil",
			fields: fields{
				name:         "test",
				propertyType: Int8,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Int16",
			fields: fields{
				name:         "test",
				propertyType: Int16,
				value:        int16(10),
			},
			want: "10",
		},
		{
			name: "Int16 Nil",
			fields: fields{
				name:         "test",
				propertyType: Int16,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Int32",
			fields: fields{
				name:         "test",
				propertyType: Int32,
				value:        int32(10),
			},
			want: "10",
		},
		{
			name: "Int32 Nil",
			fields: fields{
				name:         "test",
				propertyType: Int32,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Int64",
			fields: fields{
				name:         "test",
				propertyType: Int64,
				value:        int64(10),
			},
			want: "10",
		},
		{
			name: "Int64 Nil",
			fields: fields{
				name:         "test",
				propertyType: Int64,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Uint",
			fields: fields{
				name:         "test",
				propertyType: Uint,
				value:        uint(10),
			},
			want: "10",
		},
		{
			name: "Uint Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Uint8",
			fields: fields{
				name:         "test",
				propertyType: Uint8,
				value:        uint8(10),
			},
			want: "10",
		},
		{
			name: "Uint8 Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint8,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Uint16",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        uint16(10),
			},
			want: "10",
		},
		{
			name: "Uint16 Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Uint32",
			fields: fields{
				name:         "test",
				propertyType: Uint32,
				value:        uint32(10),
			},
			want: "10",
		},
		{
			name: "Uint32 Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint32,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Uint64",
			fields: fields{
				name:         "test",
				propertyType: Uint64,
				value:        uint64(10),
			},
			want: "10",
		},
		{
			name: "Uint64 Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint64,
				value:        nil,
			},
			want: "0",
		},
		{
			name: "Float32",
			fields: fields{
				name:         "test",
				propertyType: Float32,
				value:        float32(10),
			},
			want: "10.000000",
		},
		{
			name: "Float32 Nil",
			fields: fields{
				name:         "test",
				propertyType: Float32,
				value:        nil,
			},
			want: "0.000000",
		},
		{
			name: "Float64",
			fields: fields{
				name:         "test",
				propertyType: Float64,
				value:        10.0,
			},
			want: "10.000000",
		},
		{
			name: "Float64 Nil",
			fields: fields{
				name:         "test",
				propertyType: Float64,
				value:        nil,
			},
			want: "0.000000",
		},
		{
			name: "ByteArray",
			fields: fields{
				name:         "test",
				propertyType: ByteArray,
				value:        []byte("test"),
			},
			want: "test",
		},
		{
			name: "ByteArray Nil",
			fields: fields{
				name:         "test",
				propertyType: ByteArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "String",
			fields: fields{
				name:         "test",
				propertyType: String,
				value:        "test",
			},
			want: "test",
		},
		{
			name: "String Nil",
			fields: fields{
				name:         "test",
				propertyType: String,
				value:        nil,
			},
			want: "",
		},
		{
			name: "Invalid String Nil",
			fields: fields{
				name:         "test",
				propertyType: String,
				value:        struct{}{},
			},
			want: "",
		},
		{
			name: "DateTime",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        time.Date(1970, 1, 1, 0, 0, 0, 123000000, time.UTC),
			},
			want: "1970-01-01T00:00:00.123Z",
		},
		{
			name: "DateTime Nil",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        nil,
			},
			want: "1970-01-01T00:00:00Z",
		},
		{
			name: "Duration",
			fields: fields{
				name:         "test",
				propertyType: Duration,
				value:        time.Second,
			},
			want: "1s",
		},
		{
			name: "Duration Nil",
			fields: fields{
				name:         "test",
				propertyType: Duration,
				value:        nil,
			},
			want: "0s",
		},
		{
			name: "Decimal",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        decimal.New(1, 1),
			},
			want: "10.000000",
		},
		{
			name: "Decimal Nil",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        nil,
			},
			want: "0.000000",
		},
		{
			name: "BoolArray",
			fields: fields{
				name:         "test",
				propertyType: BoolArray,
				value:        []bool{true, false, true},
			},
			want: "[true,false,true]",
		},
		{
			name: "BoolArray Nil",
			fields: fields{
				name:         "test",
				propertyType: BoolArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "IntArray",
			fields: fields{
				name:         "test",
				propertyType: IntArray,
				value:        []int{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "IntArray Nil",
			fields: fields{
				name:         "test",
				propertyType: IntArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Int8Array",
			fields: fields{
				name:         "test",
				propertyType: Int8Array,
				value:        []int8{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Int8Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Int8Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Int16Array",
			fields: fields{
				name:         "test",
				propertyType: Int16Array,
				value:        []int16{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Int16Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Int16Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Int32Array",
			fields: fields{
				name:         "test",
				propertyType: Int32Array,
				value:        []int32{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Int32Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Int32Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Int64Array",
			fields: fields{
				name:         "test",
				propertyType: Int64Array,
				value:        []int64{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Int64Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Int64Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "UintArray",
			fields: fields{
				name:         "test",
				propertyType: UintArray,
				value:        []uint{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "UintArray Nil",
			fields: fields{
				name:         "test",
				propertyType: UintArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Uint8Array",
			fields: fields{
				name:         "test",
				propertyType: Uint8Array,
				value:        []uint8{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Uint8Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint8Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Uint16Array",
			fields: fields{
				name:         "test",
				propertyType: Uint16Array,
				value:        []uint16{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Uint16Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint16Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Uint32Array",
			fields: fields{
				name:         "test",
				propertyType: Uint32Array,
				value:        []uint32{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Uint32Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint32Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Uint64Array",
			fields: fields{
				name:         "test",
				propertyType: Uint64Array,
				value:        []uint64{10, 20, 30},
			},
			want: "[10,20,30]",
		},
		{
			name: "Uint64Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Uint64Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Float32Array",
			fields: fields{
				name:         "test",
				propertyType: Float32Array,
				value:        []float32{10, 20, 30},
			},
			want: "[10.000000,20.000000,30.000000]",
		},
		{
			name: "Float32Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Float32Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Float64Array",
			fields: fields{
				name:         "test",
				propertyType: Float64Array,
				value:        []float64{10, 20, 30},
			},
			want: "[10.000000,20.000000,30.000000]",
		},
		{
			name: "Float64Array Nil",
			fields: fields{
				name:         "test",
				propertyType: Float64Array,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "StringArray",
			fields: fields{
				name:         "test",
				propertyType: StringArray,
				value:        []string{"this", "is", "a", "test"},
			},
			want: "[\"this\",\"is\",\"a\",\"test\"]",
		},
		{
			name: "StringArray Nil",
			fields: fields{
				name:         "test",
				propertyType: StringArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "DateTimeArray",
			fields: fields{
				name:         "test",
				propertyType: DateTimeArray,
				value: []time.Time{
					time.Date(1970, 1, 1, 0, 0, 0, 123000000, time.UTC),
					time.Date(1970, 1, 1, 1, 0, 0, 123000000, time.UTC),
				},
			},
			want: "[1970-01-01T00:00:00.123Z,1970-01-01T01:00:00.123Z]",
		},
		{
			name: "DateTimeArray Nil",
			fields: fields{
				name:         "test",
				propertyType: DateTimeArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "DurationArray",
			fields: fields{
				name:         "test",
				propertyType: DurationArray,
				value: []time.Duration{
					time.Second,
					time.Minute,
				},
			},
			want: "[1s,1m0s]",
		},
		{
			name: "DurationArray Nil",
			fields: fields{
				name:         "test",
				propertyType: DurationArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "DecimalArray",
			fields: fields{
				name:         "test",
				propertyType: DecimalArray,
				value:        []decimal.Decimal{decimal.New(1, 1), decimal.New(2, 1), decimal.New(3, 1)},
			},
			want: "[10.000000,20.000000,30.000000]",
		},
		{
			name: "DecimalArray Nil",
			fields: fields{
				name:         "test",
				propertyType: DecimalArray,
				value:        nil,
			},
			want: "[]",
		},
		{
			name: "Undefined",
			fields: fields{
				name:         "test",
				propertyType: Undefined,
				value:        nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.String()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_StringArray(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "Bool",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    []string{"true"},
			wantErr: false,
		},
		{
			name: "Bool Error",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        "not a bool",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        10,
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Int Error",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        "not an int",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int8",
			fields: fields{
				name:         "test",
				propertyType: Int8,
				value:        int8(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Int8 Error",
			fields: fields{
				name:         "test",
				propertyType: Int8,
				value:        "not an int",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int16",
			fields: fields{
				name:         "test",
				propertyType: Int16,
				value:        int16(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Int16 Error",
			fields: fields{
				name:         "test",
				propertyType: Int16,
				value:        "not an int",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int32",
			fields: fields{
				name:         "test",
				propertyType: Int32,
				value:        int32(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Int32 Error",
			fields: fields{
				name:         "test",
				propertyType: Int32,
				value:        "not an int",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int64",
			fields: fields{
				name:         "test",
				propertyType: Int64,
				value:        int64(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Int64 Error",
			fields: fields{
				name:         "test",
				propertyType: Int64,
				value:        "not an int",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint",
			fields: fields{
				name:         "test",
				propertyType: Uint,
				value:        uint(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Uint Error",
			fields: fields{
				name:         "test",
				propertyType: Uint,
				value:        "not an uint",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint8",
			fields: fields{
				name:         "test",
				propertyType: Uint8,
				value:        uint8(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Uint8 Error",
			fields: fields{
				name:         "test",
				propertyType: Uint8,
				value:        "not an uint",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint16",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        uint16(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Uint16 Error",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        "not an uint",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint16 Error",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        "not an uint",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint32",
			fields: fields{
				name:         "test",
				propertyType: Uint32,
				value:        uint32(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Uint32 Error",
			fields: fields{
				name:         "test",
				propertyType: Uint32,
				value:        "not an uint",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint64",
			fields: fields{
				name:         "test",
				propertyType: Uint64,
				value:        uint64(10),
			},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name: "Uint64 Error",
			fields: fields{
				name:         "test",
				propertyType: Uint64,
				value:        "not an uint",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Float32",
			fields: fields{
				name:         "test",
				propertyType: Float32,
				value:        float32(10),
			},
			want:    []string{"10.000000"},
			wantErr: false,
		},
		{
			name: "Float32 Error",
			fields: fields{
				name:         "test",
				propertyType: Float32,
				value:        "not a float",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Float64",
			fields: fields{
				name:         "test",
				propertyType: Float64,
				value:        10.0,
			},
			want:    []string{"10.000000"},
			wantErr: false,
		},
		{
			name: "Float64 Error",
			fields: fields{
				name:         "test",
				propertyType: Float64,
				value:        "not a float",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ByteArray",
			fields: fields{
				name:         "test",
				propertyType: ByteArray,
				value:        []byte("test"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ByteArray Error",
			fields: fields{
				name:         "test",
				propertyType: ByteArray,
				value:        "not a byte array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "String",
			fields: fields{
				name:         "test",
				propertyType: String,
				value:        "test",
			},
			want:    []string{"test"},
			wantErr: false,
		},
		{
			name: "String Error",
			fields: fields{
				name:         "test",
				propertyType: String,
				value:        struct{}{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "DateTime",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        time.Date(1970, 1, 1, 0, 0, 0, 123000000, time.UTC),
			},
			want:    []string{"1970-01-01T00:00:00.123Z"},
			wantErr: false,
		},
		{
			name: "DateTime Error",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        "not a date",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Duration",
			fields: fields{
				name:         "test",
				propertyType: Duration,
				value:        time.Second,
			},
			want:    []string{"1s"},
			wantErr: false,
		},
		{
			name: "Duration Error",
			fields: fields{
				name:         "test",
				propertyType: Duration,
				value:        "not a duration",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Decimal",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        decimal.New(1, 1),
			},
			want:    []string{"10.000000"},
			wantErr: false,
		},
		{
			name: "Decimal Error",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        "not a decimal",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "BoolArray",
			fields: fields{
				name:         "test",
				propertyType: BoolArray,
				value:        []bool{true, false, true},
			},
			want:    []string{"true", "false", "true"},
			wantErr: false,
		},
		{
			name: "BoolArray Error",
			fields: fields{
				name:         "test",
				propertyType: BoolArray,
				value:        "not a bool array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "IntArray",
			fields: fields{
				name:         "test",
				propertyType: IntArray,
				value:        []int{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "IntArray Error",
			fields: fields{
				name:         "test",
				propertyType: IntArray,
				value:        "not an int array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int8Array",
			fields: fields{
				name:         "test",
				propertyType: Int8Array,
				value:        []int8{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Int8Array Error",
			fields: fields{
				name:         "test",
				propertyType: Int8Array,
				value:        "not an int array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int16Array",
			fields: fields{
				name:         "test",
				propertyType: Int16Array,
				value:        []int16{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Int16Array Error",
			fields: fields{
				name:         "test",
				propertyType: Int16Array,
				value:        "not an int array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int32Array",
			fields: fields{
				name:         "test",
				propertyType: Int32Array,
				value:        []int32{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Int32Array Error",
			fields: fields{
				name:         "test",
				propertyType: Int32Array,
				value:        "not an int array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Int64Array",
			fields: fields{
				name:         "test",
				propertyType: Int64Array,
				value:        []int64{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Int64Array Error",
			fields: fields{
				name:         "test",
				propertyType: Int64Array,
				value:        "not an int array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "UintArray",
			fields: fields{
				name:         "test",
				propertyType: UintArray,
				value:        []uint{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "UintArray Error",
			fields: fields{
				name:         "test",
				propertyType: UintArray,
				value:        "not an uint array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint8Array",
			fields: fields{
				name:         "test",
				propertyType: Uint8Array,
				value:        []uint8{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Uint8Array Error",
			fields: fields{
				name:         "test",
				propertyType: Uint8Array,
				value:        "not an uint array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint16Array",
			fields: fields{
				name:         "test",
				propertyType: Uint16Array,
				value:        []uint16{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Uint16Array Error",
			fields: fields{
				name:         "test",
				propertyType: Uint16Array,
				value:        "not an uint array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint32Array",
			fields: fields{
				name:         "test",
				propertyType: Uint32Array,
				value:        []uint32{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Uint32Array Error",
			fields: fields{
				name:         "test",
				propertyType: Uint32Array,
				value:        "not an uint array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Uint64Array",
			fields: fields{
				name:         "test",
				propertyType: Uint64Array,
				value:        []uint64{10, 20, 30},
			},
			want:    []string{"10", "20", "30"},
			wantErr: false,
		},
		{
			name: "Uint64Array Error",
			fields: fields{
				name:         "test",
				propertyType: Uint64Array,
				value:        "not an uint array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Float32Array",
			fields: fields{
				name:         "test",
				propertyType: Float32Array,
				value:        []float32{10, 20, 30},
			},
			want:    []string{"10.000000", "20.000000", "30.000000"},
			wantErr: false,
		},
		{
			name: "Float32Array Error",
			fields: fields{
				name:         "test",
				propertyType: Float32Array,
				value:        "not a float array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Float64Array",
			fields: fields{
				name:         "test",
				propertyType: Float64Array,
				value:        []float64{10, 20, 30},
			},
			want:    []string{"10.000000", "20.000000", "30.000000"},
			wantErr: false,
		},
		{
			name: "Float64Array Error",
			fields: fields{
				name:         "test",
				propertyType: Float64Array,
				value:        "not a float array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "StringArray",
			fields: fields{
				name:         "test",
				propertyType: StringArray,
				value:        []string{"this", "is", "a", "test"},
			},
			want:    []string{"this", "is", "a", "test"},
			wantErr: false,
		},
		{
			name: "StringArray Error",
			fields: fields{
				name:         "test",
				propertyType: StringArray,
				value:        "not a string array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "DateTimeArray",
			fields: fields{
				name:         "test",
				propertyType: DateTimeArray,
				value: []time.Time{
					time.Date(1970, 1, 1, 0, 0, 0, 123000000, time.UTC),
					time.Date(1970, 1, 1, 1, 0, 0, 123000000, time.UTC),
				},
			},
			want:    []string{"1970-01-01T00:00:00.123Z", "1970-01-01T01:00:00.123Z"},
			wantErr: false,
		},
		{
			name: "DateTimeArray Error",
			fields: fields{
				name:         "test",
				propertyType: DateTimeArray,
				value:        "not a DateTime array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "DurationArray",
			fields: fields{
				name:         "test",
				propertyType: DurationArray,
				value: []time.Duration{
					time.Minute,
					time.Hour,
				},
			},
			want:    []string{"1m0s", "1h0m0s"},
			wantErr: false,
		},
		{
			name: "DurationArray Error",
			fields: fields{
				name:         "test",
				propertyType: DurationArray,
				value:        "not a Duration array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "DecimalArray",
			fields: fields{
				name:         "test",
				propertyType: DecimalArray,
				value:        []decimal.Decimal{decimal.New(1, 1), decimal.New(2, 1), decimal.New(3, 1)},
			},
			want:    []string{"10.000000", "20.000000", "30.000000"},
			wantErr: false,
		},
		{
			name: "DecimalArray Error",
			fields: fields{
				name:         "test",
				propertyType: DecimalArray,
				value:        "not a Decimal array",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Undefined",
			fields: fields{
				name:         "test",
				propertyType: Undefined,
				value:        nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.StringArray()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_StringArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "Bool",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []string{"default"},
			},
			want: []string{"true"},
		},
		{
			name: "Int",
			fields: fields{
				name:         "test",
				propertyType: Int,
				value:        10,
			},
			args: args{
				d: []string{"default"},
			},
			want: []string{"10"},
		},
		{
			name: "Int8",
			fields: fields{
				name:         "test",
				propertyType: Int8,
				value:        int8(10),
			},
			args: args{
				d: []string{"default"},
			},
			want: []string{"10"},
		},
		{
			name: "Int16",
			fields: fields{
				name:         "test",
				propertyType: Int16,
				value:        int16(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Int32",
			fields: fields{
				name:         "test",
				propertyType: Int32,
				value:        int32(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Int64",
			fields: fields{
				name:         "test",
				propertyType: Int64,
				value:        int64(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Uint",
			fields: fields{
				name:         "test",
				propertyType: Uint,
				value:        uint(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Uint8",
			fields: fields{
				name:         "test",
				propertyType: Uint8,
				value:        uint8(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Uint16",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        uint16(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Uint32",
			fields: fields{
				name:         "test",
				propertyType: Uint32,
				value:        uint32(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Uint64",
			fields: fields{
				name:         "test",
				propertyType: Uint64,
				value:        uint64(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10"},
		},
		{
			name: "Float32",
			fields: fields{
				name:         "test",
				propertyType: Float32,
				value:        float32(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10.000000"},
		},
		{
			name: "Float64",
			fields: fields{
				name:         "test",
				propertyType: Float64,
				value:        float64(10),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10.000000"},
		},
		{
			name: "ByteArray",
			fields: fields{
				name:         "test",
				propertyType: ByteArray,
				value:        []byte("test"),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"Default"},
		},
		{
			name: "String",
			fields: fields{
				name:         "test",
				propertyType: String,
				value:        "test",
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"test"},
		},
		{
			name: "DateTime",
			fields: fields{
				name:         "test",
				propertyType: DateTime,
				value:        time.Date(1970, 1, 1, 0, 0, 0, 123000000, time.UTC),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"1970-01-01T00:00:00.123Z"},
		},
		{
			name: "Decimal",
			fields: fields{
				name:         "test",
				propertyType: Decimal,
				value:        decimal.New(1, 1),
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10.000000"},
		},
		{
			name: "BoolArray",
			fields: fields{
				name:         "test",
				propertyType: BoolArray,
				value:        []bool{true, false, true},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"true", "false", "true"},
		},
		{
			name: "IntArray",
			fields: fields{
				name:         "test",
				propertyType: IntArray,
				value:        []int{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Int8Array",
			fields: fields{
				name:         "test",
				propertyType: Int8Array,
				value:        []int8{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Int16Array",
			fields: fields{
				name:         "test",
				propertyType: Int16Array,
				value:        []int16{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Int32Array",
			fields: fields{
				name:         "test",
				propertyType: Int32Array,
				value:        []int32{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Int64Array",
			fields: fields{
				name:         "test",
				propertyType: Int64Array,
				value:        []int64{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "UintArray",
			fields: fields{
				name:         "test",
				propertyType: UintArray,
				value:        []uint{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Uint8Array",
			fields: fields{
				name:         "test",
				propertyType: Uint8Array,
				value:        []uint8{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Uint16Array",
			fields: fields{
				name:         "test",
				propertyType: Uint16Array,
				value:        []uint16{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Uint32Array",
			fields: fields{
				name:         "test",
				propertyType: Uint32Array,
				value:        []uint32{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Uint64Array",
			fields: fields{
				name:         "test",
				propertyType: Uint64Array,
				value:        []uint64{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10", "20", "30"},
		},
		{
			name: "Float32Array",
			fields: fields{
				name:         "test",
				propertyType: Float32Array,
				value:        []float32{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10.000000", "20.000000", "30.000000"},
		},
		{
			name: "Float64Array",
			fields: fields{
				name:         "test",
				propertyType: Float64Array,
				value:        []float64{10, 20, 30},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10.000000", "20.000000", "30.000000"},
		},
		{
			name: "StringArray",
			fields: fields{
				name:         "test",
				propertyType: StringArray,
				value:        []string{"this", "is", "a", "test"},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"this", "is", "a", "test"},
		},
		{
			name: "DateTimeArray",
			fields: fields{
				name:         "test",
				propertyType: DateTimeArray,
				value: []time.Time{
					time.Date(1970, 1, 1, 0, 0, 0, 123000000, time.UTC),
					time.Date(1970, 1, 1, 1, 0, 0, 123000000, time.UTC),
				},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"1970-01-01T00:00:00.123Z", "1970-01-01T01:00:00.123Z"},
		},
		{
			name: "DecimalArray",
			fields: fields{
				name:         "test",
				propertyType: DecimalArray,
				value:        []decimal.Decimal{decimal.New(1, 1), decimal.New(2, 1), decimal.New(3, 1)},
			},
			args: args{
				d: []string{"Default"},
			},
			want: []string{"10.000000", "20.000000", "30.000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.StringArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint,
				value:        uint(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint16(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint16
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        uint16(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint16()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint16Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []uint16
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint16Array,
				value:        []uint16{10, 20, 30},
			},
			want:    []uint16{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint16Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint16ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint16
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint16Array,
				value:        []uint16{10, 20, 30},
			},
			args: args{
				d: []uint16{100, 200, 300},
			},
			want: []uint16{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []uint16{100, 200, 300},
			},
			want: []uint16{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint16ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint16OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint16
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint16,
				value:        uint16(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint16OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint32(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint32
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint32,
				value:        uint32(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint32()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint32Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []uint32
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint32Array,
				value:        []uint32{10, 20, 30},
			},
			want:    []uint32{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint32Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint32ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint32
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint32Array,
				value:        []uint32{10, 20, 30},
			},
			args: args{
				d: []uint32{100, 200, 300},
			},
			want: []uint32{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []uint32{100, 200, 300},
			},
			want: []uint32{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint32ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint32OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint32,
				value:        uint32(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint32OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint64(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint64,
				value:        uint64(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64() got = %v, want %v", got, tt.want)
			}
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint64Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []uint64
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint64Array,
				value:        []uint64{10, 20, 30},
			},
			want:    []uint64{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint64Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint64ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint64
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint64Array,
				value:        []uint64{10, 20, 30},
			},
			args: args{
				d: []uint64{100, 200, 300},
			},
			want: []uint64{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []uint64{100, 200, 300},
			},
			want: []uint64{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint64ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint64OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint64,
				value:        uint64(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint64OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint8(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint8
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint8,
				value:        uint8(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint8() got = %v, want %v", got, tt.want)
			}
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint8Array(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []uint8
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint8Array,
				value:        []uint8{10, 20, 30},
			},
			want:    []uint8{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.Uint8Array()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Uint8ArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint8
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint8Array,
				value:        []uint8{10, 20, 30},
			},
			args: args{
				d: []uint8{50, 60, 70},
			},
			want: []uint8{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []uint8{50, 60, 70},
			},
			want: []uint8{50, 60, 70},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint8ArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_Uint8OrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint8,
				value:        uint8(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.Uint8OrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_UintArray(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []uint
		wantErr bool
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: UintArray,
				value:        []uint{10, 20, 30},
			},
			want:    []uint{10, 20, 30},
			wantErr: false,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got, err := p.UintArray()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_UintArrayOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d []uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: UintArray,
				value:        []uint{10, 20, 30},
			},
			args: args{
				d: []uint{100, 200, 300},
			},
			want: []uint{10, 20, 30},
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: []uint{100, 200, 300},
			},
			want: []uint{100, 200, 300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.UintArrayOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_UintOrDefault(t *testing.T) {
	type fields struct {
		name         string
		propertyType Type
		value        interface{}
	}
	type args struct {
		d uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		{
			name: "valid property type",
			fields: fields{
				name:         "test",
				propertyType: Uint,
				value:        uint(10),
			},
			args: args{
				d: 100,
			},
			want: 10,
		},
		{
			name: "invalid property type",
			fields: fields{
				name:         "test",
				propertyType: Bool,
				value:        true,
			},
			args: args{
				d: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Property{
				Name:  tt.fields.name,
				Type:  tt.fields.propertyType,
				Value: tt.fields.value,
			}
			got := p.UintOrDefault(tt.args.d)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestStringProperty(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "test string",
			args: args{
				name:  "test",
				value: "test",
			},
			want: Property{
				Name:  "test",
				Type:  String,
				Value: "test",
			},
		},
		{
			name: "empty string",
			args: args{
				name:  "test",
				value: "",
			},
			want: Property{
				Name:  "test",
				Type:  String,
				Value: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint16ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []uint16
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Uint16Array,
				Value: []uint16(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []uint16{},
			},
			want: Property{
				Name:  "test",
				Type:  Uint16Array,
				Value: []uint16{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []uint16{u16v01},
			},
			want: Property{
				Name:  "test",
				Type:  Uint16Array,
				Value: []uint16{u16v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []uint16{u16v01, u16v02, u16v03, u16v04, u16v05},
			},
			want: Property{
				Name:  "test",
				Type:  Uint16Array,
				Value: []uint16{u16v01, u16v02, u16v03, u16v04, u16v05},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint16Property(t *testing.T) {
	type args struct {
		name  string
		value uint16
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "u16v01",
			args: args{
				name:  "test",
				value: u16v01,
			},
			want: Property{
				Name:  "test",
				Type:  Uint16,
				Value: u16v01,
			},
		},
		{
			name: "u16v02",
			args: args{
				name:  "test",
				value: u16v02,
			},
			want: Property{
				Name:  "test",
				Type:  Uint16,
				Value: u16v02,
			},
		},
		{
			name: "u16v03",
			args: args{
				name:  "test",
				value: u16v03,
			},
			want: Property{
				Name:  "test",
				Type:  Uint16,
				Value: u16v03,
			},
		},
		{
			name: "u16v04",
			args: args{
				name:  "test",
				value: u16v04,
			},
			want: Property{
				Name:  "test",
				Type:  Uint16,
				Value: u16v04,
			},
		},
		{
			name: "u16v05",
			args: args{
				name:  "test",
				value: u16v05,
			},
			want: Property{
				Name:  "test",
				Type:  Uint16,
				Value: u16v05,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint32ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []uint32
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32Array,
				Value: []uint32(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []uint32{},
			},
			want: Property{
				Name:  "test",
				Type:  Uint32Array,
				Value: []uint32{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []uint32{u32v01},
			},
			want: Property{
				Name:  "test",
				Type:  Uint32Array,
				Value: []uint32{u32v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []uint32{u32v01, u32v02, u32v03, u32v04, u32v05, u32v06, u32v07},
			},
			want: Property{
				Name:  "test",
				Type:  Uint32Array,
				Value: []uint32{u32v01, u32v02, u32v03, u32v04, u32v05, u32v06, u32v07},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint32Property(t *testing.T) {
	type args struct {
		name  string
		value uint32
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "u3201",
			args: args{
				name:  "test",
				value: u32v01,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32,
				Value: u32v01,
			},
		},
		{
			name: "u3202",
			args: args{
				name:  "test",
				value: u32v02,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32,
				Value: u32v02,
			},
		},
		{
			name: "u3203",
			args: args{
				name:  "test",
				value: u32v03,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32,
				Value: u32v03,
			},
		},
		{
			name: "u3204",
			args: args{
				name:  "test",
				value: u32v04,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32,
				Value: u32v04,
			},
		},
		{
			name: "u3205",
			args: args{
				name:  "test",
				value: u32v05,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32,
				Value: u32v05,
			},
		},
		{
			name: "u3206",
			args: args{
				name:  "test",
				value: u32v06,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32,
				Value: u32v06,
			},
		},
		{
			name: "u3207",
			args: args{
				name:  "test",
				value: u32v07,
			},
			want: Property{
				Name:  "test",
				Type:  Uint32,
				Value: u32v07,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint64ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []uint64
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64Array,
				Value: []uint64(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []uint64{},
			},
			want: Property{
				Name:  "test",
				Type:  Uint64Array,
				Value: []uint64{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []uint64{u64v01},
			},
			want: Property{
				Name:  "test",
				Type:  Uint64Array,
				Value: []uint64{u64v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []uint64{u64v01, u64v02, u64v03, u64v04, u64v05, u64v06, u64v07, u64v08, u64v09},
			},
			want: Property{
				Name:  "test",
				Type:  Uint64Array,
				Value: []uint64{u64v01, u64v02, u64v03, u64v04, u64v05, u64v06, u64v07, u64v08, u64v09},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint64ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint64Property(t *testing.T) {
	type args struct {
		name  string
		value uint64
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "u64v01",
			args: args{
				name:  "test",
				value: u64v01,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v01,
			},
		},
		{
			name: "u64v02",
			args: args{
				name:  "test",
				value: u64v02,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v02,
			},
		},
		{
			name: "u64v03",
			args: args{
				name:  "test",
				value: u64v03,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v03,
			},
		},
		{
			name: "u64v04",
			args: args{
				name:  "test",
				value: u64v04,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v04,
			},
		},
		{
			name: "u64v05",
			args: args{
				name:  "test",
				value: u64v05,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v05,
			},
		},
		{
			name: "u64v06",
			args: args{
				name:  "test",
				value: u64v06,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v06,
			},
		},
		{
			name: "u64v07",
			args: args{
				name:  "test",
				value: u64v07,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v07,
			},
		},
		{
			name: "u64v08",
			args: args{
				name:  "test",
				value: u64v08,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v08,
			},
		},
		{
			name: "u64v09",
			args: args{
				name:  "test",
				value: u64v09,
			},
			want: Property{
				Name:  "test",
				Type:  Uint64,
				Value: u64v09,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint64Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint8ArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []uint8
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  Uint8Array,
				Value: []uint8(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []uint8{},
			},
			want: Property{
				Name:  "test",
				Type:  Uint8Array,
				Value: []uint8{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []uint8{u8v01},
			},
			want: Property{
				Name:  "test",
				Type:  Uint8Array,
				Value: []uint8{u8v01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []uint8{u8v01, u8v02, u8v03},
			},
			want: Property{
				Name:  "test",
				Type:  Uint8Array,
				Value: []uint8{u8v01, u8v02, u8v03},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUint8Property(t *testing.T) {
	type args struct {
		name  string
		value uint8
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "u8v01",
			args: args{
				name:  "test",
				value: u8v01,
			},
			want: Property{
				Name:  "test",
				Type:  Uint8,
				Value: u8v01,
			},
		},
		{
			name: "u8v02",
			args: args{
				name:  "test",
				value: u8v02,
			},
			want: Property{
				Name:  "test",
				Type:  Uint8,
				Value: u8v02,
			},
		},
		{
			name: "u8v03",
			args: args{
				name:  "test",
				value: u8v03,
			},
			want: Property{
				Name:  "test",
				Type:  Uint8,
				Value: u8v03,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8Property(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUintArrayProperty(t *testing.T) {
	type args struct {
		name  string
		value []uint
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "nil",
			args: args{
				name:  "test",
				value: nil,
			},
			want: Property{
				Name:  "test",
				Type:  UintArray,
				Value: []uint(nil),
			},
		},
		{
			name: "empty slice",
			args: args{
				name:  "test",
				value: []uint{},
			},
			want: Property{
				Name:  "test",
				Type:  UintArray,
				Value: []uint{},
			},
		},
		{
			name: "one item",
			args: args{
				name:  "test",
				value: []uint{u01},
			},
			want: Property{
				Name:  "test",
				Type:  UintArray,
				Value: []uint{u01},
			},
		},
		{
			name: "multiple items",
			args: args{
				name:  "test",
				value: []uint{u01, u02, u03, u04, u05, u06, u07},
			},
			want: Property{
				Name:  "test",
				Type:  UintArray,
				Value: []uint{u01, u02, u03, u04, u05, u06, u07},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UintArrayProperty(tt.args.name, tt.args.value...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUintProperty(t *testing.T) {
	type args struct {
		name  string
		value uint
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "u01",
			args: args{
				name:  "test",
				value: u01,
			},
			want: Property{
				Name:  "test",
				Type:  Uint,
				Value: u01,
			},
		},
		{
			name: "u02",
			args: args{
				name:  "test",
				value: u02,
			},
			want: Property{
				Name:  "test",
				Type:  Uint,
				Value: u02,
			},
		},
		{
			name: "u03",
			args: args{
				name:  "test",
				value: u03,
			},
			want: Property{
				Name:  "test",
				Type:  Uint,
				Value: u03,
			},
		},
		{
			name: "u04",
			args: args{
				name:  "test",
				value: u04,
			},
			want: Property{
				Name:  "test",
				Type:  Uint,
				Value: u04,
			},
		},
		{
			name: "u05",
			args: args{
				name:  "test",
				value: u05,
			},
			want: Property{
				Name:  "test",
				Type:  Uint,
				Value: u05,
			},
		},
		{
			name: "u06",
			args: args{
				name:  "test",
				value: u06,
			},
			want: Property{
				Name:  "test",
				Type:  Uint,
				Value: u06,
			},
		},
		{
			name: "u07",
			args: args{
				name:  "test",
				value: u07,
			},
			want: Property{
				Name:  "test",
				Type:  Uint,
				Value: u07,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UintProperty(tt.args.name, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_stringToBool(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "y",
			args: args{
				v: "y",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Y",
			args: args{
				v: "Y",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "yes",
			args: args{
				v: "yes",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "YES",
			args: args{
				v: "YES",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "n",
			args: args{
				v: "n",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "N",
			args: args{
				v: "N",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "no",
			args: args{
				v: "no",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "NO",
			args: args{
				v: "NO",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "t",
			args: args{
				v: "t",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "T",
			args: args{
				v: "T",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true",
			args: args{
				v: "true",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "TRUE",
			args: args{
				v: "TRUE",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "f",
			args: args{
				v: "f",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "F",
			args: args{
				v: "F",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "false",
			args: args{
				v: "false",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "FALSE",
			args: args{
				v: "FALSE",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				v: "fail",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringToBool(tt.args.v)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestProperty_Name(t *testing.T) {
	p := StringProperty("key", "value")
	assert.Equal(t, "key", p.Name)
}

func TestNew(t *testing.T) {
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Property
	}{
		{
			name: "Bool",
			args: args{
				name:  "bool",
				value: true,
			},
			want: Property{
				Name:  "bool",
				Type:  Bool,
				Value: true,
			},
		},
		{
			name: "Int",
			args: args{
				name:  "int",
				value: 1,
			},
			want: Property{
				Name:  "int",
				Type:  Int,
				Value: 1,
			},
		},
		{
			name: "Int8",
			args: args{
				name:  "int8",
				value: int8(1),
			},
			want: Property{
				Name:  "int8",
				Type:  Int8,
				Value: int8(1),
			},
		},
		{
			name: "Int16",
			args: args{
				name:  "int16",
				value: int16(1),
			},
			want: Property{
				Name:  "int16",
				Type:  Int16,
				Value: int16(1),
			},
		},
		{
			name: "Int32",
			args: args{
				name:  "int32",
				value: int32(1),
			},
			want: Property{
				Name:  "int32",
				Type:  Int32,
				Value: int32(1),
			},
		},
		{
			name: "Int64",
			args: args{
				name:  "int64",
				value: int64(1),
			},
			want: Property{
				Name:  "int64",
				Type:  Int64,
				Value: int64(1),
			},
		},
		{
			name: "Uint",
			args: args{
				name:  "uint",
				value: uint(1),
			},
			want: Property{
				Name:  "uint",
				Type:  Uint,
				Value: uint(1),
			},
		},
		{
			name: "Uint8",
			args: args{
				name:  "uint8",
				value: uint8(1),
			},
			want: Property{
				Name:  "uint8",
				Type:  Uint8,
				Value: uint8(1),
			},
		},
		{
			name: "Uint16",
			args: args{
				name:  "uint16",
				value: uint16(1),
			},
			want: Property{
				Name:  "uint16",
				Type:  Uint16,
				Value: uint16(1),
			},
		},
		{
			name: "Uint32",
			args: args{
				name:  "uint32",
				value: uint32(1),
			},
			want: Property{
				Name:  "uint32",
				Type:  Uint32,
				Value: uint32(1),
			},
		},
		{
			name: "Uint64",
			args: args{
				name:  "uint64",
				value: uint64(1),
			},
			want: Property{
				Name:  "uint64",
				Type:  Uint64,
				Value: uint64(1),
			},
		},
		{
			name: "Float32",
			args: args{
				name:  "float32",
				value: float32(1),
			},
			want: Property{
				Name:  "float32",
				Type:  Float32,
				Value: float32(1),
			},
		},
		{
			name: "Float64",
			args: args{
				name:  "float64",
				value: float64(1),
			},
			want: Property{
				Name:  "float64",
				Type:  Float64,
				Value: float64(1),
			},
		},
		{
			name: "ByteArray",
			args: args{
				name:  "bytearray",
				value: []byte("test"),
			},
			want: Property{
				Name:  "bytearray",
				Type:  ByteArray,
				Value: []byte("test"),
			},
		},
		{
			name: "String",
			args: args{
				name:  "string",
				value: "test",
			},
			want: Property{
				Name:  "string",
				Type:  String,
				Value: "test",
			},
		},
		{
			name: "DateTime",
			args: args{
				name:  "datetime",
				value: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: Property{
				Name:  "datetime",
				Type:  DateTime,
				Value: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Duration",
			args: args{
				name:  "duration",
				value: time.Second,
			},
			want: Property{
				Name:  "duration",
				Type:  Duration,
				Value: time.Second,
			},
		},
		{
			name: "Decimal",
			args: args{
				name:  "decimal",
				value: decimal.Zero,
			},
			want: Property{
				Name:  "decimal",
				Type:  Decimal,
				Value: decimal.Zero,
			},
		},
		{
			name: "BoolArray",
			args: args{
				name:  "boolarray",
				value: []bool{true, false},
			},
			want: Property{
				Name:  "boolarray",
				Type:  BoolArray,
				Value: []bool{true, false},
			},
		},
		{
			name: "IntArray",
			args: args{
				name:  "intarray",
				value: []int{1, 2, 3},
			},
			want: Property{
				Name:  "intarray",
				Type:  IntArray,
				Value: []int{1, 2, 3},
			},
		},
		{
			name: "Int8Array",
			args: args{
				name:  "int8array",
				value: []int8{1, 2, 3},
			},
			want: Property{
				Name:  "int8array",
				Type:  Int8Array,
				Value: []int8{1, 2, 3},
			},
		},
		{
			name: "Int16Array",
			args: args{
				name:  "int16array",
				value: []int16{1, 2, 3},
			},
			want: Property{
				Name:  "int16array",
				Type:  Int16Array,
				Value: []int16{1, 2, 3},
			},
		},
		{
			name: "Int32Array",
			args: args{
				name:  "int32array",
				value: []int32{1, 2, 3},
			},
			want: Property{
				Name:  "int32array",
				Type:  Int32Array,
				Value: []int32{1, 2, 3},
			},
		},
		{
			name: "Int64Array",
			args: args{
				name:  "int64array",
				value: []int64{1, 2, 3},
			},
			want: Property{
				Name:  "int64array",
				Type:  Int64Array,
				Value: []int64{1, 2, 3},
			},
		},
		{
			name: "UintArray",
			args: args{
				name:  "uintarray",
				value: []uint{1, 2, 3},
			},
			want: Property{
				Name:  "uintarray",
				Type:  UintArray,
				Value: []uint{1, 2, 3},
			},
		},
		{
			name: "Uint16Array",
			args: args{
				name:  "uint16array",
				value: []uint16{1, 2, 3},
			},
			want: Property{
				Name:  "uint16array",
				Type:  Uint16Array,
				Value: []uint16{1, 2, 3},
			},
		},
		{
			name: "Uint32Array",
			args: args{
				name:  "uint32array",
				value: []uint32{1, 2, 3},
			},
			want: Property{
				Name:  "uint32array",
				Type:  Uint32Array,
				Value: []uint32{1, 2, 3},
			},
		},
		{
			name: "Uint64Array",
			args: args{
				name:  "uint64array",
				value: []uint64{1, 2, 3},
			},
			want: Property{
				Name:  "uint64array",
				Type:  Uint64Array,
				Value: []uint64{1, 2, 3},
			},
		},
		{
			name: "Float32Array",
			args: args{
				name:  "float32array",
				value: []float32{1.1, 2.2, 3.3},
			},
			want: Property{
				Name:  "float32array",
				Type:  Float32Array,
				Value: []float32{1.1, 2.2, 3.3},
			},
		},
		{
			name: "Float64Array",
			args: args{
				name:  "float64array",
				value: []float64{1.1, 2.2, 3.3},
			},
			want: Property{
				Name:  "float64array",
				Type:  Float64Array,
				Value: []float64{1.1, 2.2, 3.3},
			},
		},
		{
			name: "StringArray",
			args: args{
				name:  "stringarray",
				value: []string{"a", "b", "c"},
			},
			want: Property{
				Name:  "stringarray",
				Type:  StringArray,
				Value: []string{"a", "b", "c"},
			},
		},
		{
			name: "DateTimeArray",
			args: args{
				name: "datetimearray",
				value: []time.Time{
					time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
				},
			},
			want: Property{
				Name: "datetimearray",
				Type: DateTimeArray,
				Value: []time.Time{
					time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "DurationArray",
			args: args{
				name:  "durationarray",
				value: []time.Duration{1, 2, 3},
			},
			want: Property{
				Name:  "durationarray",
				Type:  DurationArray,
				Value: []time.Duration{1, 2, 3},
			},
		},
		{
			name: "DecimalArray",
			args: args{
				name: "decimalarray",
				value: []decimal.Decimal{
					decimal.NewFromInt(1),
					decimal.NewFromInt(2),
					decimal.NewFromInt(3),
				},
			},
			want: Property{
				Name: "decimalarray",
				Type: DecimalArray,
				Value: []decimal.Decimal{
					decimal.NewFromInt(1),
					decimal.NewFromInt(2),
					decimal.NewFromInt(3),
				},
			},
		},
		{
			name: "InterfaceArray",
			args: args{
				name:  "interfacearray",
				value: []interface{}{1, 2, 3},
			},
			want: Property{
				Name:  "interfacearray",
				Type:  InterfaceArray,
				Value: []interface{}{1, 2, 3},
			},
		},
		{
			name: "Unsupported",
			args: args{
				name:  "map",
				value: map[int]string{1: "a", 2: "b", 3: "c"},
			},
			want: Property{
				Name:  "map",
				Type:  Interface,
				Value: map[int]string{1: "a", 2: "b", 3: "c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.name, tt.args.value)
			assert.Equal(t, tt.want, got, "New() = %v, want %v", got, tt.want)
		})
	}
}

func TestProperty_MarshalBinary(t *testing.T) {
	type args struct {
		name  string
		value interface{}
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Bool",
			args: args{
				name:  "bool",
				value: true,
			},
			want:    []byte(`{"name":"bool","type":1,"value":true}`),
			wantErr: false,
		},
		{
			name: "BoolArray",
			args: args{
				name:  "boolarray",
				value: []bool{true, false, true},
			},
			want:    []byte(`{"name":"boolarray","type":20,"value":[true,false,true]}`),
			wantErr: false,
		},
		{
			name: "Int",
			args: args{
				name:  "int",
				value: 10,
			},
			want:    []byte(`{"name":"int","type":2,"value":10}`),
			wantErr: false,
		},
		{
			name: "IntArray",
			args: args{
				name:  "intarray",
				value: []int{10, 20, 30},
			},
			want:    []byte(`{"name":"intarray","type":21,"value":[10,20,30]}`),
			wantErr: false,
		},
		{
			name: "Int8",
			args: args{
				name:  "int8",
				value: int8(10),
			},
			want:    []byte(`{"name":"int8","type":3,"value":10}`),
			wantErr: false,
		},
		{
			name: "Int8Array",
			args: args{
				name:  "int8array",
				value: []int8{10, 20, 30},
			},
			want:    []byte(`{"name":"int8array","type":22,"value":[10,20,30]}`),
			wantErr: false,
		},
		{
			name: "Duration",
			args: args{
				name:  "duration",
				value: time.Minute,
			},
			want:    []byte(`{"name":"duration","type":17,"value":60000000000}`),
			wantErr: false,
		},
		{
			name: "DurationArray",
			args: args{
				name:  "durationarray",
				value: []time.Duration{time.Minute, time.Minute * 2, time.Minute * 3},
			},
			want:    []byte(`{"name":"durationarray","type":35,"value":[60000000000,120000000000,180000000000]}`),
			wantErr: false,
		},
		{
			name: "DateTime",
			args: args{
				name:  "datetime",
				value: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want:    []byte(`{"name":"datetime","type":16,"value":"2020-01-01T00:00:00Z"}`),
			wantErr: false,
		},
		{
			name: "DateTimeArray",
			args: args{
				name: "datetimearray",
				value: []time.Time{
					time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
				},
			},
			want:    []byte(`{"name":"datetimearray","type":34,"value":["2020-01-01T00:00:00Z","2020-01-02T00:00:00Z","2020-01-03T00:00:00Z"]}`),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.args.name, tt.args.value)
			got, err := p.MarshalBinary()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProperty_UnmarshalBinary(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    Property
		wantErr bool
	}{
		{
			name: "bool",
			args: args{
				data: []byte(`{"name":"bool","type":1,"value":true}`),
			},
			want: Property{
				Name:  "bool",
				Type:  Bool,
				Value: true,
			},
			wantErr: false,
		},
		{
			name: "bool array",
			args: args{
				data: []byte(`{"name":"boolarray","type":20,"value":[true,false,true]}`),
			},
			want: Property{
				Name:  "boolarray",
				Type:  BoolArray,
				Value: []bool{true, false, true},
			},
			wantErr: false,
		},
		{
			name: "int",
			args: args{
				data: []byte(`{"name":"int","type":2,"value":10}`),
			},
			want: Property{
				Name:  "int",
				Type:  Int,
				Value: 10,
			},
			wantErr: false,
		},
		{
			name: "int array",
			args: args{
				data: []byte(`{"name":"intarray","type":21,"value":[1,2,3]}`),
			},
			want: Property{
				Name:  "intarray",
				Type:  IntArray,
				Value: []int{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "int8",
			args: args{
				data: []byte(`{"name":"int8","type":3,"value":10}`),
			},
			want: Property{
				Name:  "int8",
				Type:  Int8,
				Value: int8(10),
			},
			wantErr: false,
		},
		{
			name: "int8 array",
			args: args{
				data: []byte(`{"name":"int8array","type":22,"value":[1,2,3]}`),
			},
			want: Property{
				Name:  "int8array",
				Type:  Int8Array,
				Value: []int8{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "uint",
			args: args{
				data: []byte(`{"name":"uint","type":7,"value":10}`),
			},
			want: Property{
				Name:  "uint",
				Type:  Uint,
				Value: uint(10),
			},
			wantErr: false,
		},
		{
			name: "uint array",
			args: args{
				data: []byte(`{"name":"uintarray","type":26,"value":[1,2,3]}`),
			},
			want: Property{
				Name:  "uintarray",
				Type:  UintArray,
				Value: []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "float32",
			args: args{
				data: []byte(`{"name":"float32","type":12,"value":10}`),
			},
			want: Property{
				Name:  "float32",
				Type:  Float32,
				Value: float32(10),
			},
			wantErr: false,
		},
		{
			name: "float32 array",
			args: args{
				data: []byte(`{"name":"float32array","type":31,"value":[1,2,3]}`),
			},
			want: Property{
				Name:  "float32array",
				Type:  Float32Array,
				Value: []float32{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "float64",
			args: args{
				data: []byte(`{"name":"float64","type":13,"value":10}`),
			},
			want: Property{
				Name:  "float64",
				Type:  Float64,
				Value: float64(10),
			},
			wantErr: false,
		},
		{
			name: "float64 array",
			args: args{
				data: []byte(`{"name":"float64array","type":32,"value":[1,2,3]}`),
			},
			want: Property{
				Name:  "float64array",
				Type:  Float64Array,
				Value: []float64{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "string",
			args: args{
				data: []byte(`{"name":"string","type":15,"value":"test"}`),
			},
			want: Property{
				Name:  "string",
				Type:  String,
				Value: "test",
			},
			wantErr: false,
		},
		{
			name: "string array",
			args: args{
				data: []byte(`{"name":"stringarray","type":33,"value":["a","b","c"]}`),
			},
			want: Property{
				Name:  "stringarray",
				Type:  StringArray,
				Value: []string{"a", "b", "c"},
			},
			wantErr: false,
		},
		{
			name: "datetime",
			args: args{
				data: []byte(`{"name":"datetime","type":16,"value":"2020-01-01T00:00:00Z"}`),
			},
			want: Property{
				Name:  "datetime",
				Type:  DateTime,
				Value: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "datetime array",
			args: args{
				data: []byte(`{"name":"datetimearray","type":34,"value":["2020-01-01T00:00:00Z", "2020-01-02T00:00:00Z"]}`),
			},
			want: Property{
				Name: "datetimearray",
				Type: DateTimeArray,
				Value: []time.Time{
					time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "duration",
			args: args{
				data: []byte(`{"name":"duration","type":17,"value":60000000000}`),
			},
			want: Property{
				Name:  "duration",
				Type:  Duration,
				Value: time.Minute,
			},
			wantErr: false,
		},
		{
			name: "duration array",
			args: args{
				data: []byte(`{"name":"durationarray","type":35,"value":[60000000000,120000000000,180000000000]}`),
			},
			want: Property{
				Name: "durationarray",
				Type: DurationArray,
				Value: []time.Duration{
					time.Minute,
					time.Minute * 2,
					time.Minute * 3,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := new(Property)
			err := got.UnmarshalBinary(tt.args.data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, *got)
		})
	}
}

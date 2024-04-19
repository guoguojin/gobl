package property_test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/stretchr/testify/assert"

	"gitlab.com/gobl/gobl/pkg/property"
)

func TestProperties_Add(t *testing.T) {
	props := property.NewProperties()
	props.Add(property.StringProperty("key", "value"))

	v := props["key"]
	assert.Equal(t, v.String(), "value")
}

func TestProperties_Get(t *testing.T) {
	props := property.NewProperties()
	props.Add(property.StringProperty("key", "value"))

	v, ok := props.Get("key")
	assert.True(t, ok)
	assert.Equal(t, v.String(), "value")

	v, ok = props.Get("not-a-key")
	assert.False(t, ok)
	assert.Equal(t, property.Property{}, v)
}

func TestProperties_GetOrElse(t *testing.T) {
	props := property.NewProperties()
	props.Add(property.StringProperty("key", "value"))

	v := props.GetOrElse("key", property.StringProperty("key", "other-value"))
	assert.Equal(t, v.String(), "value")

	v = props.GetOrElse("not-a-key", property.StringProperty("key", "other-value"))
	assert.Equal(t, v.String(), "other-value")
	_, ok := props["not-a-key"]
	assert.True(t, ok)
}

func TestProperties_MarshalBinary(t *testing.T) {
	type args struct {
		props []property.Property
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "string",
			args: args{
				props: []property.Property{
					property.StringProperty("key", "value"),
				},
			},
			want:    []byte(`{"key":{"name":"key","type":15,"value":"value"}}`),
			wantErr: false,
		},
		{
			name: "string and bool",
			args: args{
				props: []property.Property{
					property.StringProperty("key", "value"),
					property.BoolProperty("bool", true),
				},
			},
			want:    []byte(`{"bool":{"name":"bool","type":1,"value":true},"key":{"name":"key","type":15,"value":"value"}}`),
			wantErr: false,
		},
		{
			name: "string and bool and int",
			args: args{
				props: []property.Property{
					property.StringProperty("key", "value"),
					property.BoolProperty("bool", true),
					property.IntProperty("int", 1),
				},
			},
			want:    []byte(`{"bool":{"name":"bool","type":1,"value":true},"int":{"name":"int","type":2,"value":1},"key":{"name":"key","type":15,"value":"value"}}`),
			wantErr: false,
		},
		{
			name: "int8, datetime, duration and decimal",
			args: args{
				props: []property.Property{
					property.Int8Property("int8", 1),
					property.DateTimeProperty("datetime", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
					property.DurationProperty("duration", time.Minute),
					property.DecimalProperty("decimal", decimal.NewFromFloat(1.0)),
				},
			},
			want:    []byte(`{"datetime":{"name":"datetime","type":16,"value":"2020-01-01T00:00:00Z"},"decimal":{"name":"decimal","type":18,"value":"1"},"duration":{"name":"duration","type":17,"value":60000000000},"int8":{"name":"int8","type":3,"value":1}}`),
			wantErr: false,
		},
		{
			name: "uint, uint8, float32 and float64",
			args: args{
				props: []property.Property{
					property.UintProperty("uint", 1),
					property.Uint8Property("uint8", 1),
					property.Float32Property("float32", 1.0),
					property.Float64Property("float64", 1.0),
				},
			},
			want:    []byte(`{"float32":{"name":"float32","type":12,"value":1},"float64":{"name":"float64","type":13,"value":1},"uint":{"name":"uint","type":7,"value":1},"uint8":{"name":"uint8","type":8,"value":1}}`),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			props := property.NewProperties()
			for _, prop := range tt.args.props {
				props.Add(prop)
			}

			got, err := props.MarshalBinary()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProperties_UnmarshalBinary(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    property.Properties
		wantErr bool
	}{
		{
			name: "string",
			args: args{
				data: []byte(`{"key":{"name":"key","type":15,"value":"value"}}`),
			},
			want: property.Properties{
				"key": property.StringProperty("key", "value"),
			},
			wantErr: false,
		},
		{
			name: "string and bool",
			args: args{
				data: []byte(`{"bool":{"name":"bool","type":1,"value":true},"key":{"name":"key","type":15,"value":"value"}}`),
			},
			want: property.Properties{
				"key":  property.StringProperty("key", "value"),
				"bool": property.BoolProperty("bool", true),
			},
			wantErr: false,
		},
		{
			name: "string and bool and int",
			args: args{
				data: []byte(`{"bool":{"name":"bool","type":1,"value":true},"int":{"name":"int","type":2,"value":1},"key":{"name":"key","type":15,"value":"value"}}`),
			},
			want: property.Properties{
				"key":  property.StringProperty("key", "value"),
				"bool": property.BoolProperty("bool", true),
				"int":  property.IntProperty("int", 1),
			},
			wantErr: false,
		},
		{
			name: "int8, datetime, duration and decimal",
			args: args{
				data: []byte(`{"datetime":{"name":"datetime","type":16,"value":"2020-01-01T00:00:00Z"},"decimal":{"name":"decimal","type":18,"value":"1"},"duration":{"name":"duration","type":17,"value":60000000000},"int8":{"name":"int8","type":3,"value":1}}`),
			},
			want: property.Properties{
				"int8":     property.Int8Property("int8", 1),
				"datetime": property.DateTimeProperty("datetime", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
				"duration": property.DurationProperty("duration", time.Minute),
				"decimal":  property.DecimalProperty("decimal", decimal.NewFromFloat(1.0)),
			},
			wantErr: false,
		},
		{
			name: "uint, uint8, float32 and float64",
			args: args{
				data: []byte(`{"float32":{"name":"float32","type":12,"value":1},"float64":{"name":"float64","type":13,"value":1},"uint":{"name":"uint","type":7,"value":1},"uint8":{"name":"uint8","type":8,"value":1}}`),
			},
			want: property.Properties{
				"uint":    property.UintProperty("uint", 1),
				"uint8":   property.Uint8Property("uint8", 1),
				"float32": property.Float32Property("float32", 1.0),
				"float64": property.Float64Property("float64", 1.0),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := property.NewProperties()
			err := got.UnmarshalBinary(tt.args.data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

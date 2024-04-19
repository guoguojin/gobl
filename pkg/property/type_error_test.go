package property_test

import (
	"testing"

	"gitlab.com/gobl/gobl/pkg/property"
)

func TestTypeError_Error(t *testing.T) {
	type args struct {
		propertyType property.Type
		msg          string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestTypeError_Error",
			args: args{
				propertyType: property.String,
				msg:          "some error",
			},
			want: "String - some error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := property.NewTypeError(tt.args.propertyType, tt.args.msg).Error(); got != tt.want {
				t.Errorf("TypeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

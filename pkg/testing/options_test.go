package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applyDefaultEnvVars(t *testing.T) {
	type args struct {
		defaults  []string
		overrides []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should add defaults if overrides do not contain them",
			args: args{
				defaults: []string{
					"A=1",
					"B=2",
				},
				overrides: []string{
					"C=3",
				},
			},
			want: []string{
				"A=1",
				"B=2",
				"C=3",
			},
		},
		{
			name: "should only add defaults that overrides do not contain",
			args: args{
				defaults: []string{
					"A=1",
					"B=2",
				},
				overrides: []string{
					"A=5",
					"C=3",
				},
			},
			want: []string{
				"A=5",
				"B=2",
				"C=3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want,
				applyDefaultEnvVars(tt.args.defaults, tt.args.overrides),
				"applyDefaultEnvVars(%v, %v)",
				tt.args.defaults,
				tt.args.overrides)
		})
	}
}

func Test_applyEnvOverrides_Panics(t *testing.T) {
	t.Run("should panic when invalid default environment variables are passed", func(t *testing.T) {
		assert.Panicsf(t, func() {
			_ = applyDefaultEnvVars([]string{"test"}, []string{"A=1", "B=2"})
		}, "should panic when invalid default environment variables are passed")
	})

	t.Run("should panic when invalid override environment variables are passed", func(t *testing.T) {
		assert.Panicsf(t, func() {
			_ = applyDefaultEnvVars([]string{"A=1", "B=2"}, []string{"test"})
		}, "should panic when invalid override environment variables are passed")
	})
}

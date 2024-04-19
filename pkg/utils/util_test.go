package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsTrue(t *testing.T) {
	type args struct {
		value string
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid",
			args: args{value: "invalid"},
			want: false,
		},
		{
			name: "empty",
			args: args{value: ""},
			want: false,
		},
		{
			name: "nil",
			args: args{value: ""},
			want: false,
		},
		{
			name: "true",
			args: args{value: "true"},
			want: true,
		},
		{
			name: "false",
			args: args{value: "false"},
			want: false,
		},
		{
			name: "TRUE",
			args: args{value: "TRUE"},
			want: true,
		},
		{
			name: "FALSE",
			args: args{value: "FALSE"},
			want: false,
		},
		{
			name: "t",
			args: args{value: "t"},
			want: true,
		},
		{
			name: "f",
			args: args{value: "f"},
			want: false,
		},
		{
			name: "T",
			args: args{value: "T"},
			want: true,
		},
		{
			name: "F",
			args: args{value: "F"},
			want: false,
		},
		{
			name: "yes",
			args: args{value: "yes"},
			want: true,
		},
		{
			name: "no",
			args: args{value: "no"},
			want: false,
		},
		{
			name: "YES",
			args: args{value: "YES"},
			want: true,
		},
		{
			name: "NO",
			args: args{value: "NO"},
			want: false,
		},
		{
			name: "y",
			args: args{value: "y"},
			want: true,
		},
		{
			name: "n",
			args: args{value: "n"},
			want: false,
		},
		{
			name: "Y",
			args: args{value: "Y"},
			want: true,
		},
		{
			name: "N",
			args: args{value: "N"},
			want: false,
		},
		{
			name: "1",
			args: args{value: "1"},
			want: true,
		},
		{
			name: "0",
			args: args{value: "0"},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			got := IsTrue(tc.args.value)
			assert.Equal(t, tc.want, got)
		})
	}
}

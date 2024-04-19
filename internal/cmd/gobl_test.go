package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "v1 is valid",
			args: args{
				version: "v1",
			},
			want: true,
		},
		{
			name: "1 is invalid",
			args: args{
				version: "1",
			},
			want: false,
		},
		{
			name: "v100 is valid",
			args: args{
				version: "v100",
			},
			want: true,
		},
		{
			name: "v2b is invalid",
			args: args{
				version: "v2b",
			},
			want: false,
		},
		{
			name: "v2.0 is invalid",
			args: args{
				version: "v2.0",
			},
			want: false,
		},
		{
			name: "v20 is valid",
			args: args{
				version: "v20",
			},
			want: true,
		},
		{
			name: "v100 is valid",
			args: args{
				version: "v20",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkVersion(tt.args.version); got != tt.want {
				t.Errorf("checkVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newGoblConfig(t *testing.T) {
	type args struct {
		moduleName    string
		projectFolder string
	}
	tests := []struct {
		name    string
		args    args
		want    *goblConfig
		wantErr bool
	}{
		{
			name: "github.com/birchwood-langham/gobl should not raise an error and provide a config",
			args: args{
				moduleName: "github.com/birchwood-langham/gobl",
			},
			want: &goblConfig{"github.com/birchwood-langham/gobl", "github.com", "birchwood-langham", "gobl",
				"gobl", "", "gobl", loremIpsum, templateFS},
			wantErr: false,
		},
		{
			name: "github.com/birchwood-langham/gobl/v2 should not raise an error and provide a config",
			args: args{
				moduleName: "github.com/birchwood-langham/gobl/v2",
			},
			want: &goblConfig{"github.com/birchwood-langham/gobl/v2", "github.com", "birchwood-langham",
				"gobl", "gobl", "v2", "gobl", loremIpsum, templateFS},
			wantErr: false,
		},
		{
			name: "github.com/gobl should raise an error",
			args: args{
				moduleName: "github.com/gobl",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "github.com should raise an error",
			args: args{
				moduleName: "github.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "gitlab.com/birchwood-langham/subfolder/subfolder2/gobl/v2 should not raise an error and provide a config",
			args: args{
				moduleName: "gitlab.com/birchwood-langham/subfolder/subfolder2/gobl/v2",
			},
			want: &goblConfig{"gitlab.com/birchwood-langham/subfolder/subfolder2/gobl/v2", "gitlab.com", "birchwood-langham/subfolder/subfolder2",
				"gobl", "gobl", "v2", "gobl", loremIpsum, templateFS},
			wantErr: false,
		},
		{
			name: "project folder can be overridden",
			args: args{
				moduleName:    "gitlab.com/birchwood-langham/subfolder/subfolder2/gobl/v2",
				projectFolder: "my-project-folder",
			},
			want: &goblConfig{"gitlab.com/birchwood-langham/subfolder/subfolder2/gobl/v2", "gitlab.com", "birchwood-langham/subfolder/subfolder2",
				"gobl", "my-project-folder", "v2", "gobl", loremIpsum, templateFS,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newGoblConfig(tt.args.moduleName, tt.args.projectFolder, templateFS)
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

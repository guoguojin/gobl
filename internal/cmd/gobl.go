package cmd

import (
	"errors"
	"io/fs"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var (
	goblCmd = &cobra.Command{
		Run: gobl,
	}
	ErrInvalidModuleName = errors.New("the module name appears to be invalid")
)

const (
	GoblUsage     = "gobl"
	GoblShortDesc = "bootstrap a Go application"
	GoblLongDesc  = `gobl allows you to setup a Go service application quickly by

    - initialising a Go module
    - creating the directory structure for your application
    - creating a docker build for your application
    - providing docker-compose script stubs for you to build development, testing environments
    - simple make file
    - basic .gitignore, .dockerignore, .gitattributes, readme.md files etc. to get your project started
`
)

type goblConfig struct {
	GoModule           string
	Registry           string
	ProjectRepo        string
	ProjectName        string
	ProjectFolder      string
	Version            string
	ProjectTitle       string
	ProjectDescription string
	ProjectFS          fs.FS
}

const expectedMinParts = 3

type bootstrapFn func(config *goblConfig) error

// GoblCmd is the command that will be run by the gobl CLI application
func GoblCmd() *cobra.Command {
	return goblCmd
}

func gobl(c *cobra.Command, _ []string) {
	_ = c.Help()
}

func newGoblConfig(moduleName, projectFolder string, projectFS fs.FS) (*goblConfig, error) {
	parts := strings.Split(moduleName, "/")

	if len(parts) < expectedMinParts {
		return nil, ErrInvalidModuleName
	}

	var version, projectName string

	registry := parts[0]
	parts = parts[1:]
	projectNameIndex := len(parts) - 1

	if checkVersion(parts[len(parts)-1]) {
		version = parts[len(parts)-1]
		projectName = parts[len(parts)-2]
		projectNameIndex--
	} else {
		projectName = parts[len(parts)-1]
	}

	builder := strings.Builder{}
	for i := 0; i < projectNameIndex; i++ {
		if i > 0 {
			builder.WriteRune('/')
		}
		builder.WriteString(parts[i])
	}

	projectRepo := builder.String()

	if projectFolder == "" {
		projectFolder = projectName
	}

	title := projectName
	description := loremIpsum

	if projectTitle != "" {
		title = projectTitle
	}

	if projectDescription != "" {
		description = projectDescription
	}

	return &goblConfig{
		GoModule:           moduleName,
		Registry:           registry,
		ProjectRepo:        projectRepo,
		ProjectName:        projectName,
		ProjectFolder:      projectFolder,
		Version:            version,
		ProjectTitle:       title,
		ProjectDescription: description,
		ProjectFS:          projectFS,
	}, nil
}

func checkVersion(version string) bool {
	regex := regexp.MustCompile(`^v[1-9]\d*$`)
	return regex.MatchString(version)
}

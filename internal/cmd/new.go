package cmd

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"gitlab.com/gobl/gobl/pkg/logger"
)

//nolint:gochecknoglobals
var (
	newCmd = &cobra.Command{
		Use:   "new",
		Short: "bootstrap a new go project",
		Long: `gobl new <go-module-name>

example:
    gobl new github.com/my-github-space/my-new-project

    This will bootstrap a new application with default folder structures
    and a basic configuration for building a docker based build as well as
    docker-compose configurations for development, testing and production
    environments.

    A basic make file is also provided ready for you to extend for your own
    purposes.

    The application name will be derived from the Go module name provided,
    for example, the above command will create an application called
    my-new-project.`,
		Run: bootstrapNewApplication,
	}

	errGitSubFolder    = errors.New("project will be created under an existing git project")
	ignoreGit          bool
	projectFolder      string
	projectTitle       string
	projectDescription string

	//go:embed templates/*
	templateFS embed.FS
)

const (
	loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt
ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore
eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
mollit anim id est laborum.`
	folderPermissions = 0o755
	gitExec           = "git"
	goExec            = "go"
	preCommitExec     = "pre-commit"
)

func NewCmd() *cobra.Command {
	newCmd.Flags().BoolVar(&ignoreGit, "ignore-git", false,
		"ignore any git parent folders allowing you to create a new project under and existing git project")
	newCmd.Flags().StringVar(&projectFolder, "project-folder", "",
		"specify the folder the project should be created, otherwise the project name will be used")
	newCmd.Flags().StringVar(&projectTitle, "title", "",
		"Title for the readme.md page (project name will be used as default if not specified)")
	newCmd.Flags().StringVar(&projectDescription, "description", "",
		"description for the readme.md page (Lorem Ipsum will be used if not specified)")
	return newCmd
}

//nolint:forbidigo
func bootstrapNewApplication(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
		return
	}

	moduleName := args[0]

	l := logger.Logger()
	l.Info("Bootstrapping Go module", zap.String("module-name", moduleName))

	// if we try to create the project inside an existing git repo
	// then this will fail and terminate.
	validatePath(moduleName)

	config, err := newGoblConfig(moduleName, projectFolder, templateFS)
	if err != nil {
		fmt.Printf("cannot bootstrap %s: %s\n", moduleName, err)
		return
	}

	bootstrapProject(
		config,
		createProjectFolder,
		createProjectFiles,
		initModule,
		initGit,
		initPreCommit,
		initCommit,
	)
}

func validatePath(moduleName string) {
	l := logger.Logger().With(zap.String("module-name", moduleName))
	homeDir, err := os.UserHomeDir()
	if err != nil {
		l.Fatal("cannot determine user's home directory", zap.Error(err))
	}

	cwd, err := os.Getwd()
	l.Info("Current folder", zap.String("cwd", cwd))

	if err != nil {
		l.Fatal("cannot determine current directory", zap.Error(err))
	}

	if err := checkForGitParent(homeDir, cwd, os.Chdir, os.Stat); err != nil {
		l.Fatal("cannot bootstrap module", zap.Error(err))
	}

	// make sure we're back in the directory where the command was called
	if err := os.Chdir(cwd); err != nil {
		l.Fatal("could not bootstrap", zap.Error(err))
	}
}

func checkForGitParent(homeDir, currentDir string, chdir func(string) error, stat func(string) (os.FileInfo, error)) error {
	if ignoreGit {
		return nil
	}

	if err := chdir(currentDir); err != nil {
		return err
	}

	if _, err := stat(".git"); err == nil {
		return errGitSubFolder
	}

	if currentDir == homeDir {
		return nil
	}

	if strings.HasPrefix(currentDir, homeDir) {
		parent := filepath.Dir(currentDir)
		return checkForGitParent(homeDir, parent, chdir, stat)
	}

	return nil
}

func bootstrapProject(config *goblConfig, fns ...bootstrapFn) {
	l := logger.Logger().With(zap.String("module-name", config.GoModule))
	for _, fn := range fns {
		if err := fn(config); err != nil {
			l.Fatal("cannot bootstrap project", zap.Error(err))
		}
	}
}

func createProjectFolder(config *goblConfig) (err error) {
	l := logger.Logger().With(zap.String("module-name", config.GoModule))
	l.Info("creating project folder", zap.String("project-folder", config.ProjectFolder))

	if err = makeDirs(config.ProjectFolder); err != nil {
		return
	}

	return os.Chdir(config.ProjectFolder)
}

func createProjectFiles(config *goblConfig) (err error) {
	l := logger.Logger().With(zap.String("module-name", config.GoModule))
	fsys, err := fs.Sub(config.ProjectFS, "templates")
	if err != nil {
		return err
	}

	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			return nil
		}

		if d.IsDir() {
			l.Info("creating directory", zap.String("path", path))
			return makeDirs(path)
		}

		outputFileName := path

		if !strings.HasSuffix(outputFileName, ".tmpl") {
			return nil
		}

		outputFileName = outputFileName[:len(outputFileName)-len(".tmpl")]

		l.Info("creating file from template", zap.String("file", outputFileName))
		return processTemplate(fsys, path, outputFileName, config)
	})
}

func processTemplate(fsys fs.FS, path, outputPath string, config *goblConfig) (err error) {
	var tmpl *template.Template
	var f *os.File
	if tmpl, err = template.ParseFS(fsys, path); err != nil {
		return
	}

	if f, err = os.Create(outputPath); err != nil {
		return
	}

	defer f.Close()

	return tmpl.Execute(f, config)
}

func logError(l *zap.Logger, err error, output []byte) {
	var e *exec.ExitError
	ok := errors.As(err, &e)
	if !ok {
		l.Warn(string(output))
	} else {
		l.Warn(string(e.Stderr))
	}
}

func initModule(cfg *goblConfig) (err error) {
	l := logger.Logger().With(zap.String("module-name", cfg.GoModule))

	var args []string
	var output []byte

	l.Info("initialising Go module:")
	args = []string{"mod", "init", cfg.GoModule}
	if output, err = exec.Command(goExec, args...).Output(); err != nil {
		logError(l, err, output)
		return err
	}

	args = []string{"mod", "tidy"}
	if output, err = exec.Command(goExec, args...).Output(); err != nil {
		logError(l, err, output)
		return err
	}

	return err
}

func initGit(cfg *goblConfig) (err error) {
	l := logger.Logger().With(zap.String("module-name", cfg.GoModule))
	var args []string
	var output []byte

	l.Info("initialising Git")
	args = []string{"init", "-b", "main"}
	if output, err = exec.Command(gitExec, args...).Output(); err != nil {
		logError(l, err, output)
	}

	return err
}

func initCommit(cfg *goblConfig) (err error) {
	l := logger.Logger().With(zap.String("module-name", cfg.GoModule))
	var args []string
	var output []byte

	l.Info("Creating initial commit")
	args = []string{"add", "."}
	if output, err = exec.Command(gitExec, args...).Output(); err != nil {
		logError(l, err, output)
		return err
	}

	args = []string{"commit", "-m", "chore: initial commit"}
	if output, err = exec.Command(gitExec, args...).Output(); err != nil {
		logError(l, err, output)
	}

	return err
}

func makeDirs(dirPaths ...string) error {
	for _, p := range dirPaths {
		if err := os.MkdirAll(p, folderPermissions); err != nil {
			return err
		}
	}

	return nil
}

func initPreCommit(cfg *goblConfig) (err error) {
	l := logger.Logger().With(zap.String("module-name", cfg.GoModule))
	var args []string
	var output []byte

	l.Info("initialising pre-commit")
	args = []string{"install", "-t", "commit-msg"}
	if output, err = exec.Command(preCommitExec, args...).Output(); err != nil {
		logError(l, err, output)
		return err
	}

	return err
}

package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/viper"

	icmd "gitlab.com/gobl/gobl/internal/cmd"
	"gitlab.com/gobl/gobl/pkg/application"
	"gitlab.com/gobl/gobl/pkg/bootstrap"
	"gitlab.com/gobl/gobl/pkg/cmd"
	"gitlab.com/gobl/gobl/pkg/io/fs"
)

const (
	goblPrefix             = "GOBL"
	goblServiceName        = "GOBL_SERVICE_NAME"
	defaultGoblLibraryPath = "gitlab.com/gobl/gobl"
	configDirMask          = 0700
)

var (
	//go:embed config/configuration.yaml
	configFS embed.FS
)

type goblConfigFileProps struct {
	GoProjectAddress string
	GoblLogPath      string
}

func main() {
	cmd.SetRootCmd(icmd.GoblCmd())
	cmd.AddCommand(icmd.NewCmd())
	cmd.AddBootstrapInitFuncs(false, initGoblConfig)
	cmd.SetCliProperties(icmd.GoblUsage, icmd.GoblShortDesc, icmd.GoblLongDesc)

	application.BindEnvVars(goblPrefix)
	app := bootstrap.New()

	cmd.Execute(context.Background(), app, nil)
}

func initGoblConfig() {
	// We are going to check if the gobl configuration file already exists
	// if not we will try to create it, or at the very least set the address
	// for the gobl library so we can `go get` it if a new bootstrap application
	// is being created
	homeDir, err := os.UserHomeDir()
	if err != nil {
		os.Setenv(goblServiceName, defaultGoblLibraryPath)
	}

	goblConfigDir := fmt.Sprintf("%s/.config/gobl", homeDir)
	goblConfigPath := fmt.Sprintf("%s/configuration.yaml", goblConfigDir)

	var exists bool
	if exists, err = fs.FileExists(goblConfigPath); exists && err != nil {
		return
	}

	props := goblConfigFileProps{
		GoProjectAddress: defaultGoblLibraryPath,
		GoblLogPath:      fmt.Sprintf("%s/logs/gobl.log", goblConfigDir),
	}

	if _, err = os.Stat(goblConfigDir); err != nil {
		_ = os.MkdirAll(goblConfigDir, configDirMask)
	}

	var tmpl *template.Template
	var f *os.File

	if tmpl, err = template.ParseFS(configFS, "config/configuration.yaml"); err != nil {
		os.Setenv(goblServiceName, defaultGoblLibraryPath)
		return
	}

	if f, err = os.Create(goblConfigPath); err != nil {
		os.Setenv(goblServiceName, defaultGoblLibraryPath)
		return
	}

	defer f.Close()

	if err = tmpl.Execute(f, props); err != nil {
		os.Setenv(goblServiceName, defaultGoblLibraryPath)
		return
	}

	viper.AddConfigPath(goblConfigDir)
	viper.SetConfigName("configuration")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		os.Setenv(goblServiceName, defaultGoblLibraryPath)
	}

	cmd.SetupLogger()
}

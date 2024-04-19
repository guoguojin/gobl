package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gitlab.com/gobl/gobl/pkg/config"
)

// Version returns the current version of the application
//
//nolint:gochecknoglobals,forbidigo
var Version = &cobra.Command{
	Use:   "version",
	Short: "Current version",
	Long:  "Get the current version number of the service",
	Run: func(*cobra.Command, []string) {
		fmt.Printf("Version: %s\n", viper.GetString(config.VersionKey))
	},
}

// init is used by Cobra to initialise a command
//
//nolint:gochecknoinits
func init() {
	AddCommand(Version)
}

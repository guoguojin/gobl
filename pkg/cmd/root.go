package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"gitlab.com/gobl/gobl/pkg/io/strings"
	"gitlab.com/gobl/gobl/pkg/logger"
	"gitlab.com/gobl/gobl/pkg/service"
)

//nolint:gochecknoglobals
var (
	CfgFile          string
	l                *zap.Logger
	ctx              context.Context
	app              service.Service
	state            service.State
	cobraInitFuncs   []func()
	defaultInit      bool
	Usage            string
	ShortDescription string
	LongDescription  string
)

// SetCliProperties sets the Cobra command CLI properties so that information
// on how to use the application is provided when help is needed.
func SetCliProperties(usage, short, long string) {
	Usage = usage
	ShortDescription = short
	LongDescription = long
}

//nolint:gochecknoglobals
var rootCmd = &cobra.Command{
	Run: startService,
}

//nolint:gochecknoglobals
var overrideRoot *cobra.Command

// SetRootCmd sets an override root command to run instead of the default root command
func SetRootCmd(c *cobra.Command) {
	overrideRoot = c
	overrideRoot.PersistentFlags().StringVar(&CfgFile, "config", "", "configuration file to use for the service")
}

func startService(*cobra.Command, []string) {
	if err := app.Init(ctx, state); err != nil {
		log.Fatal("could not initialise the application", zap.Error(err))
	}

	if app.RunFunction() != nil {
		if err := app.RunFunction()(ctx, state); err != nil {
			l.Error("Command failed", zap.Error(err))
		}

		if err := app.Cleanup(state); err != nil {
			l.Fatal("could not execute cleanup", zap.Error(err))
		}

		return
	}

	l.Info("Starting service. Ctrl-C to terminate")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	incoming := <-signalCh

	l.Warn("Caught signal, terminating", zap.String("signal", incoming.String()))

	if err := app.Cleanup(state); err != nil {
		l.Fatal("could not execute cleanup", zap.Error(err))
	}
}

//nolint:gochecknoinits
func init() {
	l = logger.ConsoleLogger()

	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "configuration file to use for the service")
}

// SetupCobraInit sets the cobra initialisation functions
func SetupCobraInit() {
	if cobraInitFuncs == nil {
		cobraInitFuncs = append(cobraInitFuncs, InitConfig)
	}

	cobra.OnInitialize(cobraInitFuncs...)
}

// AddBootstrapInitFuncs adds initialisation functions that are run by Cobra
// on application initialisation
func AddBootstrapInitFuncs(includeDefaultInit bool, fs ...func()) {
	if includeDefaultInit && !defaultInit {
		defaultInit = true
		// make sure we call the default init function first
		initFns := []func(){InitConfig}
		cobraInitFuncs = append(append(initFns, cobraInitFuncs...), fs...)
		return
	}

	cobraInitFuncs = append(cobraInitFuncs, fs...)
	SetupCobraInit()
}

// InitConfig is the function called to initialise the service configuration file
func InitConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("could not get user home directory", zap.Error(err))
		}

		executable, err := os.Executable()
		if err != nil {
			log.Fatal("could not get the current executable name", zap.Error(err))
		}

		executablePath := strings.SplitAndTrimSpace(executable, string(os.PathSeparator))

		appName := executablePath[len(executablePath)-1]

		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath(fmt.Sprintf("%s/.config/%s", home, appName))
		viper.SetConfigName("configuration")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		//nolint:forbidigo
		fmt.Printf("could not read application configuration file: %s\n\n", err)
	} else {
		l.Debug("using configuration", zap.String("file-path", viper.ConfigFileUsed()))
	}

	SetupLogger()
}

// SetupLogger sets up the logging configuration based on defaults or properties set in the
// application configuration file
func SetupLogger() {
	l = logger.Get(logger.ApplicationLogLevel(), logger.ConfiguredLumberjackLogger())
	zap.ReplaceGlobals(l)
}

// AddCommand allows you to add a sub-command to the root command
func AddCommand(commands ...*cobra.Command) {
	runCmd := rootCmd
	if overrideRoot != nil {
		runCmd = overrideRoot
	}
	runCmd.AddCommand(commands...)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(c context.Context, a service.Service, appState service.State) {
	ctx = c
	app = a
	state = appState

	runCmd := rootCmd

	if overrideRoot != nil {
		runCmd = overrideRoot
	}

	runCmd.Use = Usage
	runCmd.Short = ShortDescription
	runCmd.Long = LongDescription

	// make sure we setup the cobra initialisation properly
	SetupCobraInit()

	if err := runCmd.Execute(); err != nil {
		//nolint:forbidigo
		fmt.Println(err)
		os.Exit(1)
	}
}

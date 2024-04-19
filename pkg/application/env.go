package application

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// there are the default bindings that correspond to the default configuration file
//
//nolint:gochecknoglobals
var bindings = map[string]string{
	"version":         "",
	"service.name":    "",
	"log.filepath":    "",
	"log.level":       "",
	"log.max-size":    "",
	"log.max-backups": "",
	"log.max-age":     "",
	"log.compress":    "",
}

// BindEnvVars binds any environment variables that have been defined with the
// specified viper configuration. The default settings have been predefined so you
// only need to set additional bindings specific to your application.
// The prefix will be added to all defined environment variables. For example,
// if your chosen prefix is myapp, then any environment variable defined for the
// binding will be prefixed as MYAPP_<ENV_VAR>.
// E.g.
//
// SetEnvVarBinding("service.address", "SERVICE_ADDRESS")
//
// will bind the service.address configuration to the MYAPP_SERVICE_ADDRESS
// environment variable.
func BindEnvVars(prefix string) {
	log := zap.L()
	viper.SetEnvPrefix(prefix)
	for k, v := range bindings {
		vars := []string{k}

		if v != "" {
			vars = append(vars, v)
		}

		if err := viper.BindEnv(vars...); err != nil {
			log.Error("BindEnv", zap.String("config", k))
		}
	}
	viper.AutomaticEnv()
}

// SetEnvVarBinding adds an environment variable binding for viper configuration
// This function should be called for each binding you wish to add before calling
// the BindEnvVars function
func SetEnvVarBinding(config, env string) {
	bindings[config] = env
}

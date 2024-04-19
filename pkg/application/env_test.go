package application

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"gitlab.com/gobl/gobl/pkg/config"
)

func TestEnv(t *testing.T) {
	viper.SetConfigName("testconfig")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("could not read test config: %v", err)
	}

	SetEnvVarBinding("dummy.config", "DUMMY_CONFIG")
	BindEnvVars("TEST")

	t.Run("should read from the viper config file if the environment variable doesn't exist", func(t *testing.T) {
		version := config.Get("version").String("")
		assert.Equal(t, "0.1.0", version)
	})

	t.Run("should read from the environment variable if it has been set", func(t *testing.T) {
		assert.NoError(t, os.Setenv("TEST_VERSION", "0.2.0"))
		assert.NoError(t, os.Setenv("TEST_LOG.LEVEL", "ERROR"))

		version := config.Get("version").String("")
		assert.Equal(t, "0.2.0", version)

		logLevel := config.Get("log", "level").String("")
		assert.Equal(t, "ERROR", logLevel)
	})

	t.Run("should use the provided environment variable name if it has been set", func(t *testing.T) {
		assert.NoError(t, os.Setenv("DUMMY_CONFIG", "dummy value"))

		dummyConfig := config.Get("dummy", "config").String("")
		assert.Equal(t, "dummy value", dummyConfig)
	})
}

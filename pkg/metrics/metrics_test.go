package metrics_test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	viper.SetConfigName("test_config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Fatal("could not read test config", zap.Error(err))
	}

	os.Exit(m.Run())
}

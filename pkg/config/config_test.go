package config_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/gobl/gobl/pkg/config"

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

type GoodServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type BadServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type GoodUserConfig struct {
	FirstName string `mapstructure:"first-name"`
	LastName  string `mapstructure:"last-name"`
	Age       int    `mapstructure:"age"`
}

type BadUserConfig struct {
	FirstName string `mapstructure:"first-name"`
	LastName  string `mapstructure:"last-name"`
	Age       int    `mapstructure:"age"`
}

func (c GoodServerConfig) Validate() error {
	return nil
}

func (c GoodUserConfig) Validate() error {
	return nil
}

func (c BadServerConfig) Validate() error {
	return errors.New("bad server config")
}

func (c BadUserConfig) Validate() error {
	return errors.New("bad user config")
}

var (
	defGoodSrvConf = GoodServerConfig{
		Host: "localhost",
		Port: 8080,
	}

	defGoodUsrConf = GoodUserConfig{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	defBadSrvConf = BadServerConfig{
		Host: "localhost",
		Port: 8080,
	}

	defBadUsrConf = BadUserConfig{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
)

func TestReadConfigFromFile(t *testing.T) {
	t.Run("Should read good configs from file", testGoodConfigs)
	t.Run("Should return default config and error if config not found", testConfigNotFound)
	t.Run("Should return default config and error if config is invalid", testInvalidConfig)
}

func testGoodConfigs(t *testing.T) {
	var (
		srvConf GoodServerConfig
		usrConf GoodUserConfig
	)

	wantSrvConf := GoodServerConfig{
		Host: "some-host",
		Port: 1234,
	}

	wantUsrConf := GoodUserConfig{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       42,
	}

	err := config.ReadConfigFromFile("server-config", &srvConf, defGoodSrvConf)
	assert.NoError(t, err)

	assert.Equal(t, wantSrvConf, srvConf)

	err = config.ReadConfigFromFile("user-config", &usrConf, defGoodUsrConf)
	assert.NoError(t, err)

	assert.Equal(t, wantUsrConf, usrConf)
}

func testConfigNotFound(t *testing.T) {
	var (
		srvConf GoodServerConfig
		usrConf GoodUserConfig
	)

	err := config.ReadConfigFromFile("some-server-config", &srvConf, defGoodSrvConf)
	assert.Error(t, err)

	assert.Equal(t, defGoodSrvConf, srvConf)

	err = config.ReadConfigFromFile("some-user-config", &usrConf, defGoodUsrConf)
	assert.Error(t, err)

	assert.Equal(t, defGoodUsrConf, usrConf)
}

func testInvalidConfig(t *testing.T) {
	var (
		srvConf BadServerConfig
		usrConf BadUserConfig
	)

	err := config.ReadConfigFromFile("server-config", &srvConf, defBadSrvConf)
	assert.Error(t, err)

	assert.Equal(t, defBadSrvConf, srvConf)

	err = config.ReadConfigFromFile("user-config", &usrConf, defBadUsrConf)
	assert.Error(t, err)

	assert.Equal(t, defBadUsrConf, usrConf)
}

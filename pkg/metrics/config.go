package metrics

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"gitlab.com/gobl/gobl/pkg/service"
)

type Config struct {
	Host                        string        `mapstructure:"host"`
	Port                        int           `mapstructure:"port"`
	Path                        string        `mapstructure:"path"`
	Enabled                     bool          `mapstructure:"enabled"`
	HTTPServerTimeout           time.Duration `mapstructure:"http-server-timeout"`
	HTTPServerReadHeaderTimeout time.Duration `mapstructure:"http-server-read-header-timeout"`
}

const (
	MinPathLength = 2
	MaxPathLength = 64
	MinPort       = 1
)

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required, validation.Min(MinPort), validation.Max(service.MaxPort())),
		validation.Field(&c.Path, validation.Required, validation.Length(MinPathLength, MaxPathLength)),
		validation.Field(&c.HTTPServerTimeout, validation.Required),
		validation.Field(&c.HTTPServerReadHeaderTimeout, validation.Required),
	)
}

func DefaultConfig() Config {
	return Config{
		Port:                        2022,
		Path:                        "/metrics",
		Enabled:                     false,
		HTTPServerTimeout:           time.Minute,
		HTTPServerReadHeaderTimeout: time.Minute,
	}
}

package metrics_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/gobl/gobl/pkg/metrics"

	"github.com/stretchr/testify/assert"
)

func TestConfig_DefaultConfig(t *testing.T) {
	t.Run("DefaultConfig should return the default configuration", func(t *testing.T) {
		want := metrics.Config{
			Port:                        2022,
			Path:                        "/metrics",
			Enabled:                     false,
			HTTPServerTimeout:           time.Minute,
			HTTPServerReadHeaderTimeout: time.Minute,
		}
		got := metrics.DefaultConfig()
		assert.Equal(t, want, got)
	})

	t.Run("Default configuration should pass validation", func(t *testing.T) {
		config := metrics.DefaultConfig()
		err := config.Validate()
		require.NoError(t, err)
	})
}

func TestConfig_Validate(t *testing.T) {
	t.Run("Validate should fail if Port is less than MinPort", testValidatePortLessThanMin)
	t.Run("Validate should fail if Port is greater than service.MaxPort", testValidatePortMoreThanMax)
	t.Run("Validate should fail if Path is not set", testValidatePathNotSet)
	t.Run("Validate should fail if Path is too long", testValidatePathTooLong)
	t.Run("Validate should fail if HTTPServerTimeout is not set", testValidateServerTimeout)
	t.Run("Validate should fail if HTTPServerReadHeaderTimeout is not set", testValidateServerHeaderReadTimeout)
}

func testValidatePortLessThanMin(t *testing.T) {
	c := metrics.DefaultConfig()
	c.Port = -1

	err := c.Validate()
	assert.Error(t, err)
	assert.Equal(t, "Port: must be no less than 1.", err.Error())

	c.Port = 0

	err = c.Validate()
	assert.Error(t, err)
	assert.Equal(t, "Port: cannot be blank.", err.Error())
}

func testValidatePortMoreThanMax(t *testing.T) {
	c := metrics.DefaultConfig()
	c.Port = 100000

	err := c.Validate()
	assert.Error(t, err)
	assert.Equal(t, "Port: must be no greater than 65535.", err.Error())
}

func testValidatePathNotSet(t *testing.T) {
	c := metrics.DefaultConfig()
	c.Path = ""

	err := c.Validate()
	assert.Error(t, err)
	assert.Equal(t, "Path: cannot be blank.", err.Error())
}

func testValidatePathTooLong(t *testing.T) {
	c := metrics.DefaultConfig()
	c.Path = "/really-really-really-really-really-really-really-really-really-really-really-really-long-endpoint-to-be-using-for-something-like-metrics"

	err := c.Validate()
	assert.Error(t, err)
	assert.Equal(t, "Path: the length must be between 2 and 64.", err.Error())
}

func testValidateServerTimeout(t *testing.T) {
	c := metrics.DefaultConfig()
	c.HTTPServerTimeout = 0
	err := c.Validate()
	assert.Error(t, err)
	assert.Equal(t, "HTTPServerTimeout: cannot be blank.", err.Error())
}

func testValidateServerHeaderReadTimeout(t *testing.T) {
	c := metrics.DefaultConfig()
	c.HTTPServerReadHeaderTimeout = 0
	err := c.Validate()
	assert.Error(t, err)
	assert.Equal(t, "HTTPServerReadHeaderTimeout: cannot be blank.", err.Error())
}

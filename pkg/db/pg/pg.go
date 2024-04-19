package pg

import "fmt"

// Configuration holds configuration information necessary for connecting to
// a Postgres database
type Configuration struct {
	host     string
	port     int
	user     string
	password string
	database string
	sslMode  string
}

// PgConnectionString returns the Postgres connection string using the information
// provided in the configuration
func (c Configuration) PgConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.user, c.password, c.host, c.port, c.database, c.sslMode)
}

// NewConfiguration creates a Configuration instance
func NewConfiguration(host string, port int, user, password, database, sslMode string) Configuration {
	return Configuration{
		host, port, user, password, database, sslMode,
	}
}

// Host returns the host specified in the configuration
func (c Configuration) Host() string {
	return c.host
}

// Port returns the port specified in the configuration
func (c Configuration) Port() int {
	return c.port
}

// User returns the user specified in the configuration
func (c Configuration) User() string {
	return c.user
}

// Password returns the password specified in the configuration
func (c Configuration) Password() string {
	return c.password
}

// Database returns the password specified in the configuration
func (c Configuration) Database() string {
	return c.database
}

// SSLMode returns the SSL mode specified in the configuration
func (c Configuration) SSLMode() string {
	return c.sslMode
}

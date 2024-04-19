package testing

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"testing"
	"time"

	// import pgx drivers
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pressly/goose/v3"
)

const (
	postgresPortID = "5432/tcp"
)

type PgTestScaffold struct {
	Scaffold
	DB                 *sql.DB
	PgConnectionString string
}

type PgTestConfig struct {
	User          string
	Password      string
	DBName        string
	MigrationsFS  fs.FS
	MigrationsDir string
}

// SetupPg creates a test scaffold with a Dockerized Postgres database server,
// and executes the migration scripts to create the database structure needed for your
// tests. This allows you to execute tests against a real database instead of using a mocked
// database.
// This method uses the background context by default
//
// Deprecated: Use SetupPgContext and pass in an appropriate context instead.
func SetupPg(t *testing.T, config PgTestConfig, options ...Option) (PgTestScaffold, error) {
	return SetupPgContext(context.Background(), t, config, options...)
}

// SetupPgContext creates a test scaffold with a Dockerized Postgres database server,
// and executes the migration scripts to create the database structure needed for your
// tests. This allows you to execute tests against a real database instead of using a mocked
// database.
func SetupPgContext(ctx context.Context, t *testing.T, config PgTestConfig, options ...Option) (PgTestScaffold, error) {
	t.Helper()
	t.Log("Setting up constructs for tests")

	opts, err := applyPostgresOptions(options...)
	if err != nil {
		return PgTestScaffold{}, err
	}

	var setup PgTestScaffold
	pool, err := dockertest.NewPool("")
	if err != nil {
		return setup, fmt.Errorf("creating docker test pool: %w", err)
	}

	defaultEnvs := []string{
		fmt.Sprintf("POSTGRES_PASSWORD=%s", config.Password),
		fmt.Sprintf("POSTGRES_USER=%s", config.User),
		fmt.Sprintf("POSTGRES_DB=%s", config.DBName),
		"listen_addresses = '*'",
	}

	env := applyDefaultEnvVars(defaultEnvs, opts.env)

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: opts.repository,
		Tag:        opts.tag,
		Env:        env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return setup, fmt.Errorf("creating Postgres docker container: %w", err)
	}

	setup.Port = resource.GetPort(postgresPortID)

	setup.PgConnectionString = fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		config.User,
		config.Password,
		resource.GetHostPort(postgresPortID),
		config.DBName,
	)

	const maxSeconds = 120

	t.Logf("Connecting to database on url: %s\n", setup.PgConnectionString)
	if err := resource.Expire(maxSeconds); err != nil {
		return setup, err
	}

	pool.MaxWait = maxSeconds * time.Second
	if err := pool.Retry(func() error {
		var dbErr error

		setup.DB, dbErr = sql.Open("pgx", setup.PgConnectionString)
		if dbErr != nil {
			t.Logf("Cannot connect to database: %v\n", dbErr)
			return dbErr
		}

		dbErr = setup.DB.Ping()

		if dbErr != nil {
			t.Logf("could not ping database: %v", dbErr)
		}

		return dbErr
	}); err != nil {
		return setup, fmt.Errorf("connecting to test database: %w", err)
	}

	t.Log("Running database migrations")
	goose.SetBaseFS(config.MigrationsFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return setup, fmt.Errorf("setting database dialect: %w", err)
	}
	if err := goose.RunContext(ctx, "up", setup.DB, config.MigrationsDir); err != nil {
		return setup, fmt.Errorf("running database migrations: %w", err)
	}

	t.Log("Database migrations completed successfully")
	setup.Teardown = func(t *testing.T) error {
		t.Log("Tearing down test constructs")
		if err := setup.DB.Close(); err != nil {
			return fmt.Errorf("closing connection to test database: %w", err)
		}

		return pool.Purge(resource)
	}

	return setup, nil
}

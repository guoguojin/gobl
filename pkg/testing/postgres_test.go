package testing

import (
	"context"
	"embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

//nolint:govet
func TestSetupPostgres(t *testing.T) {
	if os.Getenv("GITLAB_CI") != "" {
		t.Skip("skipping Setup Postgres tests on GitLab CI")
	}

	config := PgTestConfig{
		User:          "postgres",
		Password:      "postgres",
		DBName:        "postgres",
		MigrationsFS:  migrationsFS,
		MigrationsDir: "migrations",
	}

	scaffold, err := SetupPg(t, config)
	defer func() {
		if err := scaffold.Teardown(t); err != nil {
			t.Logf("could not tear down postgres: %v", err)
		}
	}()

	require.NoError(t, err, "setting up Postgres scaffold")
	assert.NotEqual(t, "", scaffold.Port)

	rows, err := scaffold.DB.QueryContext(context.Background(), "select user_name, age from users")
	require.NoError(t, err, "querying database")
	defer func() {
		if err := rows.Close(); err != nil {
			t.Logf("could not close rows: %v", err)
		}
	}()

	assert.NoError(t, rows.Err())

	rowCount := 0
	type User struct {
		Username string
		Age      int
	}

	users := []User{
		{
			Username: "John Doe",
			Age:      26,
		},
		{
			Username: "Jane Doe",
			Age:      24,
		},
	}

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Username, &user.Age)

		assert.NoError(t, err)
		want := users[rowCount]

		assert.Equal(t, want.Username, user.Username)
		assert.Equal(t, want.Age, user.Age)

		rowCount++
	}

	assert.Equal(t, len(users), rowCount)
}

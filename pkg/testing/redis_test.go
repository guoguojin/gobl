package testing

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:govet
func TestSetupRedis(t *testing.T) {
	if os.Getenv("GITLAB_CI") != "" {
		t.Skip("skipping Setup Redis tests on GitLab CI")
	}

	timeoutSeconds := 10
	scaffold, err := SetupRedis(t, timeoutSeconds)
	assert.NoError(t, err)
	assert.NotEqual(t, "", scaffold.Port)

	defer func() {
		if err := scaffold.Teardown(t); err != nil {
			t.Logf("could not teardown Redis: %v", err)
		}
	}()

	err = scaffold.Rdb.Set(context.Background(), "Hello", "World", 0).Err()
	assert.NoError(t, err)

	v, err := scaffold.Rdb.Get(context.Background(), "Hello").Result()
	assert.NoError(t, err)
	assert.Equal(t, "World", v)
}

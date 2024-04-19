package testing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest/v3"
)

const (
	redisPort  = "6379/tcp"
	maxSeconds = 120
)

type RedisScaffold struct {
	Scaffold
	Rdb *redis.Client
}

// SetupRedis creates a test scaffold for with a Dockerized Redis server.
// This allows you to execute tests against a real database instead of using a mocked
// service.
func SetupRedis(t *testing.T, timeoutSeconds int, options ...Option) (RedisScaffold, error) {
	var s RedisScaffold
	t.Helper()

	opts, err := applyRedisOptions(options...)

	if err != nil {
		return s, err
	}

	t.Log("Setting up constructs for tests")
	pool, err := dockertest.NewPool("")
	if err != nil {
		return s, fmt.Errorf("creating docker test pool: %w", err)
	}

	resource, err := pool.Run(opts.repository, opts.tag, opts.env)
	if err != nil {
		return s, fmt.Errorf("creating docker redis container: %w", err)
	}

	if err := resource.Expire(maxSeconds); err != nil {
		return s, fmt.Errorf("setting redis container expiry timeout: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeoutSeconds))
	defer cancel()

	s.Port = resource.GetHostPort(redisPort)

	pool.MaxWait = maxSeconds * time.Second
	if err := pool.Retry(func() error {
		s.Rdb = redis.NewClient(&redis.Options{
			Addr: s.Port,
		})

		dbErr := s.Rdb.Ping(ctx).Err()

		if dbErr != nil {
			t.Logf("could not ping redis server: %v", dbErr)
		}

		return dbErr
	}); err != nil {
		return s, fmt.Errorf("connecting to Redis: %w", err)
	}

	s.Teardown = func(t *testing.T) error {
		t.Logf("Tearing down redis scaffold")
		if err := s.Rdb.Close(); err != nil {
			return fmt.Errorf("closing redis connection: %w", err)
		}

		return pool.Purge(resource)
	}

	return s, nil
}

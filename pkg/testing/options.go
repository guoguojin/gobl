package testing

import "strings"

const (
	splitMax = 2
)

type options struct {
	repository string
	tag        string
	env        []string
}

type Option func(options *options) error

// WithRepository allows you to replace the default docker image repository
// with the one you want to use.
//
// E.g. For Postgres, the default repository is `postgres`, the official docker image from Docker Hub.
// If you want to use `circleci/postgres` instead, you can specify it with this option.
func WithRepository(repository string) Option {
	return func(opts *options) error {
		opts.repository = repository
		return nil
	}
}

// WithTag allows you to specify the docker image tag to use.
// By default, the tag `latest` is used, this option will allow you
// to specify the tag of the repository you are using.
func WithTag(tag string) Option {
	return func(opts *options) error {
		opts.tag = tag
		return nil
	}
}

// WithEnv allows you to specify the environment variables to set on the container
// when it is spun up. These strings are the same as you would pass to the `dockertest` API
// when using it directly.
func WithEnv(env ...string) Option {
	return func(opts *options) error {
		opts.env = env
		return nil
	}
}

func applyPostgresOptions(opts ...Option) (options, error) {
	o := defaultPostgresOptions()
	for _, opt := range opts {
		err := opt(&o)
		if err != nil {
			return defaultPostgresOptions(), err
		}
	}

	return o, nil
}

func applyRedisOptions(opts ...Option) (options, error) {
	o := defaultRedisOptions()
	for _, opt := range opts {
		err := opt(&o)
		if err != nil {
			return defaultRedisOptions(), err
		}
	}

	return o, nil
}

func defaultPostgresOptions() options {
	return options{
		repository: "postgres",
		tag:        "latest",
		env:        nil,
	}
}

func defaultRedisOptions() options {
	return options{
		repository: "redis",
		tag:        "latest",
		env:        nil,
	}
}

func applyDefaultEnvVars(defaults []string, overrides []string) []string {
	envs := make([]string, len(overrides))
	copy(envs, overrides)

	for _, def := range defaults {
		defVar, err := getEnvironmentVariableName(def)

		if err != nil {
			// This is a library to help testing. If we have received invalid data
			// then we should just panic.
			panic(err)
		}

		found := false

		for _, override := range overrides {
			overrideVar, err := getEnvironmentVariableName(override)
			if err != nil {
				// Again, we should just panic as this library should only be used for testing.
				panic(err)
			}

			if defVar == overrideVar {
				found = true
				break
			}
		}

		if found {
			continue
		}

		envs = append(envs, def)
	}

	return envs
}

func getEnvironmentVariableName(envVar string) (string, error) {
	v := strings.Split(envVar, "=")
	if len(v) != splitMax {
		return envVar, ErrInvalidEnvironmentVariable
	}

	return v[0], nil
}

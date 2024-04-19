# GOBL: Go Boostrap Library

This is a bootstrap library for creating services so that common tasks of setting up the scaffolding on projects does not need to be
re-implemented each time, and we can spend our time working on just the business logic.

## Third party libraries

The third party libraries we are using are:

| Library     | Description                                    | Link                                                             |
| ----------- | ---------------------------------------------- | ---------------------------------------------------------------- |
| Cobra       | Library for creating CLI applications          | [https://github.com/spf13/cobra](https://github.com/spf13/cobra) |
| Viper       | Library for application configuration          | [https://github.com/spf13/viper](https://github.com/spf13/viper) |
| Zap         | Library for Structured, pluggable logging      | [https://github.com/uber-go/zap](https://github.com/uber-go/zap) |

<!-- markdownlint-disable MD010 -->
## Pre-requisites

This library requires Go version 1.16+

## Installation

To install the `gobl` CLI

```bash
go install gitlab.com/gobl/gobl/cmd/gobl
```

If you want to use the Bootstrap libraries on it's own without the CLI tool, you can use

```bash
go get gitlab.com/gobl/gobl
```

## Usage

To use this library, create a new project and initialise your new Go project

```bash
gobl new <go-module> [flags]
```

where `<go-module>` is the name of your new Go module (the same thing you would pass to `go mod init`).

This will create your project folder, setup the folder structure for your project and add a few files that you may regularly have
or would need to set up if you were starting from scratch. It also checks the current folder and all parent folders up to your
home folder to check that you're not creating your project underneath and existing git project.

To override this behaviour, you can pass the `--ignore-git`. To get more information about what flags you can pass the new command
you can use the `--help` flag.

```bash
gobl new --help
```

### Application Package

The prebuilt `application` package contains everything you need to get started. You can simply add your state object, initialisation
and clean up functions, then execute it to have a server process running. To stop the server, use `Ctrl-C` to gracefully stop the application.

```go
package main

import (
	"context"

	"gitlab.com/gobl/gobl/pkg/bootstrap"
	"gitlab.com/gobl/gobl/pkg/cmd"

)

const (
	usage = "myApp"
	short = "short description of my application"
	long = "long description and help for my application"
)

func main() {
	// create the application
	app := bootstrap.New()
	// optional if you want to provide information at the command line
	cmd.SetCliProperties(usage, short, long)
	// Now start the server it will do anything because we haven't defined
	// anything for it to do. You can do that using the init functions, i.e.
	// create a HTTP router and define routes and start it in a separate go routine
	// then use the cleanaup functions to make sure that everything is stopped
	// gracefully before the application exits
	cmd.Execute(context.Background(), app, nil)
}
```

### Init and cleanup functions

The bootstrap allows you to define and add init and cleanup functions that will be run sequentially in the order they were
added by the bootstrap.

The init and cleanup functions are defined as

```go
type InitFunc func(context.Context, service.StateStore) error
type CleanupFunc func(service.StateStore) error
```

To use the bootstrap, define your initialisation and cleanup functions and add them to the application.

### State

To allow the application to store state, you can implement the `service.State` interface. This will allow you to utilise
whatever state store is necessary for your service whether that is in-memory storage or using a database like redis or distributed
service like Etcd or Consul.

### Properties

Additionally, you can create properties that are required for your service by using the `property.<Type>Property` function which creates a `property.Property`
object, e.g. `property.StringProperty` to create a string property. With a `property.Property` object, you can call one of the functions to get the value of
the property as the required type, or return a default value you specify if it's not possible.

#### Example

```go
myProp := property.IntProperty(100) // this will create a property with a value 100
st := myProp.String("unspecified") // this will return the value of the property as a string or "unspecified" if it's not possible
i := myProp.Int(100) // this will return the value of the property as an int or int(100) if it's not possible
i32 := myProp.Int32(100) // this will return the value of the property as an int32 or int32(100) if it's not possible
i64 := myProp.Int64(100) // this will return the value of the property as an int64 or int64(100) if it's not possible
f64 := myProp.Float64(100) // this will return the value of the property as a float64 or float64(100) if it's not possible
f32 := myProp.Float32(100) // this will return the value of the property as a float32 or float32(100) if it's not possible
```

You can use the properties to set the `usage`, `short description` and `long description` that are used to display help information
for cobra commands in the main application.

```go
package main

import (
	"gitlab.com/gobl/gobl/pkg/bootstrap"
	"gitlab.com/gobl/gobl/pkg/property"
)

var (
	usage = "my-app"
	short = "a short description"
	long = "a longer description with extra informative information"
)

func main() {
    app := bootstrap.New().
            AddProperty(property.KeyUsage, property.StringProperty(property.KeyUsage, usage)).
            AddProperty(property.KeyShortDesc, property.StringProperty(property.KeyShortDesc, short)).
            AddProperty(property.KeyLongDesc, property.StringProperty(property.KeyLongDesc, long))
}
```

Alternatively:

```go
package main

import (
	"gitlab.com/gobl/gobl/pkg/bootstrap"
	"gitlab.com/gobl/gobl/pkg/property"
)

var (
	usage = "my-app"
	short = "a short description"
	long = "a longer description with extra informative information"
)

func main() {
	app := bootstrap.New()
	app.SetProperties(property.Properties{
		property.KeyUsage: property.StringProperty(property.KeyUsage, usage),
		property.KeyShortDesc: property.StringProperty(property.KeyShortDesc, short),
		property.KeyLongDesc: property.StringProperty(property.KeyLongDesc, long),
    })
}
```

### Configuration file

The bootstrap expects a configuration file with the following settings as a minimum:

```yaml
version: 0.1.0
service:
    name: bootstrap
log:
    filepath: ./logs/bootstrap.log
    level: DEBUG
    max-size: 100
    max-backup: 5
    max-age: 30
    compress: false
```

The configuration file must be called configuration.&lt;ext&gt; where ext is any format supported by viper.

You can add your own configuration to the file and access them using viper.

#### Binding with environment variables

You can bind any settings in the configuration file to environment variables by calling the `service.SetEnvVarBinding` function for each configuration,
before calling the `service.BindEnvVars`.

```go
package main

import "gitlab.com/gobl/gobl/pkg/application"

var envVars = map[string]string{
	"service.address":         "SERVICE_ADDRESS",
	"request.default-timeout": "REQUEST_DEFAULT_TIMEOUT",
}

func main() {
	for k, v := range envVars {
		application.SetEnvVarBinding(k, v)
	}

	application.BindEnvVars("myapp")
}
```

On your server, you can set the environment variables `MYAPP_SERVICE_ADDRESS` and `MYAPP_REQUEST_DEFAULT_TIMEOUT` to override any configuration
found in the configuration file.

### CLI commands

To add your own CLI commands, you can just create a command, and add them before calling the `cmd.Execute()` function. For example:

```go
package main

// omitted for clarity

var helloCmd = &cobra.Command{
    Use: "hello",
    Short: "Says hello",
    Long: "The obligatory hello world function",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello, World!")
    },
}

func main() {
    cmd.AddCommand(helloCmd)
    cmd.Execute(New())
}
```

### Run Function

If you want the main application to perform a task and then exit immediately without running as a server, you can set the RunFunction of the Application. The RunFunc type
is defined as:

```go
package service

import (
	"context"
	"gitlab.com/gobl/gobl/pkg/application"
	"gitlab.com/gobl/gobl/pkg/service"
)

type RunFunc func(context.Context, service.State) error
```

Define your function as you would any other Go function and pass it to the Application using `WithRunFunc`

```go
package main

import (
	"context"

	"gitlab.com/gobl/gobl/pkg/application"
	"gitlab.com/gobl/gobl/pkg/bootstrap"
	"gitlab.com/gobl/gobl/pkg/service"
	"gitlab.com/gobl/gobl/pkg/cmd"
)

func myInitFunc(ctx context.Context, state service.State) error {
	return nil
}

func myCleanupFunc(state service.State) error {
	return nil
}

func myRunFunc(ctx context.Context, state service.State) error {
	// Your function defined here
	return nil
}

// state you want to store for your application
type appState struct {
	// we must use composition and add the state.State interface here
	// so we can pass this state where it's needed
	service.State
}

func main() {
    app := bootstrap.New().
        AddInitFunc(myInitFunc).
        AddCleanupFunc(myCleanupFunc).
        WithRunFunc(myRunFunc)

    st := appState{}

    cmd.Execute(context.Background(), app, st)
}

```

With the RunFunc function, the application will perform the initialisation, execute the run function and then perform the cleanup.

If the RunFunc function is not defined, then the application will run the initialisation and wait for an interrupt signal to stop the
application. Once it receives the interrupt signal, it will perform the cleanup and exit.

### Overriding the default application

You can override the default application by specifying your own RootCmd

```go
package main

import (
	"context"
	"fmt"

	"gitlab.com/gobl/gobl/pkg/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func main() {
	// to specify that you want to use a viper config file, you need to add a flag to your override command
	// and set the config file for the root
	rootCmd..PersistentFlags().StringVar(&cmd.CfgFile, "config", "", "configuration file to use for the service")
	cmd.SetRootCmd(rootCmd)
	cmd.Execute(context.Background(), nil, nil)
}
```

### Bootstrap initialisation function

By default, a bootstrap application will initialise a [Viper](https://github.com/spf13/viper) configuration file using either the configuration
file passed in using the --config flag, or by searching in the current working directory, a config sub-folder of the current working directory
or in the user's $HOME/.config/&lt;application name&gt; folder for a file called configuration. If not found, the application will report an error
reporting that it could not find the configuration file.

If you want to add additional initialisation functions that cobra will call to initialise the bootstrap before initialising your application,
you can use the `cmd.AddBootstrapInitFuncs` function.

With the `cmd.AddBootstrapInitFuncs` you can also override whether or not to use the default initialisation function mentioned above by passing a
boolean as the first parameter to the function. If `false`, then only the initialisation functions you specify will be used to initialise the
bootstrap. If `true`, then the default initialisation functions will be called first before the rest of your bootstrap initialisation functions.

If you want to use the default initialisation function, but do not want it run first, rather run in the order you want it run, then you can set
the first parameter for `cmd.AddBootstrapInitFuncs` to false, and then include the `cmd.InitConfig` function in your list of initialisation
functions in the order you want.

#### Examples

##### Add additional bootstrap initialisation function

```go
package main

// Removed for brevity

func myInitFunc() {
	// your initialisation code goes here
}


func main() {
	cmd.AddBootstrapInitFuncs(true, myInitFunc)
}
```

##### Override default bootstrap function

```go
package main

// Removed for brevity

func myInitFunc() {
	// Your initialisation code goes here
}

func main() {
	cmd.AddBootstrapInitFuncs(false, myInitFunc)
}

```

##### Override default bootstrap function but call it in the order you want

```go
package main

// Removed for brevity

func myFirstInitFunc() {
	// Your initialisation code goes here
}

func mySecondInitFunc() {
	// your initialisation code goes here
}

func main() {
	cmd.AddBootstrapInitFuncs(false, myInitFunc, cmd.InitConfig, mySecondInitFunc)
}

```

## Configuration

To make accessing configuration easier, a configuration wrapper function is available in the `gitlab.com/gobl/gobl/pkg/config`
package.

### Configuration Example

Instead of doing something like this:

```go
    serviceName := "Default Service Name"

    if viper.IsSet("service.name") {
    	serviceName = viper.Get("service.name")
    }
```

We can use:

```go
    serviceName := config.Get("service", "name").String("Default Service Name")
```

### Supported Data Types

The following table contains the translation between the viper function signatures and the config functions we have defined.

| Viper Function                   | Config Function                                          | Return Data Type       |
| -------------------------------- | -------------------------------------------------------- | ---------------------- |
| viper.Get(string)                | config.Get(...string).Value(interface{})                 | interface{}            |
| viper.GetBool(string)            | config.Get(...string).Bool(bool)                         | bool                   |
| viper.GetFloat64(string)         | config.Get(...string).Float64(float64)                   | float64                |
| viper.GetInt(string)             | config.Get(...string).Int(int)                           | int                    |
| viper.GetInt8(string)            | config.Get(...string).Int8(int)                          | int8                   |
| viper.GetInt16(string)           | config.Get(...string).Int16(int)                         | int16                  |
| viper.GetInt32(string)           | config.Get(...string).Int32(int)                         | int32                  |
| viper.GetInt64(string)           | config.Get(...string).Int64(int)                         | int64                  |
| viper.GetUint(string)            | config.Get(...string).Uint(int)                          | uint                   |
| viper.GetUint8(string)           | config.Get(...string).Uint8(int)                         | uint8                  |
| viper.GetUint16(string)          | config.Get(...string).Uint16(int)                        | uint16                 |
| viper.GetUint32(string)          | config.Get(...string).Uint32(int)                        | uint32                 |
| viper.GetUint64(string)          | config.Get(...string).Uint64(int)                        | uint64                 |
| viper.GetString(string)          | config.Get(...string).String(string)                     | string                 |
| viper.GetStringMap(string)       | config.Get(...string).StringMap(map[string]interface{})  | map[string]interface{} |
| viper.GetStringMapString(string) | config.Get(...string).StringMapString(map[string]string) | map[string]string      |
| viper.GetStringSlice(string)     | config.Get(...string).StringSlice([]string)              | []string               |
| viper.GetTime(string)            | config.Get(...string).Time(time.Time)                    | time.Time              |
| viper.GetDuration(string)        | config.Get(...string).Duration(time.Duration)            | time.Duration          |

config.Get takes a variadic string parameter that lays out the path of the configuration you need to retrieve.
The following type method takes a single parameter that is the default value, which will be returned if the
configuration is not available in the configuration file.

### Examples Folder

The `examples` folder contains some examples of how to use the bootstrap and implementing a simple state store for your application.

## Finite State Machine

A package `gitlab.com/gobl/gobl/pkg/fsm` is available with a simple framework for creating and running finite state machines.
An example of how to use the Finite State Machine is available in the `examples` folder under the turnstile project. This project creates a simple
turnstile that has two states, Locked and Unlocked and two events, insert coin and push.

Each state will act differently depending on which event they receive. When locked, pushing the turnstile will
result in an error asking for a coin, while inserting a coin will transition to the Unlocked state.

When unlocked, pushing the turnstile will transition you back to the Locked state, and inserting a coin will result
in an error telling you it has returned your coin.

The State interface is defined as:

```go
package fsm

import (
	"github.com/google/uuid"
)

type State interface {
	// ID is the unique identifier of the state so different states of the same nature
	// can be distinguishable
	ID() uuid.UUID
	// Description of the state
	Description() string
	// Execute processes the event that is passed to is
	Execute(Event) error
	// Next checks the current state and determines what state it should transition to
	Next() State
	// WithTransitions sets the transitions that are supported by each state
	WithTransitions(...Transition) State
}
```

While the Event interface is defined as:

```go
package fsm

import (
	"github.com/google/uuid"
)

type Event interface {
	// ID is a unique identifier for the event
	ID() uuid.UUID
	// Source is a unique identifier used to determine where the event came from
	Source() string
	// Name of the event
	Name() string
	// Timestamp is the time of the event as nanoseconds past epoch
	Timestamp() int64
}

```

Once you have defined your own states and events implementing these interfaces, you can create a state machine to pass
your events to.

```go
machine, errCh := fsm.New(uuid.New(), "Name of state machine", myInitialState)
```

The new function returns a state machine, and a channel where it will publish errors that occur within the state machine.

Create a channel where events will be passed to the state machine, and a context that can be cancelled to terminate the machine
if necessary, then run in a separate go routine using the state machines Run function.

```go
machineCtx, cancel := context.WithCancel(context.Background())
eventCh := make(chan fsm.Event)
cleanupFunc := func() error {
	// Do cleanup code here
}
go machine.Run(machineCtx, eventCh, cleanupFunc)
```

The example also provides an example of how to model transitions from one state to another, using the `Transition` struct and the
`CheckFn` and `NextFn` function types.

They are defined as:

```go
type CheckFn func (fsm.State) bool
type NextFn func (fsm.State) fsm.State
```

<!-- markdownlint-enable MD010 -->

## Testing

In order to test code that needs to read/write from an externals service, for example Postgres, or Redis; we use the
[Ory Docker Test Library](https://github.com/ory/dockertest) to start containerised applications using Docker.

The `testing` package contains some utility functions that will take care of the setup and tear down for the test container.

### Redis

```go
package mypackage

import (
	"context"
	"testing"

	gobltesting "gitlab.com/gobl/gobl/pkg/testing"
)

func MyRedisTest(t *testing.T) {
	timeoutSeconds := 10
	// SetupRedis will create a redis container and return a testing scaffold containing the redis connection
	// the port information and a tear down function that will stop and destroy the container when the test is
	// finished
	scaffold, err := gobltesting.SetupRedis(t, timeoutSeconds)

	if err != nil {
		t.Fatal(err)
    }

	// make sure you tear down the scaffold when the test is done
	defer scaffold.Teardown()

	// Use the Rdb client to execute commands against the Redis database
	err := scaffold.Rdb.Set(context.Background, "mykey", "myvalue").Err()

	if err != nil {
		t.Fatal(err)
    }
}
```

### Postgres

In order to test Postgres we need to do a little preparation by creating the database migration scripts we need for our project.
To do this we use the [goose](https://github.com/pressly/goose) library to manage our migration scripts.

Once you have used goose to create your migration scripts, you can then pass the relative path to the folder containing your
migration scripts for goose to execute.

When the test scaffolding creates the postgres container, it will then execute the migrations bringing your database into sync ready
for you to run your tests against it.

```go
package mypackage

import (
	"context"
	"testing"

	gobltesting "gitlab.com/gobl/gobl/pkg/testing"
)

func TestPostgres(t *testing.T) {
	// SetupPg launches a Postgres Docker container using the username, password and database name you provide.
	// the migrations folder contains migrations to initialise and build the database structure you need for the test
	scaffold, err := SetupPg(t, "postgres", "postgres", "postgres", "./migrations")

	if err != nil {
		t.Fatal(err)
    }

	// Remember to tear down the scaffold when the test is done to stop and destroy the docker container
	defer scaffold.Teardown()

	// Example query against our database
	rows, err := scaffold.DB.QueryContext(context.Background(), "select user_name, age from users")

	if err != nil {
		t.Fatal(err)
    }

	// ensure we close our rows object when we're finished
	defer rows.Close()

	rowCount := 0

	// This is an example struct that we're going to extract the data from the database into
	type User struct {
		Username string
		Age      int
	}

	// the expected contents of the data we set up for the test
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

	// iterate through the rows to validate the data
	for rows.Next() {
		var user User

		err = rows.Scan(&user.Username, &user.Age)

		want := users[rowCount]

		// Perform your tests

		rowCount++
	}
}
```

### Metrics

The `metrics` package contains a set of utility functions that can be used to create and initialise a Prometheus client.

There are two main components to the metrics package, `Server` and `Instrumentation`. The `Server` is used to create a HTTP
server that will serve the Prometheus metrics endpoint. The `Instrumentation` is used to create the instrumentations that are
required to capture the metrics.

To use the metrics package, you need to create your own instrumentations package which will define the instrumentation types and
build the instrumentations you need for your application.

```go
package instrumentation

import (
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/gobl/gobl/pkg/config"
	"gitlab.com/gobl/gobl/pkg/metrics"
)

// TODO: Add instrumentation type definitions below as needed.
const (
	_ metrics.InstrumentationType = iota
	TypeInfo
	TypeRequests
)

// SetupInstrumentation initializes the instrumentations required for
// the metrics service.
func SetupInstrumentation(namespace string) *metrics.Instrumentation {
	i := createInstrumentation(namespace)

	initInstrumentation(i)
	return i
}

// createInstrumentation creates the instrumentations for the metrics
// that are to be published.
func createInstrumentation(ns string) *metrics.Instrumentation {
	return metrics.NewInstrumentation(ns).
		WithGaugeVec(
			TypeInfo,
			"info",
			"Information about the application",
			"version",
		)
}

// initInstrumentation initializes instrumentation with any values that
// should be set on startup.
func initInstrumentation(i *metrics.Instrumentation) {
	setVersion(i)
}

// setVersion sets the version of the application.
func setVersion(i *metrics.Instrumentation) {
	version := config.Get("version").String("unknown")
	i.GaugeVecs[TypeInfo].With(prometheus.Labels{"version": version}).Set(1)
}

// TODO: Add any additional instrumentation functions as needed.
```

Then you can initialise the metrics server in your application startup code and register the instrumentations.

```go
cfg, err := metrics.ReadConfigFromFile("application.metrics")
if err != nil {
	return fmt.Errorf("failed to read metrics config: %w", err)
}

i := instrumentation.SetupInstrumentation("signal_generator")
metricsSvr := metrics.DefaultServer(cfg)
if err = metricsSvr.Register(i); err != nil {
	return fmt.Errorf("failed to register instrumentations: %w", err)
}
metricsSvr.Start()
```

When stopping your application, make sure you gracefully stop the metrics server.

```go
metricsSvr.Stop()
```

#### Metrics Middleware

##### Echo

You can create a middleware that will capture metrics for your Echo HTTP handlers.

```go
package metrics

import (
	"time"
	"your/module/instrumentation"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/gobl/gobl/pkg/metrics"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	Instrumentation *metrics.Instrumentation
}

func (m *Middleware) EchoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		m.Instrumentation.CounterVecs[instrumentation.TypeRequests].With(prometheus.Labels{
			"method": c.Request().Method,
			"path":   c.Path(),
		}).Inc()

		start := time.Now()
		err := next(c)
		elapsed := time.Since(start).Seconds()
		url := c.Request().URL.Path

		m.Instrumentation.HistogramVecs[instrumentation.TypeRequests].With(prometheus.Labels{
			"method": c.Request().Method,
			"path":   url,
		}).Observe(elapsed)

		return err
	}
}
```

Then you can use the middleware in your Echo application.

```go
e := echo.New()
e.Use(metrics.EchoMiddleware)
```

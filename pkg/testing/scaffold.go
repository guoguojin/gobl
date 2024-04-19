package testing

type Scaffold struct {
	Port     string
	Teardown TeardownFunc
}

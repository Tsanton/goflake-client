package goflake_test

import (
	"testing"

	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
)

func Test_execute_scalar_int(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()

	/* Act */
	res, err := g.ExecuteScalar[int](cli, "select 93")
	if err != nil {
		t.Errorf("unable to run ExecuteScalar")
		t.FailNow()
	}

	if res != 93 {
		t.Fail()
	}
}

func Test_execute_scalar_true(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()

	/* Act */
	res, err := g.ExecuteScalar[bool](cli, "select true")
	if err != nil {
		t.Errorf("unable to run ExecuteScalar")
		t.FailNow()
	}

	if res != true {
		t.Fail()
	}
}

func Test_execute_scalar_false(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()

	/* Act */
	res, err := g.ExecuteScalar[bool](cli, "select false")
	if err != nil {
		t.Errorf("unable to run ExecuteScalar")
		t.FailNow()
	}

	if res != false {
		t.Fail()
	}
}

func Test_execute_scalar_string(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()

	/* Act */
	res, err := g.ExecuteScalar[string](cli, "select 'Hello world!'")
	if err != nil {
		t.Errorf("unable to run ExecuteScalar")
		t.FailNow()
	}

	if res != "Hello world!" {
		t.Fail()
	}
}

func Test_execute_scalar_float(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()

	/* Act */
	res, err := g.ExecuteScalar[float32](cli, "select 3.14")
	if err != nil {
		t.Errorf("unable to run ExecuteScalar")
		t.FailNow()
	}

	if res != 3.14 {
		t.Fail()
	}
}

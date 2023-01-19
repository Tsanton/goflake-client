package models_test

import (
	"testing"

	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_create_role(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[a.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	role := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   "USERADMIN",
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))

	// exists, err := g.ExecuteScalar[bool](cli, fmt.Sprintf("SHOW DATABASE "))

	// /* Assert */
	// if !exists || err != nil {
	// 	t.Fail()
	// }
}

func Test_describe_role(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[a.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	r := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   "USERADMIN",
	}

	i.ErrorFailNow(t, g.RegisterAsset(cli, &r, &stack))

	/* Act */
	dr, err := g.Describe[e.Role](cli, &d.Role{Name: r.Name})
	i.ErrorFailNow(t, err)

	/* Assert */
	if dr.Name != r.Name {
		t.FailNow()
	}
}

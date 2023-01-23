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

func Test_create_database(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[a.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DEMO",
		Comment: "integration test goflake",
		Owner:   "SYSADMIN",
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
}

func Test_describe_database(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[a.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DEMO",
		Comment: "integration test goflake",
		Owner:   "SYSADMIN",
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))

	/* Act */
	ddb, err := g.Describe[e.Database](cli, &d.Database{Name: db.Name})
	i.ErrorFailNow(t, err)

	/* Assert */
	if ddb.Name != db.Name {
		t.FailNow()
	}
}

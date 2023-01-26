package models_test

import (
	"testing"

	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_role_hierarchy(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	/* Arrange */
	rr := a.Role{
		Name:    "SOME_SCHEMA_R",
		Comment: "integration test goflake",
		Owner:   "USERADMIN",
	}
	rrw := a.Role{
		Name:    "SOME_SCHEMA_RW",
		Comment: "integration test goflake",
		Owner:   "USERADMIN",
	}
	rrwc := a.Role{
		Name:    "SOME_SCHEMA_RWC",
		Comment: "integration test goflake",
		Owner:   "USERADMIN",
	}
	rel1 := a.RoleRelationship{
		ChildRoleName:  rr.Name,
		ParentRoleName: rrw.Name,
	}
	rel2 := a.RoleRelationship{
		ChildRoleName:  rrw.Name,
		ParentRoleName: rrwc.Name,
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rr, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rrw, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rrwc, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel1, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel2, &stack))

	/* Act */
	hier, err := g.Describe[e.RoleHierarchy](cli, &d.RoleHierarchy{RoleName: rr.Name})

	/* Assert */
	if err != nil || len(hier.InheritingRoles) != 2 {
		t.FailNow()
	}
}

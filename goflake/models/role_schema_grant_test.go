package models_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ag "github.com/tsanton/goflake-client/goflake/models/assets/grants"
	dg "github.com/tsanton/goflake-client/goflake/models/describables/grants"
	eg "github.com/tsanton/goflake-client/goflake/models/entities/grants"
	"github.com/tsanton/goflake-client/goflake/models/enums"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_grant_role_schema_privilege(t *testing.T) {
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
	schema := a.Schema{
		Database: db,
		Name:     "IGT_GRANT",
		Comment:  "integration test goflake",
		Owner:    "SYSADMIN",
	}
	role := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   "USERADMIN",
	}
	privilege := a.Grant{
		Target:     &ag.RoleSchemaGrant{RoleName: role.Name, DatabaseName: db.Name, SchemaName: schema.Name},
		Privileges: []enums.Privilege{enums.PrivilegeUsage},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.RoleGrant{RoleName: role.Name})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 1)
	schemaUsage, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeUsage })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", schemaUsage.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectSchema, schemaUsage.GrantedOn)
}

func Test_grant_role_schema_privileges(t *testing.T) {
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
	schema := a.Schema{
		Database: db,
		Name:     "IGT_GRANT",
		Comment:  "integration test goflake",
		Owner:    "SYSADMIN",
	}
	role := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   "USERADMIN",
	}
	privilege := a.Grant{
		Target:     &ag.RoleSchemaGrant{RoleName: role.Name, DatabaseName: db.Name, SchemaName: schema.Name},
		Privileges: []enums.Privilege{enums.PrivilegeUsage, enums.PrivilegeMonitor},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.RoleGrant{RoleName: role.Name})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 2)

	schemaUsage, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeUsage })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", schemaUsage.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectSchema, schemaUsage.GrantedOn)

	schemaMonitor, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeMonitor })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", schemaMonitor.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectSchema, schemaMonitor.GrantedOn)
}

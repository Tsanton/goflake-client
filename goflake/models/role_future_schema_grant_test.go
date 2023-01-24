package models_test

import (
	"fmt"
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

func Test_grant_role_future_schema_privilege(t *testing.T) {
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
		Target:     &ag.RoleFutureSchemaGrant{RoleName: role.Name, DatabaseName: db.Name, SchemaName: schema.Name, ObjectType: enums.SnowflakeObjectTable},
		Privileges: []enums.Privilege{enums.PrivilegeSelect},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleFutureGrants](cli, &dg.RoleFutureGrant{RoleName: role.Name})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 1)
	schemaFutureSelect, ok := lo.Find(res.Grants, func(i eg.RoleFutureGrant) bool {
		return i.Privilege == enums.Privilege(enums.PrivilegeSelect.String())
	})
	assert.True(t, ok)
	assert.Equal(t, schemaFutureSelect.GrantedOn, enums.SnowflakeObjectTable)
	assert.Equal(t, fmt.Sprintf("%[1]s.%[2]s.<%[3]s>", db.Name, schema.Name, enums.SnowflakeObjectTable.ToSingular()), schemaFutureSelect.GrantTargetName)
}

func Test_grant_role_future_schema_privileges(t *testing.T) {
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
	privilege1 := a.Grant{
		Target:     &ag.RoleFutureSchemaGrant{RoleName: role.Name, DatabaseName: db.Name, SchemaName: schema.Name, ObjectType: enums.SnowflakeObjectTable},
		Privileges: []enums.Privilege{enums.PrivilegeSelect, enums.PrivilegeUpdate},
	}
	privilege2 := a.Grant{
		Target:     &ag.RoleFutureSchemaGrant{RoleName: role.Name, DatabaseName: db.Name, SchemaName: schema.Name, ObjectType: enums.SnowflakeObjectView},
		Privileges: []enums.Privilege{enums.PrivilegeSelect, enums.PrivilegeReferences},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege1, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege2, &stack))

	res, err := g.Describe[*eg.RoleFutureGrants](cli, &dg.RoleFutureGrant{RoleName: role.Name})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 4)

	tableSchemaScope := fmt.Sprintf("%[1]s.%[2]s.<%[3]s>", db.Name, schema.Name, enums.SnowflakeObjectTable.ToSingular())
	_, ok := lo.Find(res.Grants, func(i eg.RoleFutureGrant) bool {
		return i.Privilege == enums.PrivilegeSelect && i.GrantTargetName == tableSchemaScope
	})
	assert.True(t, ok)

	_, ok = lo.Find(res.Grants, func(i eg.RoleFutureGrant) bool {
		return i.Privilege == enums.PrivilegeUpdate && i.GrantTargetName == tableSchemaScope
	})
	assert.True(t, ok)

	viewSchemaScope := fmt.Sprintf("%[1]s.%[2]s.<%[3]s>", db.Name, schema.Name, enums.SnowflakeObjectView.ToSingular())
	_, ok = lo.Find(res.Grants, func(i eg.RoleFutureGrant) bool {
		return i.Privilege == enums.PrivilegeSelect && i.GrantTargetName == viewSchemaScope
	})
	assert.True(t, ok)

	_, ok = lo.Find(res.Grants, func(i eg.RoleFutureGrant) bool {
		return i.Privilege == enums.PrivilegeReferences && i.GrantTargetName == viewSchemaScope
	})
	assert.True(t, ok)
}

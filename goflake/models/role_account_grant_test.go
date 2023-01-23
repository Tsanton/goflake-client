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

func Test_grant_role_account_privilege(t *testing.T) {
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
	privilege := a.Grant{
		Target:     &ag.RoleAccountGrant{RoleName: role.Name},
		Privileges: []enums.Privilege{enums.PrivilegeCreateAccount},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.RoleGrant{RoleName: "IGT_DEMO_ROLE"})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 1)
	createAcc, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool {
		return i.Privilege == enums.Privilege(enums.PrivilegeCreateAccount.String())
	})
	assert.True(t, ok)
	assert.Equal(t, "ACCOUNTADMIN", createAcc.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectAccount, createAcc.GrantedOn)
}

func Test_grant_role_account_privileges(t *testing.T) {
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
	privilege := a.Grant{
		Target:     &ag.RoleAccountGrant{RoleName: role.Name},
		Privileges: []enums.Privilege{enums.PrivilegeCreateAccount, enums.PrivilegeCreateUser},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.RoleGrant{RoleName: "IGT_DEMO_ROLE"})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 2)

	createAcc, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeCreateAccount })
	assert.True(t, ok)
	assert.Equal(t, "ACCOUNTADMIN", createAcc.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectAccount, createAcc.GrantedOn)

	createUser, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeCreateUser })
	assert.True(t, ok)
	assert.Equal(t, "USERADMIN", createUser.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectAccount, createUser.GrantedOn)
}

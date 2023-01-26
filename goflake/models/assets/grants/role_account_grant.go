package grants

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeGrant = &RoleAccountGrant[i.ISnowflakeRole]{}
)

type RoleAccountGrant[T i.ISnowflakeRole] struct {
	Role T
}

func (r *RoleAccountGrant[T]) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	if r.Role.IsDatabaseRole() {
		panic("you can't grant account level privileges to database roles")
	}
	return fmt.Sprintf("GRANT %[1]s ON ACCOUNT TO ROLE %[2]s;", privs, r.Role.GetIdentifier()), 1
}

func (r *RoleAccountGrant[T]) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	if r.Role.IsDatabaseRole() {
		panic("you can't neither grant nor revoke account level privileges to/from database roles")
	}
	return fmt.Sprintf("REVOKE %[1]s ON ACCOUNT FROM ROLE %[2]s CASCADE;", privs, r.Role.GetIdentifier()), 1
}

func (*RoleAccountGrant[T]) validatePrivileges(privileges []enum.Privilege) bool {
	allowedPrivileges := []enum.Privilege{
		enum.PrivilegeCreateAccount,
		enum.PrivilegeCreateDataExchangeListing,
		enum.PrivilegeCreateDatabase,
		enum.PrivilegeCreateIntegration,
		enum.PrivilegeCreateNetworkPolicy,
		enum.PrivilegeCreateRole,
		enum.PrivilegeCreateShare,
		enum.PrivilegeCreateUser,
		enum.PrivilegeCreateWarehouse,

		enum.PrivilegeApplyMaskingPolicy,
		// enum.PrivilegeApplyPasswordPolicy, //Missing enum
		enum.PrivilegeApplyRowAccessPolicy,
		// enum.PrivilegeApplySessionPolicy, //Missing enum
		enum.PrivilegeApplyTag,
		enum.PrivilegeAttachPolicy,
		enum.PrivilegeExecuteTask,
		enum.PrivilegeImportShare,
		enum.PrivilegeManageGrants,
		enum.PrivilegeMonitorExecution,
		enum.PrivilegeMonitorUsage,
		enum.PrivilegeOverrideShareRestrictions,
		enum.PrivilegeExecuteManagedTask,
		enum.PrivilegeOrganizationSupportCases,
		enum.PrivilegeAccountSupportCases,
		enum.PrivilegeUserSupportCases,
	}
	return lo.Every(allowedPrivileges, privileges)
}

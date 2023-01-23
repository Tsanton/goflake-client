package grants

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeGrant = &RoleAccountGrant{}
)

type RoleAccountGrant struct {
	RoleName string
}

func (r *RoleAccountGrant) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	return fmt.Sprintf("GRANT %[1]s ON ACCOUNT TO ROLE %[2]s;", privs, r.RoleName), 1
}

func (r *RoleAccountGrant) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	return fmt.Sprintf("REVOKE %[1]s ON ACCOUNT FROM ROLE %[2]s CASCADE;", privs, r.RoleName), 1
}

func (*RoleAccountGrant) validatePrivileges(privileges []enum.Privilege) bool {
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

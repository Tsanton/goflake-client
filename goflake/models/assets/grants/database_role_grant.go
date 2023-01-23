package grants

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeGrant = &AccountRoleGrant{}
)

type DatabaseRoleGrant struct {
	RoleName     string
	DatabaseName string
}

func (r *DatabaseRoleGrant) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	priv := strings.Join(stringPrivileges, ", ")
	return fmt.Sprintf("GRANT %[1]s ON DATABASE %[2]s TO ROLE %[3]s;", priv, r.DatabaseName, r.RoleName), 1
}

func (r *DatabaseRoleGrant) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	priv := strings.Join(stringPrivileges, ", ")
	return fmt.Sprintf("REVOKE %[1]s ON DATABASE %[2]s FROM ROLE %[3]s CASCADE;", priv, r.DatabaseName, r.RoleName), 1
}

func (*DatabaseRoleGrant) validatePrivileges(privileges []enum.Privilege) bool {
	allowedPrivileges := []enum.Privilege{
		// enum.PrivilegeCreateDatabaseRole, //Missing enum
		enum.PrivilegeCreateSchema,
		enum.PrivilegeImportedPrivileges,
		enum.PrivilegeModify,
		enum.PrivilegeMonitor,
		enum.PrivilegeUsage,
	}
	return lo.Every(allowedPrivileges, privileges)
}

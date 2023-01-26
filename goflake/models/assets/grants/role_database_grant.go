package grants

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeGrant = &RoleDatabaseGrant[i.ISnowflakeRole]{}
)

type RoleDatabaseGrant[T i.ISnowflakeRole] struct {
	Role         i.ISnowflakeRole
	DatabaseName string
}

func (r *RoleDatabaseGrant[T]) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	if r.Role.IsDatabaseRole() {
		return fmt.Sprintf("GRANT %[1]s ON DATABASE %[2]s TO DATABASE ROLE %[3]s;", privs, r.DatabaseName, r.Role.GetIdentifier()), 1
	}
	return fmt.Sprintf("GRANT %[1]s ON DATABASE %[2]s TO ROLE %[3]s;", privs, r.DatabaseName, r.Role.GetIdentifier()), 1
}

func (r *RoleDatabaseGrant[T]) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	if r.Role.IsDatabaseRole() {
		return fmt.Sprintf("REVOKE %[1]s ON DATABASE %[2]s FROM DATABASE ROLE %[3]s CASCADE;", privs, r.DatabaseName, r.Role.GetIdentifier()), 1
	}
	return fmt.Sprintf("REVOKE %[1]s ON DATABASE %[2]s FROM ROLE %[3]s CASCADE;", privs, r.DatabaseName, r.Role.GetIdentifier()), 1
}

func (*RoleDatabaseGrant[T]) validatePrivileges(privileges []enum.Privilege) bool {
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

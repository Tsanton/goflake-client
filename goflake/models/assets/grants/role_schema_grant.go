package grants

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeGrant = &RoleSchemaGrant[i.ISnowflakeRole]{}
)

type RoleSchemaGrant[T i.ISnowflakeRole] struct {
	Role         T
	DatabaseName string
	SchemaName   string
}

func (r *RoleSchemaGrant[T]) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	priv := strings.Join(stringPrivileges, ", ")
	if r.Role.IsDatabaseRole() {
		panic("Database role not implementer")
	}
	return fmt.Sprintf("GRANT %[1]s ON SCHEMA %[2]s.%[3]s TO ROLE %[4]s;", priv, r.DatabaseName, r.SchemaName, r.Role.GetIdentifier()), 1
}

func (r *RoleSchemaGrant[T]) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	priv := strings.Join(stringPrivileges, ", ")
	if r.Role.IsDatabaseRole() {
		panic("Database role not implementer")
	}
	return fmt.Sprintf("REVOKE %[1]s ON SCHEMA %[2]s.%[3]s FROM ROLE %[4]s CASCADE;", priv, r.DatabaseName, r.SchemaName, r.Role.GetIdentifier()), 1
}

func (*RoleSchemaGrant[T]) validatePrivileges(privileges []enum.Privilege) bool {
	allowedAccountPrivileges := []enum.Privilege{
		enum.PrivilegeModify,
		enum.PrivilegeMonitor,
		enum.PrivilegeUsage,
		enum.PrivilegeAddSearchOptimization,
		//Create
		enum.PrivilegeCreateExternalTable,
		enum.PrivilegeCreateFileFormat,
		enum.PrivilegeCreateFunction,
		enum.PrivilegeCreateMaskingPolicy,
		enum.PrivilegeCreateMaterializedView,
		// enum.PrivilegePasswordPolicy, //Missing Enum
		enum.PrivilegeCreatePipe,
		enum.PrivilegeCreateProcedure,
		enum.PrivilegeCreateRowAccessPolicy,
		// enum.CreateSessionPolicy, //Missing Enum
		enum.PrivilegeCreateSequence,
		enum.PrivilegeCreateStage,
		enum.PrivilegeCreateStream,
		enum.PrivilegeCreateTag,
		enum.PrivilegeCreateTable,
		enum.PrivilegeCreateTask,
		enum.PrivilegeCreateView,
	}
	return lo.Every(allowedAccountPrivileges, privileges)
}

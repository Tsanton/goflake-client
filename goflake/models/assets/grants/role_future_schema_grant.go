package grants

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeGrant = &RoleFutureSchemaGrant{}
)

type RoleFutureSchemaGrant struct {
	RoleName     string
	DatabaseName string
	SchemaName   string
	ObjectType   enum.SnowflakeObject
}

func (r *RoleFutureSchemaGrant) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	priv := strings.Join(stringPrivileges, ", ")
	return fmt.Sprintf("GRANT %[1]s ON FUTURE %[2]s IN SCHEMA %[3]s.%[4]s TO ROLE %[5]s;", priv, r.ObjectType.ToPlural(), r.DatabaseName, r.SchemaName, r.RoleName), 1
}

func (r *RoleFutureSchemaGrant) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	priv := strings.Join(stringPrivileges, ", ")
	return fmt.Sprintf("REVOKE %[1]s ON FUTURE %[2]s IN SCHEMA %[3]s.%[4]s FROM ROLE %[5]s CASCADE;", priv, r.ObjectType.ToPlural(), r.DatabaseName, r.SchemaName, r.RoleName), 1
}

func (r *RoleFutureSchemaGrant) validatePrivileges(privileges []enum.Privilege) bool {
	var allowedAccountPrivileges []enum.Privilege

	switch r.ObjectType {
	case enum.SnowflakeObjectTable:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeSelect,
			enum.PrivilegeInsert,
			enum.PrivilegeUpdate,
			enum.PrivilegeDelete,
			enum.PrivilegeTruncate,
			enum.PrivilegeReferences,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectView, enum.SnowflakeObjectMatView:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeSelect,
			enum.PrivilegeReferences,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectSequence, enum.SnowflakeObjectFunction, enum.SnowflakeObjectProcedure, enum.SnowflakeObjectFileFormat:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeUsage,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectInternalStage:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeRead,
			enum.PrivilegeWrite,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectExternalStage:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeUsage,
			enum.PrivilegeOwnership, //?
		}
	case enum.SnowflakeObjectPipe:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeMonitor,
			enum.PrivilegeOperate,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectStream:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeSelect,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectTask:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeMonitor,
			enum.PrivilegeOperate,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectMaskingPolicy, enum.SnowflakeObjectPasswordPolicy, enum.SnowflakeObjectRowAccessPolicy, enum.SnowflakeObjectTag:
		allowedAccountPrivileges = []enum.Privilege{
			enum.PrivilegeApply,
			enum.PrivilegeOwnership,
		}
	}
	return lo.Every(allowedAccountPrivileges, privileges)
}

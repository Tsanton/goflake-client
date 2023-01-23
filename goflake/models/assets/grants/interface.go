package grants

import (
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

type ISnowflakeGrant interface {
	GetGrantStatement(privileges []enum.Privilege) (string, int)
	GetRevokeStatement(privileges []enum.Privilege) (string, int)
	validatePrivileges(privileges []enum.Privilege) bool
}

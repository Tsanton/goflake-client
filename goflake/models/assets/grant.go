package assets

import (
	gra "github.com/tsanton/goflake-client/goflake/models/assets/grants"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeAsset = &Grant{}
)

type Grant struct {
	Target     gra.ISnowflakeGrant
	Privileges []enum.Privilege
}

func (r *Grant) GetCreateStatement() (string, int) {
	return r.Target.GetGrantStatement(r.Privileges)
}

func (r *Grant) GetDeleteStatement() (string, int) {
	return r.Target.GetRevokeStatement(r.Privileges)
}

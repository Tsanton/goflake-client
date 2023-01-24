package grants

import (
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

type RoleFutureGrant struct {
	Privilege       enum.Privilege       `json:"privilege"`
	GrantedOn       enum.SnowflakeObject `json:"grant_on"`
	GrantTargetName string               `json:"name"`
	GrantOption     string               `json:"grant_option"` //TODO: Bool and custom converter due to "true" & "false"
}

var (
	_ e.ISnowflakeEntity = &RoleFutureGrants{}
)

type RoleFutureGrants struct {
	RoleName string            `json:"role_name"`
	Grants   []RoleFutureGrant `json:"grants"`
}

func (r *RoleFutureGrants) GetIdentity() string {
	return "implements ISnowflakeEntity interface"
}

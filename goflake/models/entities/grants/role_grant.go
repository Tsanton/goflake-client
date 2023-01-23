package grants

import (
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

type RoleGrant struct {
	Privilege       enum.Privilege       `json:"privilege"`
	GrantedOn       enum.SnowflakeObject `json:"granted_on"`
	GrantTargetName string               `json:"name"`
	GrantOption     string               `json:"grant_option"` //TODO: Bool and custom converter due to "true" & "false"
	GrantedBy       string               `json:"granted_by"`
}

var (
	_ e.ISnowflakeEntity = &FutureGrant{}
)

type RoleGrants struct {
	RoleName string      `json:"role_name"`
	Grants   []RoleGrant `json:"grants"`
}

func (r *RoleGrants) GetIdentity() string {
	return "implements ISnowflakeEntity interface"
}

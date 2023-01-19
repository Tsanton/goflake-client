package entities

import (
	"fmt"
	"strings"

	u "github.com/tsanton/goflake-client/goflake/utilities"
)

var (
	_ ISnowflakeEntity = RoleRelationship{}
)

type RoleRelationship struct {
	ChildRoleName  string     `json:"child_role_name"`
	ParentRoleName string     `json:"parent_role_name"`
	GrantOption    string     `json:"grant_option"` //TODO: Bool and custom converter due to "true" & "false"
	GrantedBy      string     `json:"granted_by"`
	GrantedOn      u.SnowTime `json:"created_on"`
}

func (r RoleRelationship) GetIdentity() string {
	return strings.ToUpper(fmt.Sprintf("%[1]s.%[2]s", r.ChildRoleName, r.ParentRoleName))
}

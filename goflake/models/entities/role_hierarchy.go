package entities

type InheritedRole struct {
	RoleName           string `json:"role_name"`
	ParentRoleName     string `json:"parent_role_name"`
	DistanceFromSource int    `json:"distance_from_source"`
	GrantOption        string `json:"grant_option"` //TODO: Bool and custom converter due to "true" & "false"
	GrantedBy          string `json:"granted_by"`
	CreatedOn          string `json:"created_on"`
}

var (
	_ ISnowflakeEntity = RoleHierarchy{}
)

type RoleHierarchy struct {
	Name            string          `json:"name"`
	InheritingRoles []InheritedRole `json:"inheriting_roles"`
}

// GetIdentity implements ISnowflakeEntity
func (r RoleHierarchy) GetIdentity() string {
	return r.Name
}

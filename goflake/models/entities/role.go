package entities

import "time"

var (
	_ ISnowflakeEntity = Role{}
)

type Role struct {
	Name            string    `db:"name"`
	Owner           string    `db:"owner"`
	AssignedToUsers int       `db:"assigned_to_users"`
	GrantedToRoles  int       `db:"granted_to_roles"`
	GrantedRoles    int       `db:"granted_roles"`
	Comment         string    `db:"comment"`
	CreatedOn       time.Time `db:"created_on"`
}

func (r Role) GetIdentity() string {
	return r.Name
}

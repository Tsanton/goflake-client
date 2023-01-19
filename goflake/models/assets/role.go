package assets

import "fmt"

var (
	_ ISnowflakeAsset = &Role{}
)

type Role struct {
	Name    string
	Owner   string
	Comment string
}

func (r *Role) GetCreateStatement() (string, int) {
	return fmt.Sprintf(`
	CREATE OR REPLACE ROLE %[1]s COMMENT = '%[2]s';
	GRANT OWNERSHIP ON ROLE %[1]s TO %[3]s REVOKE CURRENT GRANTS;`,
		r.Name, r.Comment, r.Owner,
	), 2
}

func (r *Role) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP ROLE IF EXISTS %[1]s;", r.Name), 1
}

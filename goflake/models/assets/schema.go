package assets

import "fmt"

var (
	_ ISnowflakeAsset = &Schema{}
)

type Schema struct {
	Database Database
	Name     string
	Comment  string
	Owner    string
}

func (r *Schema) GetCreateStatement() (string, int) {
	return fmt.Sprintf(`
CREATE OR REPLACE SCHEMA %[1]s.%[2]s WITH MANAGED ACCESS COMMENT = '%[3]s';
GRANT OWNERSHIP ON SCHEMA %[1]s.%[2]s TO %[4]s REVOKE CURRENT GRANTS;
`,
		r.Database.Name, r.Name, r.Comment, r.Owner,
	), 2
}

func (r *Schema) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %[1]s.%[2]s CASCADE;", r.Database.Name, r.Name), 1
}

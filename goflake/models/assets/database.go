package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
)

var (
	_ i.ISnowflakeAsset = &Database{}
)

type Database struct {
	Name    string
	Comment string
	Owner   string
}

func (r *Database) GetCreateStatement() (string, int) {
	return fmt.Sprintf(`
CREATE OR REPLACE DATABASE %[1]s COMMENT = '%[2]s';
GRANT OWNERSHIP ON DATABASE %[1]s TO %[3]s;
`,
		r.Name, r.Comment, r.Owner,
	), 2
}

func (r *Database) GetDeleteStatement() (string, int) {
	return fmt.Sprintf(`DROP DATABASE IF EXISTS %[1]s CASCADE;`, r.Name), 1
}

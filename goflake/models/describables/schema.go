package describables

import "fmt"

var (
	_ ISnowflakeDescribable = &Schema{}
)

type Schema struct {
	DatabaseName string
	SchemaName   string
}

func (r *Schema) GetDescribeStatement() string {
	return fmt.Sprintf("SHOW SCHEMAS LIKE '%[2]s' IN DATABASE %[1]s;", r.DatabaseName, r.SchemaName)
}

func (r *Schema) IsProcedure() bool {
	return false
}

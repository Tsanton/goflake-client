package describables

import (
	"fmt"
	"strings"
)

var (
	_ ISnowflakeDescribable = &Database{}
)

type Database struct {
	Name string
}

func (r *Database) GetDescribeStatement() string {
	return strings.ToUpper(fmt.Sprintf("SHOW DATABASES LIKE '%[1]s'", r.Name))
}

func (r *Database) IsProcedure() bool {
	return false
}

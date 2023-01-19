package entities

import (
	"fmt"
	"time"
)

var (
	_ ISnowflakeEntity = Schema{}
)

type Schema struct {
	Name          string    `db:"name"`
	DatabaseName  string    `db:"database_name"`
	Owner         string    `db:"owner"`
	Comment       string    `db:"comment"`
	RetentionTime int       `db:"retention_time"`
	CreatedOn     time.Time `db:"created_on"`
}

func (r Schema) GetIdentity() string {
	return fmt.Sprintf("%[1]s.%[2]s", r.DatabaseName, r.Name)
}

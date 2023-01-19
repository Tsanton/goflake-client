package entities

import "time"

var (
	_ ISnowflakeEntity = Database{}
)

type Database struct {
	Name          string    `db:"name"`
	Owner         string    `db:"owner"`
	Origin        string    `db:"origin"`
	RetentionTime int       `db:"retention_time"`
	CreatedOn     time.Time `db:"created_on"`
}

func (r Database) GetIdentity() string {
	return r.Name
}

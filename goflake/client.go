package goflake

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/snowflakedb/gosnowflake"
)

type GoflakeClient struct {
	SnowflakeHost string
	SnowflakeUid  string
	SnowflakePwd  string
	SnowflakeRole string
	SnowflakeWh   string
	db            *sqlx.DB

	dbName       string
	configSchema string
}

func (g *GoflakeClient) Open() {
	var err error
	g.db, err = sqlx.Connect("snowflake", fmt.Sprintf("%s:%s@%s:443?warehouse=%s&role=%s", g.SnowflakeUid, g.SnowflakePwd, g.SnowflakeHost, g.SnowflakeWh, g.SnowflakeRole))
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}
	if err := g.db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	g.db = g.db.Unsafe()
}

func (g *GoflakeClient) Configure(dbName string, configSchema string) {
	g.dbName = dbName
	g.configSchema = configSchema
}

func (c *GoflakeClient) Close() {
	c.db.Close()
}

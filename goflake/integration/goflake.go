package integration

import (
	"os"

	g "github.com/tsanton/goflake-client/goflake"
)

func Goflake() *g.GoflakeClient {
	cli := g.GoflakeClient{
		SnowflakeHost: os.Getenv("SNOWFLAKE_HOST"),
		SnowflakeUid:  os.Getenv("SNOWFLAKE_UID"),
		SnowflakePwd:  os.Getenv("SNOWFLAKE_PWD"),
		SnowflakeRole: os.Getenv("SNOWFLAKE_ROLE"),
		SnowflakeWh:   os.Getenv("SNOWFLAKE_WH"),
	}

	cli.Open()

	return &cli
}

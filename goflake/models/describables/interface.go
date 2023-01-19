package describables

type ISnowflakeDescribable interface {
	GetDescribeStatement() string
	IsProcedure() bool
}

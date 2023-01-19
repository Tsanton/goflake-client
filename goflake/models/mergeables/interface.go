package mergeables

type ISnowflakeMergeable interface {
	MergeIntoStatement() string
	SelectStatement() string
}

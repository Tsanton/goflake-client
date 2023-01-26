package inter

type ISnowflakeRole interface {
	GetIdentifier() string
	IsDatabaseRole() bool
}

package assets

type ISnowflakeAsset interface {
	GetCreateStatement() (string, int)
	GetDeleteStatement() (string, int)
}

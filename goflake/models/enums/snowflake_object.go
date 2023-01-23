package enums

type SnowflakeObject string

func (p SnowflakeObject) String() string {
	return string(p)
}

func (p SnowflakeObject) ToSingular() string {
	return string(p)
}
func (p SnowflakeObject) ToPlural() string {
	switch p {
	case SnowflakeObjectTable:
		return "TABLES"
	case SnowflakeObjectView:
		return "VIEWS"
	case SnowflakeObjectAccount:
		return "ACCOUNTS"
	case SnowflakeObjectDatabase:
		return "DATABASES"
	case SnowflakeObjectDatabaseRole:
		return "DATABASE_ROLES"
	case SnowflakeObjectFunction:
		return "FUNCTIONS"
	case SnowflakeObjectRole:
		return "ROLES"
	case SnowflakeObjectSchema:
		return "SCHEMAS"
	case SnowflakeObjectTag:
		return "TAGS"
	case SnowflakeObjectUser:
		return "USERS"
	default:
		panic("not implemented")
	}
}

const (
	SnowflakeObjectTable        SnowflakeObject = "TABLE"
	SnowflakeObjectView         SnowflakeObject = "VIEW"
	SnowflakeObjectAccount      SnowflakeObject = "ACCOUNT"
	SnowflakeObjectDatabase     SnowflakeObject = "DATABASE"
	SnowflakeObjectDatabaseRole SnowflakeObject = "DATABASE_ROLE"
	SnowflakeObjectFunction     SnowflakeObject = "FUNCTION"
	SnowflakeObjectRole         SnowflakeObject = "ROLE"
	SnowflakeObjectSchema       SnowflakeObject = "SCHEMA"
	SnowflakeObjectTag          SnowflakeObject = "TAG"
	SnowflakeObjectUser         SnowflakeObject = "USER"
)

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
	case SnowflakeObjectMatView:
		return "MATERIALIZED VIEWS"
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
	case SnowflakeObjectSequence:
		return "SEQUENCES"
	case SnowflakeObjectProcedure:
		return "PROCEDURES"
	case SnowflakeObjectFileFormat:
		return "FILE FORMATS"
	case SnowflakeObjectInternalStage:
		return "INTERNAL STAGES"
	case SnowflakeObjectExternalStage:
		return "EXTERNAL STAGES"
	case SnowflakeObjectPipe:
		return "PIPES"
	case SnowflakeObjectStream:
		return "STREAMS"
	case SnowflakeObjectTask:
		return "TASKS"
	case SnowflakeObjectMaskingPolicy:
		return "MASKING POLICIES"
	case SnowflakeObjectPasswordPolicy:
		return "PASSWORD POLICIES"
	case SnowflakeObjectRowAccessPolicy:
		return "ROW ACCESS POLICIES"
	case SnowflakeObjectWarehouse:
		return "WAREHOUSES"
	default:
		panic("not implemented")
	}
}

const (
	SnowflakeObjectTable           SnowflakeObject = "TABLE"
	SnowflakeObjectView            SnowflakeObject = "VIEW"
	SnowflakeObjectMatView         SnowflakeObject = "MATERIALIZED VIEW"
	SnowflakeObjectAccount         SnowflakeObject = "ACCOUNT"
	SnowflakeObjectDatabase        SnowflakeObject = "DATABASE"
	SnowflakeObjectDatabaseRole    SnowflakeObject = "DATABASE_ROLE"
	SnowflakeObjectFunction        SnowflakeObject = "FUNCTION"
	SnowflakeObjectRole            SnowflakeObject = "ROLE"
	SnowflakeObjectSchema          SnowflakeObject = "SCHEMA"
	SnowflakeObjectTag             SnowflakeObject = "TAG"
	SnowflakeObjectUser            SnowflakeObject = "USER"
	SnowflakeObjectSequence        SnowflakeObject = "SEQUENCE"
	SnowflakeObjectProcedure       SnowflakeObject = "PROCEDURE"
	SnowflakeObjectFileFormat      SnowflakeObject = "FILE FORMAT"
	SnowflakeObjectInternalStage   SnowflakeObject = "INTERNAL STAGE"
	SnowflakeObjectExternalStage   SnowflakeObject = "EXTERNAL STAGE"
	SnowflakeObjectPipe            SnowflakeObject = "PIPE"
	SnowflakeObjectStream          SnowflakeObject = "STREAM"
	SnowflakeObjectTask            SnowflakeObject = "TASK"
	SnowflakeObjectMaskingPolicy   SnowflakeObject = "MASKING POLICY"
	SnowflakeObjectPasswordPolicy  SnowflakeObject = "PASSWORD POLICY"
	SnowflakeObjectRowAccessPolicy SnowflakeObject = "ROW ACCESS POLICY"
	SnowflakeObjectWarehouse       SnowflakeObject = "WAREHOUSE"
)

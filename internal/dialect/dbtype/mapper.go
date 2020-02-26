package dbtype

var matchers = []TypeMatcher{
	NewMappingType(
		"bool",
		[]string{
			"tinyint", "boolean", "bool",
		},
	),
	NewMappingType(
		"float32",
		[]string{
			"decimal", "float", "double",
			"numeric", "real", "double precision",
		},
	),
	NewMappingType(
		"time.Time",
		[]string{
			"date", "time", "datetime", "timestamp",
			"year", "interval", "timestamptz",
		},
	),
	NewMappingType(
		"int32",
		[]string{
			"smallint", "mediumint", "int",
			"bigint", "bit", "integer",
			"smallserial", "serial", "bigserial",
		},
	),
	NewMappingType(
		"string",
		[]string{
			"char", "character", "varchar", "character varying",
			"bytea", "binary", "varbinary", "tinyblob",
			"blob", "mediumblob", "longblob", "tinytext",
			"text", "mediumtext", "longtext",
		},
	),
}

func MapDBTypeToVarType(dType string) string {
	for _, m := range matchers {
		match, ok := m.Match(dType)
		if ok {
			return match
		}
	}
	return "interface{}"
}

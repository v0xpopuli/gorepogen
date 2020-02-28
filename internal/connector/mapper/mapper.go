package mapper

import (
	"github.com/v0xpopuli/gorepogen/internal/connector"
	"github.com/v0xpopuli/gorepogen/internal/param"
)

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
			"date", "time", "datetime",
			"timestamp", "year", "interval",
			"timestamptz", "timestamp with time zone",
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

func MapTablesToEntityDefinition(tables []connector.Table) (param.EntityDefinition, error) {
	entityDefinition := make(param.EntityDefinition, 0)
	for _, t := range tables {
		info := param.EntityInfo{
			EntityName:      t.TableName,
			EntityPackage:   "",
			FullPackagePath: "",
		}
		entityDefinition[info] = append(entityDefinition[info], param.Field{
			VarName: t.ColumnName,
			VarType: mapColumnTypeToVarType(t.ColumnType),
		})
	}
	return entityDefinition, nil
}

func mapColumnTypeToVarType(columnType string) string {
	for _, m := range matchers {
		match, ok := m.Match(columnType)
		if ok {
			return match
		}
	}
	return "interface{}"
}

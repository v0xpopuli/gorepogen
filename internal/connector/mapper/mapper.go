package mapper

import (
	"github.com/v0xpopuli/gorepogen/internal/connector"
	"github.com/v0xpopuli/gorepogen/internal/generator/entity"
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

func MapTablesToEntityDefinition(tables []connector.Table) (entity.Definition, error) {
	entityDefinition := make(entity.Definition, 0)
	for _, t := range tables {
		entityDefinition[t.TableName] = append(entityDefinition[t.TableName], entity.Field{
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

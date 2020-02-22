package entity

import (
	"os"
	"path/filepath"

	"github.com/v0xpopuli/gorepogen/internal/driver"
)

type generator struct{}

func NewGenerator() *generator {
	return &generator{}
}

type field struct {
	name  string
	dtype string
}

func (g generator) Generate(tables []driver.TableInfo) (string, error) {

	g.mapTableInfoToEntityInfo(tables)

	return "", nil
}

func (g generator) makeEntityDir(fileName, currentDir, entityPackageName string) string {
	entityDir := filepath.Join(currentDir, entityPackageName)
	_ = os.MkdirAll(entityDir, os.ModePerm)
	return filepath.Join(entityDir, fileName)
}

func (g generator) mapTableInfoToEntityInfo(tableInfo []driver.TableInfo) map[string][]field {

	var name string
	var fields []field
	map1 := make(map[string][]field, 0)
	for i, t := range tableInfo {
		if name == "" {
			name = t.TableName
		}
		if name == t.TableName {
			fields = append(fields, field{
				name:  t.ColumnName,
				dtype: t.DataType,
			})
		} else {
			map1[name] = fields
			fields = make([]field, 0)
			name = t.TableName
		}
		if i == len(tableInfo)-1 {
			map1[name] = fields
		}
	}

	return nil
}

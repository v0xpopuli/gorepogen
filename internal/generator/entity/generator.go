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

type entityInfo struct {
	name   string
	fields []field
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

// not the most elegant solution, PR's are welcome
func (g generator) mapTableInfoToEntityInfo(tableInfo []driver.TableInfo) (infos []entityInfo) {
	// TODO: great potential to refactoring
	var info entityInfo
	for i, t := range tableInfo {
		if info.name == "" {
			info.name = t.TableName
		}
		if info.name == t.TableName {
			info.fields = append(info.fields, field{
				name:  t.ColumnName,
				dtype: t.DataType,
			})
		} else {
			infos = append(infos, info)
			info.fields = make([]field, 0)
			info.name = t.TableName
		}
		if i == len(tableInfo)-1 {
			infos = append(infos, info)
		}
	}

	return infos
}

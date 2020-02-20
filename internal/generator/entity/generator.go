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
	fields map[string]string
}

func (g generator) Generate(tables []driver.TableInfo) (string, error) {
	return "", nil
}

func (g generator) makeEntityDir(fileName, currentDir, entityPackageName string) string {
	entityDir := filepath.Join(currentDir, entityPackageName)
	_ = os.MkdirAll(entityDir, os.ModePerm)
	return filepath.Join(entityDir, fileName)
}

func (g generator) mapTableInfoToEntityInfo(tableInfo []driver.TableInfo) (entitiesInfo []*entityInfo) {
	return entitiesInfo
}

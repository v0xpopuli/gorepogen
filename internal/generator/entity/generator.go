package entity

import (
	"os"
	"path/filepath"

	"github.com/v0xpopuli/gorepogen/internal/dialect"
)

type generator struct{}

func NewGenerator() *generator {
	return &generator{}
}

func (g generator) Generate(tables map[string][]dialect.Field) (string, error) {
	return "", nil
}

func (g generator) makeEntityDir(fileName, currentDir, entityPackageName string) string {
	entityDir := filepath.Join(currentDir, entityPackageName)
	_ = os.MkdirAll(entityDir, os.ModePerm)
	return filepath.Join(entityDir, fileName)
}

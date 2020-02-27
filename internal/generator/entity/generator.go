package entity

import (
	"os"
	"path/filepath"
)

type generator struct{}

func NewGenerator() *generator {
	return &generator{}
}

func (g generator) Generate(entityDefinition Definition) (string, error) {
	return "", nil
}

func (g generator) makeEntityDir(fileName, currentDir, entityPackageName string) string {
	entityDir := filepath.Join(currentDir, entityPackageName)
	_ = os.MkdirAll(entityDir, os.ModePerm)
	return filepath.Join(entityDir, fileName)
}

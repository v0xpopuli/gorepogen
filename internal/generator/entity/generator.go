package entity

import (
	"os"
	"path/filepath"
)

type generator struct {
	outputDir        string
	entityDefinition Definition
}

func NewGenerator(entityDefinition Definition, outputDir string) *generator {
	return &generator{
		entityDefinition: entityDefinition,
		outputDir:        outputDir,
	}
}

func (g generator) Generate() (string, error) {

	return "", nil
}

func (g generator) makeEntityDir(fileName, currentDir, entityPackageName string) string {
	if g.outputDir != "" {
		return filepath.Join(g.outputDir, fileName)
	}
	entityDir := filepath.Join(currentDir, entityPackageName)
	_ = os.MkdirAll(entityDir, os.ModePerm)
	return filepath.Join(entityDir, fileName)
}

package generator

import (
	"os"
	"path/filepath"

	rc "github.com/v0xpopuli/gorepogen/internal/repocomp"

	j "github.com/dave/jennifer/jen"
)

type generator struct {
	namesRegistry NamesRegister
	components    []rc.Appender
}

// NewGenerator make new instance of generator
func NewGenerator(namesRegistry NamesRegister) *generator {
	return &generator{
		namesRegistry: namesRegistry,
		components: []rc.Appender{
			rc.NewInterface(namesRegistry.GetInterfaceNames()),
			rc.NewStruct(namesRegistry.GetStructNames()),
			rc.NewConstructor(namesRegistry.GetConstructorNames()),
			rc.NewMethodsList(namesRegistry.GetMethodListNames()),
		},
	}
}

// Generate perform final rendering from all components
func (g generator) Generate(currentDir string) (string, error) {

	repositoryPackageName := g.namesRegistry.repositoryPackageName
	file := g.newRepository(repositoryPackageName, g.namesRegistry)

	for _, c := range g.components {
		c.AppendTo(file)
	}

	repositoryFullPath := g.makeRepositoryDir(g.namesRegistry.fileName, currentDir, repositoryPackageName)
	return repositoryFullPath, file.Save(repositoryFullPath)
}

func (g generator) newRepository(repositoryPackageName string, namesRegistry NamesRegister) *j.File {
	file := j.NewFile(repositoryPackageName)
	file.HeaderComment("THIS FILE IS AUTOGENERATED.\nFEEL FREE TO CHANGE IT.\nGOREPOGEN 1.0.0")
	file.ImportNames(
		map[string]string{
			"github.com/jinzhu/gorm":      "gorm",
			namesRegistry.fullPackageName: namesRegistry.packageName,
		},
	)
	return file
}

func (g generator) makeRepositoryDir(fileName, currentDir, repositoryPackageName string) string {
	repositoryDir := filepath.Join(currentDir, repositoryPackageName)
	_ = os.MkdirAll(repositoryDir, os.ModePerm)
	return filepath.Join(repositoryDir, fileName)
}

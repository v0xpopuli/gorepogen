package param

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type (
	// EntityInfo hold all information needed to build NameRegister
	EntityInfo struct {
		EntityName      string
		EntityPackage   string
		FullPackagePath string
	}

	Field struct {
		VarName string
		VarType string
	}

	EntityDefinition map[string][]Field

	InterfaceParams struct {
		InterfaceName   string
		EntityName      string
		FullPackageName string
	}

	StructParams struct {
		StructName string
	}

	ConstructorParams struct {
		ConstructorName string
		InterfaceName   string
		StructName      string
	}

	MethodListParams struct {
		ReceiverName          string
		EntityNameWithPackage string
	}

	GeneratorParams struct {
		FileName              string
		PackageName           string
		FullPackageName       string
		OutputDirectory       string
		RepositoryPackageName string
	}
)

func NewInterfaceParams(info *EntityInfo) *InterfaceParams {
	entityName := info.EntityName
	return &InterfaceParams{
		EntityName:      entityName,
		FullPackageName: info.FullPackagePath,
		InterfaceName:   fmt.Sprintf("%sRepository", entityName),
	}
}

func NewStructParams(info *EntityInfo) *StructParams {
	return &StructParams{
		StructName: fmt.Sprintf("%sRepository", strings.ToLower(info.EntityName)),
	}
}

func NewConstructorParams(info *EntityInfo) *ConstructorParams {
	entityName := info.EntityName
	entityNameUncapitalized := strings.ToLower(entityName)
	return &ConstructorParams{
		ConstructorName: fmt.Sprintf("New%sRepository", entityName),
		InterfaceName:   fmt.Sprintf("%sRepository", entityName),
		StructName:      fmt.Sprintf("%sRepository", entityNameUncapitalized),
	}
}

func NewMethodListParams(info *EntityInfo) *MethodListParams {
	entityName := info.EntityName
	entityNameUncapitalized := strings.ToLower(entityName)
	return &MethodListParams{
		ReceiverName:          fmt.Sprintf("r %sRepository", entityNameUncapitalized),
		EntityNameWithPackage: fmt.Sprintf("%s.%s", info.EntityPackage, entityName),
	}
}

func NewGeneratorParams(info *EntityInfo, outputDir string) *GeneratorParams {
	entityName := info.EntityName
	entityPackage := info.EntityPackage
	entityNameUncapitalized := strings.ToLower(entityName)
	repo := "repository"
	fileName := fmt.Sprintf("%s_repository.go", entityNameUncapitalized)
	return &GeneratorParams{
		FileName:              fileName,
		PackageName:           entityPackage,
		FullPackageName:       fmt.Sprintf("%s.%s", entityPackage, entityName),
		OutputDirectory:       resolveOutputDir(repo, outputDir, fileName),
		RepositoryPackageName: repo,
	}
}

func resolveOutputDir(repo, output, fileName string) string {
	if output != "" {
		filepath.Join(output, fileName)
	}
	pwd, _ := os.Getwd()
	pwd = filepath.Join(pwd, repo)
	_ = os.MkdirAll(pwd, os.ModePerm)
	return filepath.Join(pwd, fileName)
}

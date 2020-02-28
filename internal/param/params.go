package param

import (
	"fmt"
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

func NewGeneratorParams(info *EntityInfo) *GeneratorParams {
	entityName := info.EntityName
	entityPackage := info.EntityPackage
	entityNameUncapitalized := strings.ToLower(entityName)
	return &GeneratorParams{
		FileName:              fmt.Sprintf("%s_repository.go", entityNameUncapitalized),
		PackageName:           entityPackage,
		FullPackageName:       fmt.Sprintf("%s.%s", entityPackage, entityName),
		RepositoryPackageName: "repository",
	}
}

package param

import (
	"fmt"
	"strings"

	"github.com/v0xpopuli/gorepogen/internal/walker"
)

// TODO: temporary place
type (
	Field struct {
		VarName string
		VarType string
	}

	Definition map[string][]Field
)

type InterfaceParams struct {
	InterfaceName   string
	EntityName      string
	FullPackageName string
}

type StructParams struct {
	StructName string
}

type ConstructorParams struct {
	ConstructorName string
	InterfaceName   string
	StructName      string
}

type MethodListParams struct {
	ReceiverName          string
	EntityNameWithPackage string
}

type GeneratorParams struct {
	FileName              string
	PackageName           string
	FullPackageName       string
	RepositoryPackageName string
}

func NewInterfaceParams(info *walker.EntityInfo) *InterfaceParams {
	entityName := info.EntityName
	return &InterfaceParams{
		EntityName:      entityName,
		FullPackageName: info.FullPackagePath,
		InterfaceName:   fmt.Sprintf("%sRepository", entityName),
	}
}

func NewStructParams(info *walker.EntityInfo) *StructParams {
	return &StructParams{
		StructName: fmt.Sprintf("%sRepository", strings.ToLower(info.EntityName)),
	}
}

func NewConstructorParams(info *walker.EntityInfo) *ConstructorParams {
	entityName := info.EntityName
	entityNameUncapitalized := strings.ToLower(entityName)
	return &ConstructorParams{
		ConstructorName: fmt.Sprintf("New%sRepository", entityName),
		InterfaceName:   fmt.Sprintf("%sRepository", entityName),
		StructName:      fmt.Sprintf("%sRepository", entityNameUncapitalized),
	}
}

func NewMethodListParams(info *walker.EntityInfo) *MethodListParams {
	entityName := info.EntityName
	entityNameUncapitalized := strings.ToLower(entityName)
	return &MethodListParams{
		ReceiverName:          fmt.Sprintf("r %sRepository", entityNameUncapitalized),
		EntityNameWithPackage: fmt.Sprintf("%s.%s", info.EntityPackage, entityName),
	}
}

func NewGeneratorParams(info *walker.EntityInfo) *GeneratorParams {
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

package generator

import (
	"fmt"
	"strings"
)

// NamesRegister provides all existing names
// needed for auto generation
type NamesRegister struct {
	EntityName            string
	PackageName           string
	FullPackageName       string
	EntityNameWithPackage string
	InterfaceName         string
	StructName            string
	ConstructorName       string
	ReceiveName           string
	FileName              string
	RepositoryPackageName string
}

// NewNamesRegister build full list of needed names
// needed for auto generation
func NewNamesRegister(entityInfo *EntityInfo) NamesRegister {
	entityNameUncapitalized := strings.ToLower(entityInfo.EntityName)
	entityName := entityInfo.EntityName
	return NamesRegister{
		EntityName:            entityName,
		PackageName:           entityInfo.EntityPackage,
		FullPackageName:       entityInfo.FullPackagePath,
		EntityNameWithPackage: fmt.Sprintf("%s.%s", entityInfo.EntityPackage, entityName),
		InterfaceName:         fmt.Sprintf("%sRepository", entityName),
		ConstructorName:       fmt.Sprintf("New%sRepository", entityName),
		StructName:            fmt.Sprintf("%sRepository", entityNameUncapitalized),
		ReceiveName:           fmt.Sprintf("r %sRepository", entityNameUncapitalized),
		FileName:              fmt.Sprintf("%s_repository.go", entityNameUncapitalized),
		RepositoryPackageName: "repository",
	}
}

// GetInterfaceNames returns all names belongs to interface block
func (nr *NamesRegister) GetInterfaceNames() (string, string, string) {
	return nr.InterfaceName, nr.EntityName, nr.FullPackageName
}

// GetStructNames returns all names belongs to struct block
func (nr *NamesRegister) GetStructNames() string {
	return nr.StructName
}

// GetConstructorNames returns all names belongs to constructor block
func (nr *NamesRegister) GetConstructorNames() (string, string, string) {
	return nr.ConstructorName, nr.InterfaceName, nr.StructName
}

// GetMethodListNames returns all names belongs to method list
func (nr *NamesRegister) GetMethodListNames() (string, string) {
	return nr.ReceiveName, nr.EntityNameWithPackage
}

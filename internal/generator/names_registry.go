package generator

import (
	"fmt"
	"strings"
)

// NamesRegister provides all existing names
// needed for auto generation
type NamesRegister struct {
	entityName            string
	packageName           string
	fullPackageName       string
	entityNameWithPackage string
	interfaceName         string
	structName            string
	constructorName       string
	receiveName           string
	fileName              string
	repositoryPackageName string
}

// NewNamesRegister build full list of needed names
// needed for auto generation
func NewNamesRegister(entityInfo *EntityInfo) NamesRegister {
	entityNameUncapitalized := strings.ToLower(entityInfo.EntityName)
	entityName := entityInfo.EntityName
	return NamesRegister{
		entityName:            entityName,
		packageName:           entityInfo.EntityPackage,
		fullPackageName:       entityInfo.FullPackagePath,
		entityNameWithPackage: fmt.Sprintf("%s.%s", entityInfo.EntityPackage, entityName),
		interfaceName:         fmt.Sprintf("%sRepository", entityName),
		constructorName:       fmt.Sprintf("New%sRepository", entityName),
		structName:            fmt.Sprintf("%sRepository", entityNameUncapitalized),
		receiveName:           fmt.Sprintf("r %sRepository", entityNameUncapitalized),
		fileName:              fmt.Sprintf("%s_repository.go", entityNameUncapitalized),
		repositoryPackageName: "repository",
	}
}

// GetInterfaceNames returns all names belongs to interface block
func (nr *NamesRegister) GetInterfaceNames() (string, string, string) {
	return nr.interfaceName, nr.entityName, nr.fullPackageName
}

// GetStructNames returns all names belongs to struct block
func (nr *NamesRegister) GetStructNames() string {
	return nr.structName
}

// GetConstructorNames returns all names belongs to constructor block
func (nr *NamesRegister) GetConstructorNames() (string, string, string) {
	return nr.constructorName, nr.interfaceName, nr.structName
}

// GetMethodListNames returns all names belongs to method list
func (nr *NamesRegister) GetMethodListNames() (string, string) {
	return nr.receiveName, nr.entityNameWithPackage
}

package generator

// NamesRegistry provides all existing names
// needed for auto generation
type NamesRegistry struct {
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

// GetInterfaceNames returns all names belongs to interface block
func (nr *NamesRegistry) GetInterfaceNames() (string, string, string) {
	return nr.InterfaceName, nr.EntityName, nr.FullPackageName
}

// GetStructNames returns all names belongs to struct block
func (nr *NamesRegistry) GetStructNames() string {
	return nr.StructName
}

// GetConstructorNames returns all names belongs to constructor block
func (nr *NamesRegistry) GetConstructorNames() (string, string, string) {
	return nr.ConstructorName, nr.InterfaceName, nr.StructName
}

// GetMethodListNames returns all names belongs to method list
func (nr *NamesRegistry) GetMethodListNames() (string, string) {
	return nr.ReceiveName, nr.EntityNameWithPackage
}

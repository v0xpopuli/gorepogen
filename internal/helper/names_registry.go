package helper

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

func (nr *NamesRegistry) GetInterfaceNames() (string, string, string) {
	return nr.InterfaceName, nr.EntityName, nr.FullPackageName
}

func (nr *NamesRegistry) GetStructNames() string {
	return nr.StructName
}

func (nr *NamesRegistry) GetConstructorNames() (string, string, string) {
	return nr.ConstructorName, nr.InterfaceName, nr.StructName
}

func (nr *NamesRegistry) GetMethodListNames() (string, string) {
	return nr.ReceiveName, nr.EntityNameWithPackage
}

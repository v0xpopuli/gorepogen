package mapper

type TypeMatcher interface {
	Match(string) (string, bool)
}

type MappingType struct {
	varType     string
	columnTypes []string
}

func NewMappingType(varType string, columnTypes []string) TypeMatcher {
	return &MappingType{varType: varType, columnTypes: columnTypes}
}

func (m MappingType) Match(columnType string) (string, bool) {
	for _, e := range m.columnTypes {
		if e == columnType {
			return m.varType, true
		}
	}
	return "", false
}

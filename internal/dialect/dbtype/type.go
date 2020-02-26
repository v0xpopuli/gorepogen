package dbtype

type TypeMatcher interface {
	Match(string) (string, bool)
}

type MappingType struct {
	vType  string
	dTypes []string
}

func NewMappingType(vType string, dTypes []string) TypeMatcher {
	return &MappingType{vType: vType, dTypes: dTypes}
}

func (m MappingType) Match(dType string) (string, bool) {
	if m.contains(m.dTypes, dType) {
		return m.vType, true
	}
	return "", false
}

func (m MappingType) contains(a []string, t string) bool {
	for _, e := range a {
		if e == t {
			return true
		}
	}
	return false
}

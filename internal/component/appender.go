package component

import (
	j "github.com/dave/jennifer/jen"
	"github.com/v0xpopuli/gorepogen/internal/param"
)

// Appender is implemented by repository components,
// to provide ability to append they in the right order
type Appender interface {
	AppendTo(*j.File)
}

type components struct {
	components []Appender
}

func New(info *param.EntityInfo, withEntity bool) *components {
	return &components{
		components: setUpComponents(info, withEntity),
	}
}

func setUpComponents(info *param.EntityInfo, withEntity bool) []Appender {
	components := []Appender{
		NewInterface(param.NewInterfaceParams(info)),
		NewStruct(param.NewStructParams(info)),
		NewConstructor(param.NewConstructorParams(info)),
		NewMethodsList(param.NewMethodListParams(info)),
	}
	if withEntity {
		components = append([]Appender{NewStruct(param.NewStructParams(info))}, components...)
	}
	return components
}

func (c components) GetComponents() []Appender {
	return c.components
}

package component

import (
	j "github.com/dave/jennifer/jen"
	"github.com/v0xpopuli/gorepogen/internal/param"
)

type structGenerator struct {
	*param.StructParams
	fields []*j.Statement
}

// NewStruct renders struct block
func NewStruct(params *param.StructParams, fields ...*j.Statement) Appender {
	return &structGenerator{
		StructParams: params,
		fields:       fields,
	}
}

func (sg *structGenerator) AppendTo(file *j.File) {
	file.Type().
		Id(sg.StructName).
		StructFunc(func(group *j.Group) {
			group.Add(j.Op("*").Qual("github.com/jinzhu/gorm", "DB"))
			for _, f := range sg.fields {
				group.Add(f)
			}
		}).
		Line()
}

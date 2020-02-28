package component

import (
	j "github.com/dave/jennifer/jen"
	"github.com/v0xpopuli/gorepogen/internal/param"
)

type constructorGenerator struct {
	*param.ConstructorParams
	args       map[*j.Statement]*j.Statement
	structArgs j.Dict
}

// NewConstructor renders constructor block
func NewConstructor(params *param.ConstructorParams) Appender {
	return &constructorGenerator{
		ConstructorParams: params,
		args: map[*j.Statement]*j.Statement{
			j.Id("db"): j.Id("*gorm.DB"),
		},
		structArgs: j.Dict{
			j.Id("DB"): j.Id("db"),
		},
	}
}

func (cg *constructorGenerator) AppendTo(file *j.File) {
	file.Func().
		Id(cg.ConstructorName).
		ParamsFunc(cg.generateConstructorParams()).
		Params(j.Id(cg.InterfaceName)).
		Block(j.Return(j.Op("&").Id(cg.StructName).Values(cg.structArgs))).
		Line()
}

func (cg *constructorGenerator) generateConstructorParams() func(group *j.Group) {
	return func(group *j.Group) {
		for k, v := range cg.args {
			group.Add(k).Add(v)
		}
	}
}

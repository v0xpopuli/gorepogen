package repocomp

import (
	. "gorepogen/internal/helper"

	. "github.com/dave/jennifer/jen"
)

type constructorGenerator struct {
	constructorName string
	interfaceName   string
	structName      string
	args            map[*Statement]*Statement
	structArgs      Dict
}

func NewConstructor(constructorName, interfaceName, structName string) Appender {
	return &constructorGenerator{
		constructorName: constructorName,
		interfaceName:   interfaceName,
		structName:      structName,
		args: map[*Statement]*Statement{
			Id("db"): Id("*gorm.DB"),
		},
		structArgs: Dict{
			Id("DB"): Id("db"),
		},
	}
}

func (cg *constructorGenerator) AppendTo(file *File) {
	file.Func().
		Id(cg.constructorName).
		ParamsFunc(cg.generateConstructorParams()).
		Params(Id(cg.interfaceName)).
		Block(Return(Op("&").Id(cg.structName).Values(cg.structArgs))).
		Line()
}

func (cg *constructorGenerator) generateConstructorParams() func(group *Group) {
	return func(group *Group) {
		for k, v := range cg.args {
			group.Add(k).Add(v)
		}
	}
}

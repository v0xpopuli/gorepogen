package repocomp

import (
	h "gorepogen/internal/helper"

	j "github.com/dave/jennifer/jen"
)

type constructorGenerator struct {
	constructorName string
	interfaceName   string
	structName      string
	args            map[*j.Statement]*j.Statement
	structArgs      j.Dict
}

func NewConstructor(constructorName, interfaceName, structName string) h.Appender {
	return &constructorGenerator{
		constructorName: constructorName,
		interfaceName:   interfaceName,
		structName:      structName,
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
		Id(cg.constructorName).
		ParamsFunc(cg.generateConstructorParams()).
		Params(j.Id(cg.interfaceName)).
		Block(j.Return(j.Op("&").Id(cg.structName).Values(cg.structArgs))).
		Line()
}

func (cg *constructorGenerator) generateConstructorParams() func(group *j.Group) {
	return func(group *j.Group) {
		for k, v := range cg.args {
			group.Add(k).Add(v)
		}
	}
}

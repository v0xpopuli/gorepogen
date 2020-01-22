package repocomp

import (
	j "github.com/dave/jennifer/jen"
)

type interfaceGenerator struct {
	interfaceName   string
	entityName      string
	fullPackageName string
}

// NewInterface renders interface block
func NewInterface(interfaceName, entityName, fullPackageName string) Appender {
	return &interfaceGenerator{
		interfaceName:   interfaceName,
		entityName:      entityName,
		fullPackageName: fullPackageName,
	}
}

func (ig *interfaceGenerator) AppendTo(file *j.File) {
	file.Type().
		Id(ig.interfaceName).
		Interface(
			j.Id(findAllMethodName).
				Params().
				Params(j.List(j.Index().Qual(ig.fullPackageName, ig.entityName), j.Error())),
			j.Id(findByIdMethodName).
				Params(j.Uint()).
				Params(j.List(j.Qual(ig.fullPackageName, ig.entityName), j.Error())),
			j.Id(saveMethodName).
				Params(j.Qual(ig.fullPackageName, ig.entityName)).
				Params(j.Qual(ig.fullPackageName, ig.entityName), j.Error()),
			j.Id(updateMethodName).
				Params(j.Qual(ig.fullPackageName, ig.entityName)).
				Params(j.Error()),
			j.Id(deleteMethodName).
				Params(j.Qual(ig.fullPackageName, ig.entityName)).Error(),
			j.Id(countMethodName).
				Params().
				Params(j.Uint(), j.Error()),
		).
		Line()
}

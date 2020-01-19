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
			j.Id(findAllMethod).
				Params().
				Params(j.List(j.Index().Qual(ig.fullPackageName, ig.entityName), j.Error())),
			j.Id(findByIdMethod).
				Params(j.Uint()).
				Params(j.List(j.Qual(ig.fullPackageName, ig.entityName), j.Error())),
			j.Id(saveMethod).
				Params(j.Qual(ig.fullPackageName, ig.entityName)).
				Params(j.Qual(ig.fullPackageName, ig.entityName), j.Error()),
			j.Id(updateMethod).
				Params(j.Qual(ig.fullPackageName, ig.entityName)).
				Params(j.Error()),
			j.Id(deleteMethod).
				Params(j.Qual(ig.fullPackageName, ig.entityName)).Error(),
			j.Id(countMethod).
				Params().
				Params(j.Uint(), j.Error()),
		).
		Line()
}

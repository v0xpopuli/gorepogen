package repocomp

import (
	. "gorepogen/internal/helper"

	. "github.com/dave/jennifer/jen"
)

type interfaceGenerator struct {
	interfaceName   string
	entityName      string
	fullPackageName string
}

func NewInterface(interfaceName, entityName, fullPackageName string) Appender {
	return &interfaceGenerator{
		interfaceName:   interfaceName,
		entityName:      entityName,
		fullPackageName: fullPackageName,
	}
}

func (ig *interfaceGenerator) AppendTo(file *File) {
	file.Type().
		Id(ig.interfaceName).
		Interface(
			Id(findAllMethod).
				Params().
				Params(List(Index().Qual(ig.fullPackageName, ig.entityName), Error())),
			Id(findByIdMethod).
				Params(Uint()).
				Params(List(Qual(ig.fullPackageName, ig.entityName), Error())),
			Id(saveMethod).
				Params(Qual(ig.fullPackageName, ig.entityName)).
				Params(Qual(ig.fullPackageName, ig.entityName), Error()),
			Id(updateMethod).
				Params(Qual(ig.fullPackageName, ig.entityName)).
				Params(Error()),
			Id(deleteMethod).
				Params(Qual(ig.fullPackageName, ig.entityName)).Error(),
			Id(countMethod).
				Params().
				Params(Uint(), Error()),
		).
		Line()
}

package component

import (
	j "github.com/dave/jennifer/jen"
	"github.com/v0xpopuli/gorepogen/internal/param"
)

type interfaceGenerator struct {
	*param.InterfaceParams
}

// NewInterface renders interface block
func NewInterface(params *param.InterfaceParams) Appender {
	return &interfaceGenerator{
		InterfaceParams: params,
	}
}

func (ig *interfaceGenerator) AppendTo(file *j.File) {
	file.Type().
		Id(ig.InterfaceName).
		Interface(
			j.Id(findAllMethodName).
				Params().
				Params(j.List(j.Index().Qual(ig.FullPackageName, ig.EntityName), j.Error())),
			j.Id(findByIdMethodName).
				Params(j.Uint()).
				Params(j.List(j.Qual(ig.FullPackageName, ig.EntityName), j.Error())),
			j.Id(saveMethodName).
				Params(j.Qual(ig.FullPackageName, ig.EntityName)).
				Params(j.Qual(ig.FullPackageName, ig.EntityName), j.Error()),
			j.Id(updateMethodName).
				Params(j.Qual(ig.FullPackageName, ig.EntityName)).
				Params(j.Error()),
			j.Id(deleteMethodName).
				Params(j.Qual(ig.FullPackageName, ig.EntityName)).Error(),
			j.Id(countMethodName).
				Params().
				Params(j.Uint(), j.Error()),
		).
		Line()
}

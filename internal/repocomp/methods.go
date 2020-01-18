package repocomp

import (
	h "gorepogen/internal/helper"

	j "github.com/dave/jennifer/jen"
)

const (
	findAllMethod  = "FindAll"
	findByIdMethod = "FindById"
	saveMethod     = "Save"
	updateMethod   = "Update"
	deleteMethod   = "Delete"
	countMethod    = "Count"
)

var (
	idLit       = "id"
	entityLit   = "entity"
	entitiesLit = "entities"
	errLit      = "err"
	countLit    = "count"

	errorLit = "Error"
)

type method struct {
	receiverName *j.Statement
	methodName   string
	args         []h.StatementPair
	returnParams *j.Statement
	fnBody       *j.Statement
}

type methodsList struct {
	receiverName          string
	entityNameWithPackage string
	methods               []method
}

func NewMethodsList(receiverName, entityNameWithPackage string) h.Appender {
	return &methodsList{
		receiverName:          receiverName,
		entityNameWithPackage: entityNameWithPackage,
		methods: []method{
			{
				receiverName: j.Id(receiverName),
				methodName:   findAllMethod,
				args:         nil,
				returnParams: j.List(j.Index().Id(entityNameWithPackage), j.Error()),
				fnBody:       findAllFuncBody(entityNameWithPackage),
			},
			{
				receiverName: j.Id(receiverName),
				methodName:   findByIdMethod,
				args: []h.StatementPair{
					{
						ArgName: j.Id(idLit),
						ArgType: j.Uint(),
					},
				},
				returnParams: j.List(j.Id(entityNameWithPackage), j.Error()),
				fnBody:       findByIdFuncBody(entityNameWithPackage),
			},
			{
				receiverName: j.Id(receiverName),
				methodName:   saveMethod,
				args: []h.StatementPair{
					{
						ArgName: j.Id(entityLit),
						ArgType: j.Id(entityNameWithPackage),
					},
				},
				returnParams: j.List(j.Id(entityNameWithPackage), j.Error()),
				fnBody:       saveFuncBody(),
			},
			{
				receiverName: j.Id(receiverName),
				methodName:   updateMethod,
				args: []h.StatementPair{
					{
						ArgName: j.Id(entityLit),
						ArgType: j.Id(entityNameWithPackage),
					},
				},
				returnParams: j.Error(),
				fnBody:       updateFuncBody(),
			},
			{
				receiverName: j.Id(receiverName),
				methodName:   deleteMethod,
				args: []h.StatementPair{
					{
						ArgName: j.Id(entityLit),
						ArgType: j.Id(entityNameWithPackage),
					},
				},
				returnParams: j.Error(),
				fnBody:       deleteFuncBody(),
			},
			{
				receiverName: j.Id(receiverName),
				methodName:   countMethod,
				args:         nil,
				returnParams: j.List(j.Uint(), j.Error()),
				fnBody:       countFuncBody(entityNameWithPackage),
			},
		},
	}
}

func (ml *methodsList) AppendTo(file *j.File) {
	for _, m := range ml.methods {
		file.Func().
			Params(m.receiverName).
			Id(m.methodName).
			ParamsFunc(ml.generateMethodParams(m.args)).
			Params(m.returnParams).
			Block(m.fnBody).
			Line()
	}
}

func (ml *methodsList) generateMethodParams(args []h.StatementPair) func(group *j.Group) {
	return func(group *j.Group) {
		for _, p := range args {
			group.Add(p.ArgName).Add(p.ArgType)
		}
	}
}

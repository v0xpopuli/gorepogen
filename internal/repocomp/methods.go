package repocomp

import (
	. "gorepogen/internal/helper"

	. "github.com/dave/jennifer/jen"
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
	receiverName *Statement
	methodName   string
	args         []StatementPair
	returnParams *Statement
	fnBody       *Statement
}

type methodsList struct {
	receiverName          string
	entityNameWithPackage string
	methods               []method
}

func NewMethodsList(receiverName, entityNameWithPackage string) Appender {
	return &methodsList{
		receiverName:          receiverName,
		entityNameWithPackage: entityNameWithPackage,
		methods: []method{
			{
				receiverName: Id(receiverName),
				methodName:   findAllMethod,
				args:         nil,
				returnParams: List(Index().Id(entityNameWithPackage), Error()),
				fnBody:       findAllFuncBody(entityNameWithPackage),
			},
			{
				receiverName: Id(receiverName),
				methodName:   findByIdMethod,
				args: []StatementPair{
					{
						ArgName: Id(idLit),
						ArgType: Uint(),
					},
				},
				returnParams: List(Id(entityNameWithPackage), Error()),
				fnBody:       findByIdFuncBody(entityNameWithPackage),
			},
			{
				receiverName: Id(receiverName),
				methodName:   saveMethod,
				args: []StatementPair{
					{
						ArgName: Id(entityLit),
						ArgType: Id(entityNameWithPackage),
					},
				},
				returnParams: List(Id(entityNameWithPackage), Error()),
				fnBody:       saveFuncBody(),
			},
			{
				receiverName: Id(receiverName),
				methodName:   updateMethod,
				args: []StatementPair{
					{
						ArgName: Id(entityLit),
						ArgType: Id(entityNameWithPackage),
					},
				},
				returnParams: Error(),
				fnBody:       updateFuncBody(),
			},
			{
				receiverName: Id(receiverName),
				methodName:   deleteMethod,
				args: []StatementPair{
					{
						ArgName: Id(entityLit),
						ArgType: Id(entityNameWithPackage),
					},
				},
				returnParams: Error(),
				fnBody:       deleteFuncBody(),
			},
			{
				receiverName: Id(receiverName),
				methodName:   countMethod,
				args:         nil,
				returnParams: List(Uint(), Error()),
				fnBody:       countFuncBody(entityNameWithPackage),
			},
		},
	}
}

func (ml *methodsList) AppendTo(file *File) {
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

func (ml *methodsList) generateMethodParams(args []StatementPair) func(group *Group) {
	return func(group *Group) {
		for _, p := range args {
			group.Add(p.ArgName).Add(p.ArgType)
		}
	}
}

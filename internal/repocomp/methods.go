package repocomp

import (
	j "github.com/dave/jennifer/jen"
)

const (
	findAllMethodName  = "FindAll"
	findByIdMethodName = "FindById"
	saveMethodName     = "Save"
	updateMethodName   = "Update"
	deleteMethodName   = "Delete"
	countMethodName    = "Count"
)

var (
	idLit       = "id"
	entityLit   = "entity"
	entitiesLit = "entities"
	errLit      = "err"
	countLit    = "count"

	errorLit = "Error"
)

type statementPair struct {
	ArgName *j.Statement
	ArgType *j.Statement
}

type method struct {
	receiverName *j.Statement
	methodName   string
	args         []statementPair
	returnParams *j.Statement
	fnBody       *j.Statement
}

type methodsList struct {
	receiverName          string
	entityNameWithPackage string
	methods               []method
}

// NewMethodsList renders method list
func NewMethodsList(receiverName, entityNameWithPackage string) Appender {
	return &methodsList{
		receiverName:          receiverName,
		entityNameWithPackage: entityNameWithPackage,
		methods: []method{
			findAllMethod(receiverName, entityNameWithPackage),
			findByIdMethod(receiverName, entityNameWithPackage),
			saveMethod(receiverName, entityNameWithPackage),
			updateMethod(receiverName, entityNameWithPackage),
			deleteMethod(receiverName, entityNameWithPackage),
			countMethod(receiverName, entityNameWithPackage),
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

func (ml *methodsList) generateMethodParams(args []statementPair) func(group *j.Group) {
	return func(group *j.Group) {
		for _, p := range args {
			group.Add(p.ArgName).Add(p.ArgType)
		}
	}
}

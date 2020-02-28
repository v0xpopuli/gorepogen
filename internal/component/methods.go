package component

import (
	j "github.com/dave/jennifer/jen"
	"github.com/v0xpopuli/gorepogen/internal/param"
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
	*param.MethodListParams
	methods []method
}

// NewMethodsList renders method list
func NewMethodsList(params *param.MethodListParams) Appender {
	return &methodsList{
		MethodListParams: params,
		methods: []method{
			findAllMethod(params.ReceiverName, params.EntityNameWithPackage),
			findByIdMethod(params.ReceiverName, params.EntityNameWithPackage),
			saveMethod(params.ReceiverName, params.EntityNameWithPackage),
			updateMethod(params.ReceiverName, params.EntityNameWithPackage),
			deleteMethod(params.ReceiverName, params.EntityNameWithPackage),
			countMethod(params.ReceiverName, params.EntityNameWithPackage),
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

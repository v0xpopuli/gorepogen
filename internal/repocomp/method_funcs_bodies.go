package repocomp

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

func findAllMethod(receiverName, entityNameWithPackage string) method {
	return method{
		receiverName: j.Id(receiverName),
		methodName:   findAllMethodName,
		args:         nil,
		returnParams: j.List(j.Index().Id(entityNameWithPackage), j.Error()),
		fnBody:       findAllFuncBody(entityNameWithPackage),
	}
}

func findAllFuncBody(entityName string) *j.Statement {
	return j.Var().Id(entitiesLit).Index().Id(entityName).
		Line().
		Id(fmt.Sprintf("%s := r.DB.Find(&%s).%s", errLit, entitiesLit, errorLit)).
		Line().
		Return(j.Id(entitiesLit), j.Id(errLit))
}

func findByIdMethod(receiverName, entityNameWithPackage string) method {
	return method{
		receiverName: j.Id(receiverName),
		methodName:   findByIdMethodName,
		args: []statementPair{
			{
				ArgName: j.Id(idLit),
				ArgType: j.Uint(),
			},
		},
		returnParams: j.List(j.Id(entityNameWithPackage), j.Error()),
		fnBody:       findByIdFuncBody(entityNameWithPackage),
	}
}

func findByIdFuncBody(entityName string) *j.Statement {
	return j.Var().Id(entityLit).Id(entityName).
		Line().
		Id(fmt.Sprintf("%s := r.DB.First(&%s,%s).%s", errLit, entityLit, idLit, errorLit)).
		Line().
		Return(j.Id(entityLit), j.Id(errLit))
}

func saveMethod(receiverName, entityNameWithPackage string) method {
	return method{
		receiverName: j.Id(receiverName),
		methodName:   saveMethodName,
		args: []statementPair{
			{
				ArgName: j.Id(entityLit),
				ArgType: j.Id(entityNameWithPackage),
			},
		},
		returnParams: j.List(j.Id(entityNameWithPackage), j.Error()),
		fnBody:       saveFuncBody(),
	}
}

func saveFuncBody() *j.Statement {
	return j.Id(fmt.Sprintf("%s := r.DB.Create(&%s).%s", errLit, entityLit, errorLit)).
		Line().
		Return(j.Id(entityLit), j.Id(errLit))
}

func updateMethod(receiverName, entityNameWithPackage string) method {
	return method{
		receiverName: j.Id(receiverName),
		methodName:   updateMethodName,
		args: []statementPair{
			{
				ArgName: j.Id(entityLit),
				ArgType: j.Id(entityNameWithPackage),
			},
		},
		returnParams: j.Error(),
		fnBody:       updateFuncBody(),
	}
}

func updateFuncBody() *j.Statement {
	return j.Return(j.Id(fmt.Sprintf("r.DB.UpdateColumns(&%s).%s", entityLit, errorLit)))
}

func deleteMethod(receiverName, entityNameWithPackage string) method {
	return method{
		receiverName: j.Id(receiverName),
		methodName:   deleteMethodName,
		args: []statementPair{
			{
				ArgName: j.Id(entityLit),
				ArgType: j.Id(entityNameWithPackage),
			},
		},
		returnParams: j.Error(),
		fnBody:       deleteFuncBody(),
	}
}

func deleteFuncBody() *j.Statement {
	return j.Return(j.Id(fmt.Sprintf("r.DB.Delete(&%s).%s", entityLit, errorLit)))
}

func countMethod(receiverName, entityNameWithPackage string) method {
	return method{
		receiverName: j.Id(receiverName),
		methodName:   countMethodName,
		args:         nil,
		returnParams: j.List(j.Uint(), j.Error()),
		fnBody:       countFuncBody(entityNameWithPackage),
	}
}

func countFuncBody(entityName string) *j.Statement {
	return j.Var().Id(countLit).Uint().
		Line().
		Id(fmt.Sprintf("%s := r.DB.Model(&%s{}).Count(&%s).%s", errLit, entityName, countLit, errorLit)).
		Line().
		Return(j.Id(countLit), j.Id(errLit))
}

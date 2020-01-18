package repocomp

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

func findAllFuncBody(entityName string) *j.Statement {
	return j.Var().Id(entitiesLit).Index().Id(entityName).
		Line().
		Id(fmt.Sprintf("%s := r.DB.Find(&%s).%s", errLit, entitiesLit, errorLit)).
		Line().
		Return(j.Id(entitiesLit), j.Id(errLit))
}

func findByIdFuncBody(entityName string) *j.Statement {
	return j.Var().Id(entityLit).Id(entityName).
		Line().
		Id(fmt.Sprintf("%s := r.DB.First(&%s,%s).%s", errLit, entityLit, idLit, errorLit)).
		Line().
		Return(j.Id(entityLit), j.Id(errLit))
}

func saveFuncBody() *j.Statement {
	return j.Id(fmt.Sprintf("%s := r.DB.Create(&%s).%s", errLit, entityLit, errorLit)).
		Line().
		Return(j.Id(entityLit), j.Id(errLit))
}

func updateFuncBody() *j.Statement {
	return j.Return(j.Id(fmt.Sprintf("r.DB.UpdateColumns(&%s).%s", entityLit, errorLit)))
}

func deleteFuncBody() *j.Statement {
	return j.Return(j.Id(fmt.Sprintf("r.DB.Delete(&%s).%s", entityLit, errorLit)))
}

func countFuncBody(entityName string) *j.Statement {
	return j.Var().Id(countLit).Uint().
		Line().
		Id(fmt.Sprintf("%s := r.DB.Model(&%s{}).Count(&%s).%s", errLit, entityName, countLit, errorLit)).
		Line().
		Return(j.Id(countLit), j.Id(errLit))
}

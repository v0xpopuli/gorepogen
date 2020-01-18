package repocomp

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func findAllFuncBody(entityName string) *Statement {
	return Var().Id(entitiesLit).Index().Id(entityName).
		Line().
		Id(fmt.Sprintf("%s := r.DB.Find(&%s).%s", errLit, entitiesLit, errorLit)).
		Line().
		Return(Id(entitiesLit), Id(errLit))
}

func findByIdFuncBody(entityName string) *Statement {
	return Var().Id(entityLit).Id(entityName).
		Line().
		Id(fmt.Sprintf("%s := r.DB.First(&%s,%s).%s", errLit, entityLit, idLit, errorLit)).
		Line().
		Return(Id(entityLit), Id(errLit))
}

func saveFuncBody() *Statement {
	return Id(fmt.Sprintf("%s := r.DB.Create(&%s).%s", errLit, entityLit, errorLit)).
		Line().
		Return(Id(entityLit), Id(errLit))
}

func updateFuncBody() *Statement {
	return Return(Id(fmt.Sprintf("r.DB.UpdateColumns(&%s).%s", entityLit, errorLit)))
}

func deleteFuncBody() *Statement {
	return Return(Id(fmt.Sprintf("r.DB.Delete(&%s).%s", entityLit, errorLit)))
}

func countFuncBody(entityName string) *Statement {
	return Var().Id(countLit).Uint().
		Line().
		Id(fmt.Sprintf("%s := r.DB.Model(&%s{}).Count(&%s).%s", errLit, entityName, countLit, errorLit)).
		Line().
		Return(Id(countLit), Id(errLit))
}

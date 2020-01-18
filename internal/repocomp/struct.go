package repocomp

import (
	. "gorepogen/internal/helper"

	. "github.com/dave/jennifer/jen"
)

type structGenerator struct {
	structName string
	dbField    *Statement
}

func NewStruct(structName string) Appender {
	return &structGenerator{
		structName: structName,
		dbField:    Op("*").Qual("github.com/jinzhu/gorm", "DB"),
	}
}

func (sg *structGenerator) AppendTo(file *File) {
	file.Type().
		Id(sg.structName).
		Struct(sg.dbField).
		Line()
}

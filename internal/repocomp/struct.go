package repocomp

import (
	j "github.com/dave/jennifer/jen"
)

type structGenerator struct {
	structName string
	dbField    *j.Statement
}

// NewStruct renders struct block
func NewStruct(structName string) Appender {
	return &structGenerator{
		structName: structName,
		dbField:    j.Op("*").Qual("github.com/jinzhu/gorm", "DB"),
	}
}

func (sg *structGenerator) AppendTo(file *j.File) {
	file.Type().
		Id(sg.structName).
		Struct(sg.dbField).
		Line()
}

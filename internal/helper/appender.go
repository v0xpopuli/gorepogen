package helper

import j "github.com/dave/jennifer/jen"

type Appender interface {
	AppendTo(*j.File)
}

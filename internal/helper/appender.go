package helper

import . "github.com/dave/jennifer/jen"

type Appender interface {
	AppendTo(*File)
}

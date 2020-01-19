package repocomp

import j "github.com/dave/jennifer/jen"

// Appender is implemented by repository components,
// to provide ability to append they in the right order
type Appender interface {
	AppendTo(*j.File)
}

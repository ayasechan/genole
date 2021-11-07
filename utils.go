package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

var (
	IFErrNENil = If(Err().Op("!=").Nil())
)

func GetVtableName(name string) string {
	return fmt.Sprintf("%sVtbl", name)
}

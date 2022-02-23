package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	. "github.com/dave/jennifer/jen"
)

var (
	IFErrNENil = If(Err().Op("!=").Nil())
)

func GetVtableName(name string) string {
	return fmt.Sprintf("%sVtbl", name)
}

type ExtraPackageFlag map[string]string

var _ flag.Value = (*ExtraPackageFlag)(nil)

func (m *ExtraPackageFlag) String() string {
	return fmt.Sprintf("%+v", *m)
}

// 使用 : 分割
// 前面是包的路径
// 后面是包的别名
func (m *ExtraPackageFlag) Set(value string) error {
	frags := strings.Split(value, ":")
	switch len(frags) {
	case 1:
		(*m)[frags[0]] = ""
	case 2:
		(*m)[frags[0]] = frags[1]
	default:
		return errors.New("too many")
	}
	return nil
}

func isInteger(s string) bool {
	l := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
	}
	for _, v := range l {
		if s == v {
			return true
		}
	}
	return false
}

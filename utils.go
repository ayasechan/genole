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

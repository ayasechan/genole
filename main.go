package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	input         = flag.String("input", "", "输入的 yaml 文件")
	output        = flag.String("output", "", "输出的 go 文件")
	pkgName       = flag.String("pkg", "", "包的名称")
	extraPackages = make(ExtraPackageFlag)
)

func main() {
	flag.Var(&extraPackages, "e", "自定义额外导入的包 可使用多次 格式为 pkg/path[:alias]")
	flag.Parse()
	srcFd, err := os.Open(*input)
	if err != nil {
		panic(err)
	}

	var classes []ComClass
	err = yaml.NewDecoder(srcFd).Decode(&classes)
	if err != nil {
		panic(fmt.Errorf("decode %s error: %v", *input, err))
	}

	content := GenHeader(*pkgName)
	content.Add(GenExtraImport(extraPackages))
	for _, cls := range classes {
		content.Add(GenClass(cls))
		for _, m := range cls.Methods {
			content.Add(GenMethod(cls.Name, m))
		}
	}
	err = content.Save(*output)
	if err != nil {
		panic(err)
	}
}

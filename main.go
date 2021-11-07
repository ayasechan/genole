package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	input   = flag.String("input", "", "")
	output  = flag.String("output", "", "")
	pkgName = flag.String("pkg", "", "")
)

func main() {

	flag.Parse()
	srcFd, err := os.Open(*input)
	if err != nil {
		panic(err)
	}

	var classes []ComClass
	yaml.NewDecoder(srcFd).Decode(&classes)

	content := GenHeader(*pkgName)
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

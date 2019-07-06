package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
)

// referenced github.com/golang/tools/blob/master/cmd/stringer
type GoCodeGenerator struct {
	code     bytes.Buffer
	packages map[string]string

	generator string
	pkg       string
}

func NewGoCodeGenerator(generator string, pkg string) *GoCodeGenerator {
	g := &GoCodeGenerator{
		packages:  make(map[string]string),
		generator: generator,
		pkg:       pkg,
	}

	return g
}

func (g *GoCodeGenerator) AddPackage(alias, packagePath string) {
	g.packages[packagePath] = alias
}

func (g *GoCodeGenerator) Printf(format string, a ...interface{}) (int, error) {
	return g.code.WriteString(fmt.Sprintf(format, a...))
}

func (g GoCodeGenerator) OutputFormattedCode() ([]byte, error) {
	code := bytes.Buffer{}
	code.WriteString(fmt.Sprintf("// Code generated by \"%s\"; DO NOT EDIT.\n", g.generator))
	code.WriteString(fmt.Sprintf("\n"))
	code.WriteString(fmt.Sprintf("package %s\n", g.pkg))
	code.WriteString(fmt.Sprintf("\n"))

	// import
	code.WriteString("import (\n")
	for k, v := range g.packages {
		if v == "" {
			code.WriteString(fmt.Sprintf("	\"%s\"\n", k))
		} else {
			code.WriteString(fmt.Sprintf("	%s \"%s\"\n", v, k))
		}
	}
	code.WriteString(")\n")
	code.WriteString("\n")

	code.Write(g.code.Bytes())

	return format.Source(code.Bytes())
}

func (g GoCodeGenerator) WriteFile(filename string) error {
	src, err := g.OutputFormattedCode()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, src, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (g GoCodeGenerator) WriteWithTemplateFileName() error {
	return g.WriteFile(fmt.Sprintf("%s.go", g.generator))
}

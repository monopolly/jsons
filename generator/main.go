package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/monopolly/file"
)

func main() {

	//f := file.OpenE("test.go")

	files := os.Args
	if len(files) != 2 {
		fmt.Println("Usage:", files[0]+" file_struct.go")
		return
	}

	in := files[1]
	//in = "test.go"
	out := strings.TrimSuffix(in, ".go") + "js.go"

	fields := [][3]string{}
	var structed bool
	var name string
	var pac string
	file.PlayBytes(in, func(line []byte) {
		if bytes.Index(line, []byte("package")) != -1 {
			p := bytes.Fields(line)
			if len(p) > 1 {
				pac = string(p[1])
			}
		}
		if bytes.Index(line, []byte("struct")) != -1 {
			structed = true
			p := bytes.Fields(line)
			if len(p) < 2 {
				structed = false
				return
			}
			name = string(p[1])
			return
		}
		if !structed {
			return
		}
		f := bytes.Fields(line)
		if len(f) < 2 {
			return
		}
		var comment string
		if len(f) > 2 {
			comment = string(bytes.Join(f[2:], []byte(" ")))
		}
		fields = append(fields, [3]string{string(f[0]), string(f[1]), comment})
	})

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("package %s\n", pac))

	b.WriteString(`import "github.com/monopolly/jsons"`)
	b.WriteString("\n")

	structName := strings.Title(name)

	var functions []string

	functions = append(functions, fmt.Sprintf(`
		func New() (a %s) {
			a = []byte("{}")
			return a
		}
	`, structName))

	functions = append(functions, fmt.Sprintf("type %s []byte\n", structName))

	functions = append(functions, fmt.Sprintf(`
		func (a *%s) Set(k string, v interface{}) *%s {
			(*a) = jsons.Set((*a), k,v)
			return a
		}
	`, structName, structName))

	functions = append(functions, fmt.Sprintf(`
		func (a *%s) Get(k string) jsons.Result {
			return jsons.Get((*a), k)
		}
	`, structName))

	var constantsFields []string

	for _, x := range fields {
		fname := x[0]
		funcType := x[1]
		funcName := strings.Title(fname)
		field := fmt.Sprintf("f%s", fname) //fid, fname
		constantsFields = append(constantsFields, fmt.Sprintf(`%s = "%s" // %s %s`, field, fname, funcType, x[2]))
		if fname == "id" {
			funcName = "ID"
		}
		var jsonsValue string
		funcTypeSet := funcType
		switch funcType {
		case "int":
			jsonsValue = "Int"
		case "int64":
			jsonsValue = "Int64"
		case "string":
			jsonsValue = "String"
		case "bool":
			jsonsValue = "Bool"
		case "[]string":
			funcTypeSet = "...string"
			jsonsValue = "ArrayString"
		case "[]int":
			jsonsValue = "ArrayInt"
			funcTypeSet = "...int"
		case "map[string]string":
			jsonsValue = "MapString"
		case "map[string]int":
			jsonsValue = "MapInt"
		case "interface{}":
			jsonsValue = "Interface"
		}

		functions = append(functions, fmt.Sprintf(`//%s get value`, funcName))
		functions = append(functions, fmt.Sprintf("func (a *%s) %s() %s {return jsons.%s((*a), %s)}", structName, funcName, funcType, jsonsValue, field))
		functions = append(functions, fmt.Sprintf(`//Set%s set value`, funcName))
		functions = append(functions, fmt.Sprintf("func (a *%s) ", structName)+fmt.Sprintf("Set%s", funcName)+fmt.Sprintf("(v %s) ", funcTypeSet)+fmt.Sprintf("*%s {", structName)+fmt.Sprintf(`return a.Set(%s, v)}`, field))

		//functions = append(functions, "}")
	}

	b.WriteString(fmt.Sprintf(`
	   const (
		   %s
	   )
	`, strings.Join(constantsFields, "\n")))

	b.WriteString(strings.Join(functions, "\n"))
	//fmt.Println(string(name), fields)

	//ioutil.WriteFile(out, b.Bytes(), os.ModePerm)
	file.Save(out, b.Bytes())

}

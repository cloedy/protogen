package cmd

import (
	"os"
	"strings"
	"text/template"
)

type datam struct {
	Name  string
	Name2 string
}

const prototpl = `
message {{.Name}} {
  string {{.Name2}}_id = 1 [json_name = "{{.Name2}}_id"];
}

message {{.Name}}Item {
  {{.Name}} items = 1;
}

message {{.Name}}Request {
  string id = 1 [json_name = "id"];
}

message {{.Name}}Return {
  int64 code = 1;
  string msg = 2;
}
`

func protogen(name string) {
	name2 := name
	name = strings.ToUpper(name[0:1]) + name[1:]
	data := datam{Name: name, Name2: name2}

	tmpl, err := template.New("p3").Parse(prototpl)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}

}

package cmd

import (
    "os"
    "strings"
    "text/template"
    "unicode"
)

type datam struct {
    Name  string
    Name2 string
}

const prototpl = `
message {{.Name}} {
  string {{.Name2}}_id = 1 [json_name = "{{.Name2}}_id", (validate.rules).string.len = 32];
  string {{.Name2}}_name = 2 [json_name = "{{.Name2}}_id", (validate.rules).string = {min_len: 1, max_len: 255}];
}

message {{.Name}}Item {
  {{.Name}} items = 1;
  int64 total = 2;
}

message {{.Name}}Request {
  string {{.Name2}}_id = 1 [json_name = "{{.Name2}}_id", (validate.rules).string.len = 32];
  string {{.Name2}}_name = 2 [json_name = "{{.Name2}}_name", (validate.rules).string = {ignore_empty: true, max_len: 255}];
}

message {{.Name}}Return {
  int64 code = 1;
  string msg = 2;
}
`

func protogen(name string) {
    name2 := name
    name = strings.ToUpper(name[0:1]) + name[1:]
    name2 = camelToSnake(name)
    data := datam{Name: name, Name2: name2}

    tmpl, err := template.New("p3").Parse(prototpl)
    if err != nil {
        panic(err)
    }
    if err := tmpl.Execute(os.Stdout, data); err != nil {
        panic(err)
    }

}

func camelToSnake(s string) string {
    result := make([]rune, 0, len(s)*2)
    for i, runeValue := range s {
        if unicode.IsUpper(runeValue) {
            if i > 0 {
                result = append(result, '_')
            }
            result = append(result, unicode.ToLower(runeValue))
        } else {
            result = append(result, runeValue)
        }
    }
    return string(result)
}

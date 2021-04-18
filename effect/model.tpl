{{define "model.tpl"}}
package {{.Package}}

import "time"

// {{.Table.Comment}}
type {{.Table.Tag}} struct {
    {{range .Columns}}
    // {{.Comment}}
    {{.Name|camelCase}} {{.Type}} `{{.Tag}}`
    {{- end}}
}

func (*{{.Table.Tag}}) TableName() string {
	return "{{.Table.Name}}"
}
{{end}}
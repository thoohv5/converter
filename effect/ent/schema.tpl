{{define "schema.tpl"}}package schema

import (
    "github.com/facebook/ent/dialect"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// {{.Table.Tag}} holds the schema definition for the {{.Table.Tag}} entity.
type {{.Table.Tag}} struct {
	ent.Schema
}

// Mixin of the {{.Table.Tag}}.
func ({{.Table.Tag}}) Mixin() []ent.Mixin {
	return []ent.Mixin{
	}
}

func ({{.Table.Tag}}) Config() ent.Config {
	return ent.Config{
		Table: "{{.Table.Name}}",
	}
}

// Fields of the {{.Table.Tag}}.
func ({{.Table.Tag}}) Fields() []ent.Field {
	return []ent.Field{ {{range .Columns}}
	    {{- if eq .Name "label"}}
        field.{{.Type|camelCase}}("lab"){{if .Default}}.Default({{.Default}}){{else}}.Optional(){{end}}.Comment(`{{.Comment}}`).StorageKey("{{.Name}}"),
	    {{- else if eq .Name "id"}}
        field.{{.Type|camelCase}}("{{.Name}}").Comment(`{{.Comment}}`),
	    {{- else if eq .Type "int32" "int64"}}
        field.{{.Type|camelCase}}("{{.Name}}"){{if .Default}}.Default({{.Default}}){{else}}.Optional(){{end}}.Comment(`{{.Comment}}`),
	    {{- else if eq .Type "time"}}
        field.{{.Type|camelCase}}("{{.Name}}").Optional().SchemaType(map[string]string{ dialect.MySQL:"{{.OriginalType}}",}).Comment(`{{.Comment}}`),
	    {{- else}}
        field.{{.Type|camelCase}}("{{.Name}}"){{if .Default}}.Default("{{.Default}}"){{else}}.Optional(){{end}}.Comment(`{{.Comment}}`),
	    {{- end -}}
	{{- end -}}
	}
}

// Edges of the {{.Table.Tag}}.
func ({{.Table.Tag}}) Edges() []ent.Edge {
	return nil
}
{{end}}
{{define "schema.tpl"}}package schema

import (
    "entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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

func ({{.Table.Tag}}) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "{{.Table.Name}}"},
	}
}

// Fields of the {{.Table.Tag}}.
func ({{.Table.Tag}}) Fields() []ent.Field {
	return []ent.Field{ {{range .Columns}}
	    {{- if eq .Type "int32" "int64"}}
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
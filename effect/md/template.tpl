{{.Table.Comment}}：{{.Table.Name}}

| 字段       | 备注   | 类型    | 索引 |
| :--------- | :----- | :------ | :--- |
{{- range .Columns }}
| {{.Name}}  | {{.Comment}} | {{.OriginalType}} |      |
{{- end -}}
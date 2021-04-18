package parse

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/thoohv5/converter/datasource"
	"github.com/thoohv5/converter/effect"
	"github.com/thoohv5/converter/utils"
)

type parse struct {
	typeToGo map[string]string
	e        effect.IEffect
}

func New() IParse {
	return &parse{
		typeToGo: map[string]string{
			"int":                "int32",
			"integer":            "int32",
			"tinyint":            "int32",
			"smallint":           "int32",
			"mediumint":          "int32",
			"bigint":             "int64",
			"int unsigned":       "int32",
			"integer unsigned":   "int32",
			"tinyint unsigned":   "int32",
			"smallint unsigned":  "int32",
			"mediumint unsigned": "int32",
			"bigint unsigned":    "int64",
			"bit":                "int64",
			"bool":               "bool",
			"enum":               "string",
			"set":                "string",
			"varchar":            "string",
			"char":               "string",
			"tinytext":           "string",
			"mediumtext":         "string",
			"text":               "string",
			"longtext":           "string",
			"blob":               "string",
			"tinyblob":           "string",
			"mediumblob":         "string",
			"longblob":           "string",
			"date":               "time.Time", // time.Time or string
			"datetime":           "time.Time", // time.Time or string
			"timestamp":          "time.Time", // time.Time or string
			"time":               "time.Time", // time.Time or string
			"float":              "float64",
			"double":             "float64",
			"decimal":            "float64",
			"binary":             "string",
			"varbinary":          "string",
		},
		e: effect.New(),
	}
}

func (p *parse) Transform(columnType string) (fieldType string) {
	fieldType, ok := p.typeToGo[columnType]
	if !ok {
		fieldType = ""
	}
	return
}

func (p *parse) MakeTag(tagLabel string, columnName string, jsonTag bool) string {
	if jsonTag {
		return fmt.Sprintf(`%s:"%s" %s:"%s"`, tagLabel, columnName, "json", utils.Lcfirst(utils.CamelCase(columnName)))
	}
	return fmt.Sprintf(`%s:"%s"`, tagLabel, columnName)
}

func (p *parse) Model(param *Param) error {

	dsn := param.Dsn
	tables := param.Tables
	tag := param.Tag
	jsonTag := param.JsonTag
	savePath := param.SavePath
	pkg := param.Pkg

	svr, err := datasource.New(dsn)
	if nil != err {
		return err
	}

	// 保存文件
	if err := os.MkdirAll(savePath, os.ModePerm); nil != err {
		return fmt.Errorf("os Mkdir err, err:%w", err)
	}

	for _, table := range tables {
		// 表信息
		tableInfo, err := svr.GetTableInfo(table)
		if nil != err {
			return err
		}

		// 字段列表信息
		columnList, err := svr.GetColumnList(table)
		if nil != err {
			return err
		}

		// 数据处理
		columns := make([]*effect.Column, 0)
		for _, column := range columnList {
			columns = append(columns, &effect.Column{
				Column: &datasource.Column{
					Name:    column.Name,
					Type:    p.Transform(column.Type),
					Default: column.Default,
					Comment: column.Comment,
				},
				Tag: p.MakeTag(tag, column.Name, jsonTag),
			})
		}

		filePath := fmt.Sprintf("%s%s%s.go", savePath, string(os.PathSeparator), table)
		f, err := os.Create(filePath)
		if nil != err {
			if !os.IsExist(err) {
				return fmt.Errorf("os Create err, err:%w", err)
			}
			f, err = os.Open(filePath)
			if nil != err {
				return fmt.Errorf("os Open err, err:%w", err)
			}
		}
		defer func() {
			f.Close()
		}()

		// 模版渲染
		err = p.e.Assembly(f, &effect.Model{
			Package: pkg,
			Table: &effect.Table{
				Table: tableInfo,
				Tag:   utils.CamelCase(tableInfo.Name),
			},
			Columns: columns,
		})

		// 格式化
		cmd := exec.Command("gofmt", "-w", filePath)
		if err := cmd.Run(); nil != err {
			return err
		}
	}
	return nil
}

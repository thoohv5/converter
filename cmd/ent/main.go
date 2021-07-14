package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/thoohv5/converter/effect/adapter"
	"github.com/thoohv5/converter/parse"
)

var dsn string
var tables string
var savePath string

func init() {
	flag.StringVar(&dsn, "dsn", "root:password@tcp(127.0.0.1:3306)/med_interrogation?charset=utf8", "数据库链接DSN")
	flag.StringVar(&tables, "tables", "file", "表的名称，多个按照逗号隔开")
	flag.StringVar(&savePath, "path", "./model", "保存路径")
}

func main() {

	flag.Parse()

	err := parse.New(adapter.SchemaEffect).Schema(&parse.Param{
		Dsn:      dsn,
		Tables:   strings.Split(tables, ","),
		SavePath: savePath,
	})

	fmt.Println(err)
	log.Println(err)

}

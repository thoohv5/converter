package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/thoohv5/converter/parse"
)

var dsn string
var tables string
var pkg string
var tag string
var jsonTag bool
var savePath string
var effect int

func init() {
	flag.StringVar(&dsn, "dsn", "root:password@tcp(127.0.0.1:3306)/med_interrogation?charset=utf8", "数据库链接DSN")
	flag.StringVar(&tables, "tables", "file", "表的名称，多个按照逗号隔开")
	flag.StringVar(&pkg, "pkg", "model", "结构体的package")
	flag.StringVar(&tag, "tag", "db", "自定义tag名称")
	flag.BoolVar(&jsonTag, "jsonTag", false, "是否jsonTag")
	flag.StringVar(&savePath, "path", "./model", "保存路径")
	flag.IntVar(&effect, "effect", 0, "模版：0-model，1-ent，3-md, 默认0")
}

func main() {

	flag.Parse()

	err := parse.New(effect).Model(&parse.Param{
		Dsn:      dsn,
		Tables:   strings.Split(tables, ","),
		Tag:      tag,
		JsonTag:  jsonTag,
		SavePath: savePath,
		Pkg:      pkg,
	})

	if nil != err {
		fmt.Println(err)
		log.Println(err)
	}

}

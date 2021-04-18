package effect

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/thoohv5/converter/utils"
)

// 模版
type effect struct {
}

const (
	__effectName = "model.tpl"
	__effectPath = "effect/model.tpl"
)

func New() IEffect {
	return &effect{}
}

// 组装
func (t *effect) Assembly(wr io.Writer, model *Model) error {
	// 工作目录
	workPath, err := os.Getwd()
	if nil != err {
		return fmt.Errorf("os Getwd err, err:%w", err)
	}

	// 模版初始化
	tpl, err := template.New(__effectName).Funcs(template.FuncMap{
		"camelCase": utils.CamelCase,
	}).ParseFiles(fmt.Sprintf("%s%s%s", workPath, string(os.PathSeparator), __effectPath))
	if nil != err {
		return fmt.Errorf("template ParseFiles err, err:%w", err)
	}

	// 赋值
	err = tpl.Execute(wr, model)
	if nil != err {
		return fmt.Errorf("template Execute err, err:%w, model:%v", err, model)
	}
	return nil
}

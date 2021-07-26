package md

import (
	"fmt"
	"io"
	"text/template"

	"github.com/thoohv5/converter/effect"
	"github.com/thoohv5/converter/utils"
)

// 模版
type md struct {
}

const (
	__effectName = "template.tpl"
	__effectPath = "md/template.tpl"
)

func New() effect.IEffect {
	return &md{}
}

// 组装
func (t *md) Assembly(wr io.Writer, model *effect.Model) error {
	// // 工作目录
	// workPath, err := os.Getwd()
	// if nil != err {
	// 	return fmt.Errorf("os Getwd err, err:%w", err)
	// }

	// tpl
	tplByte, err := effect_md_template_tpl()
	if nil != err {
		return fmt.Errorf("effect_model_tpl err, err:%w", err)
	}

	// 模版初始化
	tpl, err := template.New(__effectName).Funcs(template.FuncMap{
		"camelCase": utils.CamelCase,
		// }).ParseFiles(fmt.Sprintf("%s%s%s", workPath, string(os.PathSeparator), __effectPath))
	}).Parse(string(tplByte))
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

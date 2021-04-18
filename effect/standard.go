package effect

import (
	"io"

	"github.com/thoohv5/converter/datasource"
)

// 模版标准化
type IEffect interface {
	// 组装
	Assembly(wr io.Writer, model *Model) error
}

type Model struct {
	// 包
	Package string
	// 表信息
	Table *Table
	// 字段列表
	Columns []*Column
}

type Table struct {
	// 信息
	*datasource.Table
	// 标签
	Tag string
}

type Column struct {
	// 字段
	*datasource.Column
	// 标签
	Tag string
}

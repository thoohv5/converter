package datasource

// 数据标准化
type IDataSource interface {
	// 表信息
	GetTableInfo(tableName string) (*Table, error)
	// 字段列表
	GetColumnList(tableName string) ([]*Column, error)
}

// 字段
type Column struct {
	// 字段名称
	Name string
	// 字段类型
	Type string
	// 是否为空
	IsNullAble string
	// 默认值
	Default *string
	// 字段备注
	Comment string
	// 原始字段类型
	OriginalType string
}

// 表
type Table struct {
	// 表名称
	Name string
	// 备注
	Comment string
}

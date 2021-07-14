package parse

// 转换标准
type IParse interface {
	// 转换
	Transform(columnType string) string
	// 生成Tag
	MakeTag(tagLabel string, columnName string, jsonTag bool) string
	// 生成Model
	Model(param *Param) error
	// 生成Schema
	Schema(param *Param) error
}

type Param struct {
	// dsn
	Dsn string
	// 表列表
	Tables []string
	// 标识
	Tag string
	// jsonTag
	JsonTag bool
	// 保存路径
	SavePath string
	// 包名
	Pkg string
}

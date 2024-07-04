package chain

import (
	"gorm.io/gorm"
	queryContext "table_desc/src/db/context"
	"table_desc/src/entity"
)

type Handler interface {
	SetNext(h Handler)

	Handle(*HandlerParam)
}

type HandlerParam struct {
	// 数据库连接参数
	Param entity.ConnectParam
	// 数据库链接
	Db *gorm.DB
	// word存放的路径
	Path string
	// 下一步查询的表名
	TableName string
	// 操作上下文
	Ctx *queryContext.OpContext
	// 是否允许关闭数据库链接
	Allows bool
}

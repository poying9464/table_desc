package db

import (
	"gorm.io/gorm"
	"table_desc/src/entity"
)

type QueryOperation interface {

	// Connect 连接数据库
	Connect(param entity.ConnectParam) *gorm.DB

	// QueryTableInfo 查询表信息
	QueryTableInfo(scheme string, db *gorm.DB) []map[string]string

	// QueryColumnInfo 查询列信息
	QueryColumnInfo(tableName, scheme string, db *gorm.DB) []map[string]string

	// Close 关闭连接
	Close(db *gorm.DB)
}

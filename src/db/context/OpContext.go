package queryContext

import (
	"gorm.io/gorm"
	"table_desc/src/db"
	"table_desc/src/entity"
)

type OpContext struct {
	queryOperation db.QueryOperation
}

func NewOpContext(queryOperation db.QueryOperation) *OpContext {
	return &OpContext{queryOperation: queryOperation}
}

func (op *OpContext) Connect(param entity.ConnectParam) *gorm.DB {
	return op.queryOperation.Connect(param)
}

func (op *OpContext) QueryTableInfo(scheme string, db *gorm.DB) []map[string]string {
	return op.queryOperation.QueryTableInfo(scheme, db)
}

func (op *OpContext) QueryColumnInfo(tableName, scheme string, db *gorm.DB) []map[string]string {
	return op.queryOperation.QueryColumnInfo(tableName, scheme, db)
}

func (op *OpContext) Close(db *gorm.DB) {
	op.queryOperation.Close(db)
}

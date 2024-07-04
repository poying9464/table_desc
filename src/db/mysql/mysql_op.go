package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"table_desc/src/entity"
)

var dsn = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

var tableInfoSql = `SELECT 
    TABLE_NAME AS tableName,
    TABLE_COMMENT AS comments
FROM 
    INFORMATION_SCHEMA.TABLES
WHERE 
    TABLE_SCHEMA = ?;
`

var colInfoSql = `SELECT 
    COLUMN_NAME AS colName,
    COLUMN_TYPE AS dataType,
    CHARACTER_MAXIMUM_LENGTH AS dataLength,
    COLUMN_COMMENT AS colComments
FROM 
    INFORMATION_SCHEMA.COLUMNS
WHERE 
    TABLE_SCHEMA = ?
    AND TABLE_NAME = ?;
`

type Operation struct {
}

func (do *Operation) Connect(param entity.ConnectParam) *gorm.DB {

	url := fmt.Sprintf(dsn, param.Username, param.Password, param.Host, param.Port, param.Scheme)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return db
	}
	panic(err)
}

func (do *Operation) QueryTableInfo(scheme string, db *gorm.DB) []map[string]string {
	var res []map[string]string
	db.Exec(tableInfoSql, scheme).Scan(&res)
	return res
}

func (do *Operation) QueryColumnInfo(tableName, scheme string, db *gorm.DB) []map[string]string {
	var res []map[string]string
	db.Exec(colInfoSql, scheme, tableName).Scan(&res)
	return res
}

func (do *Operation) Close(db *gorm.DB) {
	if s, err := db.DB(); err == nil {
		err = s.Close()
	}
}

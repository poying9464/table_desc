package dm

import (
	dameng "github.com/godoes/gorm-dameng"
	"gorm.io/gorm"
	"table_desc/src/entity"
)

var queryTableName = `
 SELECT T.TABLE_NAME AS "tableName", C.COMMENTS AS comments
 FROM sys.ALL_TABLES T
 INNER JOIN sys.ALL_TAB_COMMENTS C ON T.TABLE_NAME = C.TABLE_NAME AND C.OWNER = T.OWNER
 WHERE T.OWNER = ?;
`

var queryColumnInfo = `
SELECT C.TABLE_NAME AS "tableName", C.COLUMN_NAME AS "colName", C.DATA_TYPE AS "dataType",
 C.DATA_LENGTH AS "dataLength",U.COMMENTS AS "colComments"
 FROM SYS.ALL_TAB_COLUMNS C INNER JOIN  sys.USER_COL_COMMENTS U
 ON C.TABLE_NAME=U.TABLE_NAME AND C.COLUMN_NAME = U.COLUMN_NAME AND C.OWNER = U.OWNER
 where C.TABLE_NAME = ? AND U.OWNER = ?;
`

type Operation struct {
}

func (do *Operation) Connect(param entity.ConnectParam) *gorm.DB {

	options := map[string]string{
		"schema":         param.Scheme,
		"appName":        "table_desc",
		"connectTimeout": "30000",
	}
	dsn := dameng.BuildUrl(param.Username, param.Password, param.Host, param.Port, options)

	db, err := gorm.Open(dameng.Open(dsn), &gorm.Config{})
	if err != nil {
		return db
	}
	panic(err)
}

func (do *Operation) QueryTableInfo(scheme string, db *gorm.DB) []map[string]string {
	var result []map[string]string
	db.Exec(queryTableName, scheme).Scan(&result)
	return result
}

func (do *Operation) QueryColumnInfo(tableName, scheme string, db *gorm.DB) []map[string]string {
	var result []map[string]string
	db.Exec(queryColumnInfo, tableName, scheme).Scan(&result)
	return result
}

func (do *Operation) Close(db *gorm.DB) {
	if s, err := db.DB(); err == nil {
		err = s.Close()
	}
}

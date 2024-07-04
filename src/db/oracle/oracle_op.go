package oracle

import (
	"gorm.io/gorm"
	"table_desc/src/entity"
)

var tableInfoSql = `SELECT 
    t.table_name AS "tableName",
    c.comments AS "comments"
FROM 
    dba_tables t
LEFT JOIN 
    dba_tab_comments c ON t.owner = c.owner AND t.table_name = c.table_name
ORDER BY 
    t.owner, t.table_name where t.owner = ?;
`

var colInfoSql = `SELECT 
    cols.column_name AS "colName",
    cols.data_type AS "dataType",
    cols.data_length AS "dataLength",
    com.comments AS "colComments"
FROM 
    all_tab_columns cols
LEFT JOIN 
    all_col_comments com ON cols.owner = com.owner 
                         AND cols.table_name = com.table_name 
                         AND cols.column_name = com.column_name
WHERE 
    cols.table_name = UPPER(?) 
    AND cols.owner = ?
ORDER BY 
    cols.column_id;`

type Operation struct {
}

func (do *Operation) Connect(param entity.ConnectParam) *gorm.DB {

	//oralInfo := fmt.Sprintf("%s/%s@%s:%d/%s", param.Username, param.Password, param.Host, param.Port, param.Scheme)
	//db, err := gorm.Open(oracle.Open(oralInfo), &gorm.Config{})
	//if err != nil {
	//	log.Fatal("Error connecting to database:", err)
	//}
	//return db
	return nil
}

func (do *Operation) QueryTableInfo(scheme string, db *gorm.DB) []map[string]string {

	var res []map[string]string
	db.Exec(tableInfoSql, scheme).Scan(&res)
	return res
}

func (do *Operation) QueryColumnInfo(tableName, scheme string, db *gorm.DB) []map[string]string {
	var res []map[string]string
	db.Exec(colInfoSql, tableName, scheme).Scan(&res)
	return res
}

func (do *Operation) Close(db *gorm.DB) {

	if s, err := db.DB(); err != nil {
		err = s.Close()
	}
}

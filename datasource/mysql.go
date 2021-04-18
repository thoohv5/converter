package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	db *sql.DB
}

func New(dsn string) (IDataSource, error) {
	ds := new(mysql)
	db, err := sql.Open("mysql", dsn)
	if nil != err {
		return ds, fmt.Errorf("sql Open err, err:%w", err)
	}
	if err := db.Ping(); err != nil {
		return ds, fmt.Errorf("db Ping err, err:%w", err)
	}
	ds.db = db
	return ds, nil
}

func (d *mysql) GetTableInfo(tableName string) (*Table, error) {
	ret := new(Table)
	if err := d.db.QueryRow("SELECT TABLE_NAME, TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema = DATABASE() AND TABLE_NAME =  ?", tableName).Scan(&ret.Name, &ret.Comment); nil != err {
		return ret, fmt.Errorf("db QueryRow err,err:%w", err)
	}
	return ret, nil
}

func (d *mysql) GetColumnList(tableName string) ([]*Column, error) {
	list := make([]*Column, 0)

	// 查询
	rows, err := d.db.Query("SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,COLUMN_DEFAULT,COLUMN_COMMENT FROM information_schema.COLUMNS WHERE table_schema = DATABASE() AND TABLE_NAME = ? order by TABLE_NAME asc, ORDINAL_POSITION asc;", tableName)
	if nil != err {
		return list, fmt.Errorf("db Query err,err:%w", err)
	}
	defer rows.Close()
	for rows.Next() {
		c := new(Column)
		if err := rows.Scan(&c.Name, &c.Type, &c.isNullAble, &c.Default, &c.Comment); nil != err {
			return list, fmt.Errorf("rows Scan err, err:%w", err)
		}
		list = append(list, c)
	}
	return list, nil
}

package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMySQL(dbUser, dbPassword, dbHost, dbPort, dbName string, maxIdleConns, maxOpenConns int) error {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName))
	if nil != err {
		return err
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	return nil
}

func CloseMySQL() {
	db.Close()
}

// 获取一个事务
func GetTransaction() (*sql.Tx, error) {
	return db.Begin()
}

// 根据SQL语句获取数据库信息
func GetRows(sql string, args ...interface{}) ([]map[string]string, error) {
	stmt, err := db.Prepare(sql)
	if nil != err {
		return nil, err
	}
	defer stmt.Close()

	rst, err := stmt.Query(args...)
	if nil != err {
		return nil, err
	}
	defer rst.Close()

	columns, err := rst.Columns()
	count := len(columns)
	values := make([]string, count)
	scanArgs := make([]interface{}, count)
	for i := 0; i < count; i++ {
		scanArgs[i] = &values[i]
	}
	rows := make([]map[string]string, 0)
	for rst.Next() {
		rst.Scan(scanArgs...)
		entry := make(map[string]string, count)
		for i, col := range columns {
			entry[col] = values[i]
		}
		rows = append(rows, entry)
	}
	return rows, nil
}

// 插入数据库信息，返回插入的ID
func InsertOne(sql string, args ...interface{}) (int64, error) {
	stmt, err := db.Prepare(sql)
	if nil != err {
		return 0, err
	}
	defer stmt.Close()

	rst, err := stmt.Exec(args...)
	if nil != err {
		return 0, err
	}

	return rst.LastInsertId()
}

// 更新数据库信息，返回受影响的条数以及错误标识
func Update(sql string, args ...interface{}) (int64, error) {
	stmt, err := db.Prepare(sql)
	if nil != err {
		return 0, err
	}
	defer stmt.Close()

	rst, err := stmt.Exec(args...)
	if nil != err {
		return 0, err
	}

	return rst.RowsAffected()
}

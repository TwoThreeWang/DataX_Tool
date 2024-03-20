package appScript

import (
	"fmt"
	"strings"
)

// AppGetMysqlDb 查询所有库名
func AppGetMysqlDb(jdbcURL string) string {
	jdbcURL = GetTns(jdbcURL, "information_schema")
	// 连接到 MySQL 数据库
	mgr, err := NewMysqlManager(jdbcURL)
	if err != nil {
		return fmt.Sprintf("Mysql连接报错： %s", err.Error())
	}
	defer mgr.Close()
	// 查询数据
	rows, err := mgr.Query("SHOW DATABASES")
	if err != nil {
		return fmt.Sprintf("Mysql查询所有数据库报错： %s", err.Error())
	}
	defer rows.Close()
	var columns []string
	// 处理查询结果
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return fmt.Sprintf("Mysql查询报错： %s", err.Error())
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return fmt.Sprintf("Mysql查询报错： %s", err.Error())
	}
	return strings.Join(columns, ",")
}

// AppGetMysqlTable 查询所有表名
func AppGetMysqlTable(jdbcURL, database string) string {
	jdbcURL = GetTns(jdbcURL, database)
	// 连接到 MySQL 数据库
	mgr, err := NewMysqlManager(jdbcURL)
	if err != nil {
		return fmt.Sprintf("Mysql连接报错： %s", err.Error())
	}
	defer mgr.Close()
	// 查询数据
	rows, err := mgr.Query(fmt.Sprintf("SELECT TABLE_NAME FROM information_schema.tables WHERE TABLE_SCHEMA = '%s' AND TABLE_TYPE = 'BASE TABLE'", database))
	if err != nil {
		return fmt.Sprintf("Mysql查询所有数据表报错： %s", err.Error())
	}
	defer rows.Close()
	var columns []string
	// 处理查询结果
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return fmt.Sprintf("Mysql查询报错： %s", err.Error())
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return fmt.Sprintf("Mysql查询报错： %s", err.Error())
	}
	fmt.Println(columns)
	return strings.Join(columns, ",")
}

// AppGetMysqlColumn 查询表中所有字段名
func AppGetMysqlColumn(jdbcURL, database, table_name string) string {
	jdbcURL = GetTns(jdbcURL, database)
	// 连接到 MySQL 数据库
	mgr, err := NewMysqlManager(jdbcURL)
	if err != nil {
		return fmt.Sprintf("Mysql连接报错： %s", err.Error())
	}
	defer mgr.Close()
	// 查询数据
	fmt.Println(fmt.Sprintf("SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", database, table_name))
	rows, err := mgr.Query(fmt.Sprintf("SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", database, table_name))
	if err != nil {
		return fmt.Sprintf("Mysql查询所有字段报错： %s", err.Error())
	}
	defer rows.Close()
	var columns []string
	// 处理查询结果
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return fmt.Sprintf("Mysql查询报错： %s", err.Error())
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return fmt.Sprintf("Mysql查询报错： %s", err.Error())
	}
	return strings.Join(columns, ",")
}

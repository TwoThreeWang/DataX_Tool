package appScript

import (
	"encoding/json"
	"fmt"
)

func AppSaveDataSource(setData string) string {
	db := NewSQLiteDB()
	defer db.Close()
	id := MD5(setData)
	var jsonValue map[string]interface{}
	if err := json.Unmarshal([]byte(setData), &jsonValue); err != nil {
		fmt.Println("Error parsing JSON value: %s", err)
		return fmt.Sprintf("数据源保存错误： %s", err.Error())
	}
	insertSQL := "insert into data_source(id, alias,db_type,link,username,password) values (?,?,?,?,?,?)"
	// 插入数据
	err := db.Execute(insertSQL, id, jsonValue["alias"], jsonValue["db_type"], jsonValue["link"], jsonValue["username"], jsonValue["password"])
	if err != nil {
		return fmt.Sprintf("数据源保存错误： %s", err.Error())
	}
	return "数据源保存成功"
}

func AppGetDataSource() string {
	db := NewSQLiteDB()
	defer db.Close()
	// 查询数据
	querySQL := `SELECT id,alias,db_type,link,username,password FROM data_source`
	rows := db.Query(querySQL)
	defer rows.Close()
	result := make([]map[string]string, 0)
	for rows.Next() {
		var alias, db_type, link, id, username, password string
		if err := rows.Scan(&id, &alias, &db_type, &link, &username, &password); err != nil {
			return fmt.Sprintf("数据源获取错误： %s", err.Error())
		}
		// 创建 map
		data := map[string]string{
			"key":      id,
			"alias":    alias,
			"db_type":  db_type,
			"link":     link,
			"username": username,
			"password": password,
		}
		// 追加到 list
		result = append(result, data)
	}
	// 将map转换为JSON格式
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf("数据源获取错误： %s", err.Error())
	}
	return string(jsonData)
}

func AppUpdateDataSource(setData string) string {
	db := NewSQLiteDB()
	defer db.Close()
	var jsonValue map[string]interface{}
	if err := json.Unmarshal([]byte(setData), &jsonValue); err != nil {
		fmt.Println("Error parsing JSON value: %s", err)
		return fmt.Sprintf("数据源修改报错： %s", err.Error())
	}
	SQL := "UPDATE data_source SET alias = ? ,db_type=?, link=?,username=?,password=? WHERE id = ?"
	// 插入数据
	err := db.Execute(SQL, jsonValue["alias"], jsonValue["db_type"], jsonValue["link"], jsonValue["username"], jsonValue["password"], jsonValue["key"])
	if err != nil {
		return fmt.Sprintf("数据源修改报错： %s", err.Error())
	}
	return "数据源修改成功"
}

func AppDeleteDataSource(id string) string {
	db := NewSQLiteDB()
	defer db.Close()
	SQL := "DELETE from data_source WHERE id = ?"
	// 插入数据
	err := db.Execute(SQL, id)
	if err != nil {
		return fmt.Sprintf("数据源删除报错： %s", err.Error())
	}
	return "数据源删除成功"
}

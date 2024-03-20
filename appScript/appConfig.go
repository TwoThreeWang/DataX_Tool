package appScript

import (
	"encoding/json"
	"fmt"
)

func AppGetSetting() string {
	db := NewSQLiteDB()
	defer db.Close()
	// 查询数据
	querySQL := `SELECT config_name,config_value FROM sys_config`
	rows := db.Query(querySQL)
	defer rows.Close()
	result := make(map[string]string)
	for rows.Next() {
		var name, value string
		if err := rows.Scan(&name, &value); err != nil {
			return fmt.Sprintf("配置获取错误： %s", err.Error())
		}
		result[name] = value
	}
	// 将map转换为JSON格式
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf("配置获取错误： %s", err.Error())
	}
	return string(jsonData)
}

func AppSaveSetting(setData string) string {
	db := NewSQLiteDB()
	defer db.Close()
	var jsonValue map[string]interface{}
	if err := json.Unmarshal([]byte(setData), &jsonValue); err != nil {
		fmt.Println("Error parsing JSON value: %s", err)
		return fmt.Sprintf("配置保存错误： %s", err.Error())
	}
	truncateSql := "delete from sys_config"
	err := db.Execute(truncateSql)
	if err != nil {
		return fmt.Sprintf("配置保存错误： %s", err.Error())
	}
	insertSQL := "insert into sys_config(config_name,config_value) values (?,?)"
	// 使用for循环和range遍历map
	for key, value := range jsonValue {
		// 插入数据
		err := db.Execute(insertSQL, key, value)
		if err != nil {
			return fmt.Sprintf("配置保存错误： %s", err.Error())
		}
	}
	return "保存成功"
}

func AppGetDataxPath() string {
	db := NewSQLiteDB()
	defer db.Close()
	// 查询数据
	querySQL := `SELECT config_value FROM sys_config where config_name = 'DataxPath'`
	row := db.QueryRow(querySQL)
	var value string
	if err := row.Scan(&value); err != nil {
		return fmt.Sprintf("配置获取错误： %s", err.Error())
	}
	return value
}

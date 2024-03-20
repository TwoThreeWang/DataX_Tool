package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

func configInit() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetConfig(key string) string {
	configInit()
	// 获取配置参数
	return viper.GetString(key)
}

func SetConfig(key string, value string) (err error) {
	configInit()
	// 检查value是否为特定的JSON结构
	if value[0] == '{' && value[len(value)-1] == '}' {
		var jsonValue map[string]interface{}
		if err = json.Unmarshal([]byte(value), &jsonValue); err != nil {
			fmt.Println("Error parsing JSON value: %s", err)
			return
		}
		// 序列化后的value作为新值
		viper.Set(key, jsonValue)
	} else {
		// 直接将value作为新值
		viper.Set(key, value)
	}
	// 保存配置文件
	err = viper.WriteConfig()
	return
}

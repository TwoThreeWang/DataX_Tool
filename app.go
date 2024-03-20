package main

import (
	"context"
	"dataxTool/appScript"
	"encoding/json"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetSetting 获取设置数据
func (a *App) GetSetting() string {
	return appScript.AppGetSetting()
}

// SaveSetting 保存配置数据
func (a *App) SaveSetting(setData string) string {
	return appScript.AppSaveSetting(setData)
}

// GetDataSource 获取数据源数据
func (a *App) GetDataSource() string {
	return appScript.AppGetDataSource()
}

// SaveDataSource 保存数据源数据
func (a *App) SaveDataSource(setData string) string {
	return appScript.AppSaveDataSource(setData)
}

// UpdateDataSource 更新数据源
func (a *App) UpdateDataSource(setData string) string {
	return appScript.AppUpdateDataSource(setData)
}

// DeleteDataSource 删除数据源
func (a *App) DeleteDataSource(id string) string {
	return appScript.AppDeleteDataSource(id)
}

// GetTxtFirstLine 获取文本文件第一行内容
func (a *App) GetTxtFirstLine(filepath, encoding string) string {
	return appScript.AppReadFirstLine(filepath, encoding)
}

// GetJobJson 获取拼接后的job配置文件
func (a *App) GetJobJson(reader, readerType, writer, writerType, jsonPath string) string {
	return appScript.AppGetJobJson(reader, readerType, writer, writerType, jsonPath)
}

// GetJsonFileList 获取配置文件列表
func (a *App) GetJsonFileList() []string {
	// 获取datax路径
	filePath := appScript.AppGetDataxPath()
	res, err := appScript.AppGetFileList(filePath + "/job")
	if err != nil {
		return []string{"文件列表获取错误！"}
	}
	return res
}

// DelJsonFileList 删除配置文件
func (a *App) DelJsonFile(file string) string {
	// 获取datax路径
	filePath := appScript.AppGetDataxPath()
	res := appScript.AppDelFile(filePath + "/job/" + file)
	return res
}

// ReadFile 获取文件内容
func (a *App) ReadFile(file string) string {
	// 获取datax路径
	filePath := appScript.AppGetDataxPath()
	res, err := appScript.AppReadFile(filePath + "/job/" + file)
	if err != nil {
		return fmt.Sprintf("获取文件内容报错： %s", err.Error())
	}
	return res
}

// SaveFile 保存文件
func (a *App) SaveFile(file, content string) string {
	// 获取datax路径
	filePath := appScript.AppGetDataxPath()
	err := appScript.AppSaveFile(filePath+"/job/"+file, content)
	if err != nil {
		return fmt.Sprintf("保存文件内容报错： %s", err.Error())
	}
	return "保存成功"
}

// RunScript 运行程序
func (a *App) RunScript(file string) string {
	// 获取datax路径
	filePath := appScript.AppGetDataxPath()
	runCode := fmt.Sprintf("%s/bin/datax.py", filePath)
	fmt.Println(runCode)
	res := map[string]interface{}{
		"pid":  0,
		"info": "",
	}
	// 定义参数
	args := []string{fmt.Sprintf("%s/job/%s", filePath, file)}
	pid, err := appScript.AppStartScript(runCode, filePath+"/job/log/out.log", args)
	res["pid"] = pid
	if err != nil {
		res["info"] = fmt.Sprintf("运行程序报错： %s", err.Error())
	} else {
		res["info"] = "开始运行成功"
	}
	// 将map转换为JSON格式
	jsonData, err := json.Marshal(res)
	if err != nil {
		return fmt.Sprintf("程序运行结果信息格式化报错： %s", err.Error())
	}
	return string(jsonData)
}

// KillScript 杀掉进程
func (a *App) KillScript(pid int) string {
	err := appScript.AppKillProcess(pid)
	if err != nil {
		return fmt.Sprintf("结束程序报错： %s", err.Error())
	}
	return "结束程序成功"
}

// IsProcessRunning 判断进程是否存在
func (a *App) IsProcessRunning(pid int) bool {
	return appScript.AppIsProcessRunning(pid)
}

// GetMysqlDb 查询所有库名
func (a *App) GetMysqlDb(jdbcURL string) string {
	return appScript.AppGetMysqlDb(jdbcURL)
}

// GetMysqlTable 查询所有表名
func (a *App) GetMysqlTable(jdbcURL, database string) string {
	return appScript.AppGetMysqlTable(jdbcURL, database)
}

// GetMysqlColumn 查询所有字段名
func (a *App) GetMysqlColumn(jdbcURL, database, table_name string) string {
	return appScript.AppGetMysqlColumn(jdbcURL, database, table_name)
}

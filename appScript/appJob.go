package appScript

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AppGetJobJson 拼接 job 配置文件
func AppGetJobJson(reader, readerType, writer, writerType, jsonFile string) string {
	fmt.Println(writer)
	//var readerData map[string]interface{}
	readerData := make(map[string]interface{})
	if readerType == "txt" {
		var readerJsonValue map[string]string
		if err := json.Unmarshal([]byte(reader), &readerJsonValue); err != nil {
			fmt.Println("Error parsing readerJsonValue value: ", err.Error())
			return fmt.Sprintf("reader信息解析报错： %s", err.Error())
		}
		path := strings.Replace(readerJsonValue["path"], "\\", "\\\\", -1)
		paths := strings.Split(path, ",")
		readerData = GetTxtReader(paths, readerJsonValue["encoding"], readerJsonValue["fieldDelimiter"], readerJsonValue["skipHeader"], readerJsonValue["columnNum"])
	} else if readerType == "mysql" {
		var readerDataParameter map[string]interface{}
		if err := json.Unmarshal([]byte(reader), &readerDataParameter); err != nil {
			fmt.Println("Error parsing readerJsonValue value: ", err.Error())
			return fmt.Sprintf("reader信息解析报错： %s", err.Error())
		}
		readerData["parameter"] = readerDataParameter
		readerData["name"] = "mysqlreader"
	} else {
		fmt.Println(reader)
		return "未支持的 readerType"
	}

	writerData := make(map[string]interface{})
	if writerType == "txt" {
		var writerJsonValue map[string]string
		if err := json.Unmarshal([]byte(writer), &writerJsonValue); err != nil {
			fmt.Println("Error parsing writerJsonValue value: %s", err)
			return fmt.Sprintf("writer信息解析报错： %s", err.Error())
		}
		path := strings.Replace(writerJsonValue["path"], "\\", "\\\\", -1)
		writerData = GetTxtWriter(path, writerJsonValue["fileName"], writerJsonValue["writeMode"], writerJsonValue["fileFormat"], writerJsonValue["encoding"], writerJsonValue["fieldDelimiter"])
	} else if writerType == "mysql" {
		var writerDataParameter map[string]interface{}
		if err := json.Unmarshal([]byte(writer), &writerDataParameter); err != nil {
			fmt.Println("Error parsing writerJsonValue value: ", err.Error())
			return fmt.Sprintf("writer信息解析报错： %s", err.Error())
		}
		writerData["parameter"] = writerDataParameter
		writerData["name"] = "mysqlwriter"
	} else {
		return "未支持的 writerType"
	}
	jobConfigData := JobConfigData{}
	jobConfigData.Job.Setting.Speed.Channel = 2
	jobConfigData.Job.Setting.ErrorLimit.Record = 2
	jobConfigData.Job.Setting.ErrorLimit.Percentage = 0.02
	// 如果Content没有元素，添加一个新元素
	newContent := struct {
		Reader map[string]interface{} `json:"reader"`
		Writer map[string]interface{} `json:"writer"`
	}{}
	jobConfigData.Job.Content = append(jobConfigData.Job.Content, newContent)
	jobConfigData.Job.Content[0].Reader = readerData
	jobConfigData.Job.Content[0].Writer = writerData
	jsonBytes, err := json.Marshal(jobConfigData)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return fmt.Sprintf("结果json解析报错： %s", err.Error())
	}
	if jsonFile != "" {
		// 获取datax路径
		filePath := AppGetDataxPath()
		if err = AppSaveFile(filePath+"/job/"+jsonFile+".json", string(jsonBytes)); err != nil {
			return fmt.Sprintf("保存结果json到文件报错： %s", err.Error())
		}
	}
	return string(jsonBytes)
}

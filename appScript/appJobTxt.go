package appScript

import (
	"fmt"
	"strconv"
)

// GetTxtReader 获取 txtfilereader
func GetTxtReader(path []string, encoding, fieldDelimiter, skipHeader, columnNum string) (readerRes map[string]interface{}) {
	reader := &TxtReader{
		Name: "txtfilereader",
	}
	reader.Parameter.Path = path
	reader.Parameter.Encoding = encoding
	reader.Parameter.FieldDelimiter = fieldDelimiter
	reader.Parameter.SkipHeader = skipHeader
	// 初始化一个空的Column类型的切片
	n, err := strconv.Atoi(columnNum)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	list := make([]TxtColumn, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, TxtColumn{
			Type:  "string",
			Index: i,
		})
	}
	reader.Parameter.Column = list
	readerRes = ToStruct(reader)
	return
}

// GetTxtWriter 获取 txtfilewriter
func GetTxtWriter(path, fileName, writeMode, fileFormat, encoding, fieldDelimiter string) (writerRes map[string]interface{}) {
	writer := TxtWriter{
		Name: "txtfilewriter",
	}
	writer.Parameter.Path = path
	writer.Parameter.FileName = fileName
	writer.Parameter.WriteMode = writeMode
	writer.Parameter.FileFormat = fileFormat
	writer.Parameter.Encoding = encoding
	writer.Parameter.FieldDelimiter = fieldDelimiter
	writerRes = ToStruct(writer)
	return
}

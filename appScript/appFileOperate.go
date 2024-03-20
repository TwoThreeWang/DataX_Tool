package appScript

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// AppGetFileList 获取目录下所有文件
func AppGetFileList(dirPath string) (filenames []string, err error) {
	// 读取目录中的所有文件和目录信息
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 遍历文件信息
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}
	return
}

// AppSaveFile 将内容保存到文件中
func AppSaveFile(filePath, content string) error {
	// 将内容写入文件
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	return err
}

// AppReadFile 读取文件内容
func AppReadFile(filePath string) (content string, err error) {
	con, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	content = string(con)
	return
}

// AppDelFile 删除文件
func AppDelFile(filePath string) string {
	// 删除文件
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("文件删除报错： %s", err.Error())
	}
	return fmt.Sprintf("文件已删除")
}

// AppReadFirstLine 读取指定文件的第一行，并根据指定的编码返回内容
func AppReadFirstLine(filePath string, encoding string) string {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Sprintf("获取文件第一行错误： %s", err.Error())
	}
	defer file.Close()
	// 根据编码类型创建解码器
	var reader io.Reader
	switch encoding {
	case "GBK":
		// 如果是GBK编码，使用transform.NewReader进行转换
		gbkReader := transform.NewReader(file, simplifiedchinese.GBK.NewDecoder())
		reader = gbkReader
	default:
		reader = file
	}
	// 创建bufio.Reader来逐行读取
	bufioReader := bufio.NewReader(reader)
	line, err := bufioReader.ReadString('\n')
	if err != nil && err != io.EOF {
		return fmt.Sprintf("获取文件第一行错误： %s", err.Error())
	}
	// 去除行尾的换行符
	line = strings.TrimSuffix(line, "\n")
	return line
}

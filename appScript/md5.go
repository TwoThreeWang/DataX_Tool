package appScript

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(value string) string {
	// 创建新的 MD5 哈希计算器
	h := md5.New()
	// 写入要计算的字符串，注意这里需要转换为字节切片
	io.WriteString(h, value)
	// 计算MD5校验和，Sum函数返回的是字节切片
	md5sum := h.Sum(nil)
	// 将字节切片转换为16进制字符串
	return fmt.Sprintf("%x", md5sum)
}

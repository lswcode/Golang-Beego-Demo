package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(pwd string) string {
	data := []byte(pwd)                       // md5.Sumd的参数格式是字节切片，所以要先将原数据转换为字节切片
	fmt.Println(md5.Sum(data))                // 加密后返回一个长度16的byte字节切片
	return fmt.Sprintf("%x\n", md5.Sum(data)) // 将字节切片转换为16进制的字符串
	// 使用fmt.Sprintf+ %x 将字节转换为16进制的字符串
}

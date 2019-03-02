package util

import (
	"crypto/md5"
	"fmt"
)

//GetMD5 计算字符串的MD5值并返回MD5值得字符串
func GetMD5(data string) string {
	ret := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", ret[:])
}

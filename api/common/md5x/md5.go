package md5x

import (
	"crypto/md5"
	"fmt"
)

// 生成MD5
func Md5(s string) string {
	m := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", m)
}

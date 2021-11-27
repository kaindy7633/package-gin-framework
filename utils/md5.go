package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 编写 MD5() 用于 token 编码
func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

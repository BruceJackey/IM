package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写的md5加密
func MD5Encode(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	tempStr := hash.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写的md5加密
func MD5UpperEncode(str string) string {
	return strings.ToUpper(MD5Encode(str))
}

// 生成密码
func MakePassword(password,salt string) string {
	return MD5Encode(password + salt)
}

// 验证密码
func VerifyPassword(password,salt string,hashedPassword string) bool {
	return MakePassword(password,salt) == hashedPassword
}
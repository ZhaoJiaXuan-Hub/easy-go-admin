package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// EncryptionMd5 字符串加密ß
func EncryptionMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// EncryptionPassword 密码加密
func EncryptionPassword(password string, salt string) string {
	password = EncryptionMd5(password)
	salt = EncryptionMd5(salt)
	return EncryptionMd5(password + salt)
}

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyz0123456789"
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}

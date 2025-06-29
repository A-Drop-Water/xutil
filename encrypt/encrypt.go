package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// GetMD5 返回字符串的MD5值
func GetMD5(code string) string {
	h := md5.New()
	_, _ = h.Write([]byte(code))
	return hex.EncodeToString(h.Sum(nil))
}

func GetBcrypt(code string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	return string(hash)
}

func CheckBcrypt(code, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(code))
	return err == nil
}
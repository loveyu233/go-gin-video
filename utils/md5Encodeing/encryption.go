package md5Encodeing

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/viper"
)

// 密码MD5加密
func PasswordMd5Encryption(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	hash.Write([]byte(viper.GetString("encryption.md5secretkey")))
	return hex.EncodeToString(hash.Sum(nil))
}

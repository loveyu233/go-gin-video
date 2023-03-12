package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

// 生成jwt所需的属性
type claims struct {
	jwt.RegisteredClaims
	UserId int32
	Role   int32
}

// jwt加密秘钥
var secretKey = []byte(viper.GetString("encryption.tokenSecretKey"))

// 创建token
func CreateToken(userid, role int32) (string, error) {
	if userid == 0 {
		return "", errors.New("用户名不存在")
	}
	expiresTime := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-gin-video",
			Subject:   "encryption",
			ExpiresAt: jwt.NewNumericDate(expiresTime),
		},
		UserId: userid,
		Role:   role,
	})
	signedString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// token解析
func ParseToken(tokenString string) (int32, int32, error) {
	claim := &claims{}
	_, _, err := jwt.NewParser().ParseUnverified(tokenString, claim)
	if err != nil {
		return -1, -1, errors.New("token解析失败")
	}
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return -1, -1, err
	}
	return claim.UserId, claim.Role, nil
}

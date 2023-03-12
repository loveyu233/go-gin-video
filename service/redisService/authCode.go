package redisService

import (
	"context"
	"errors"
	"fmt"
	"go-gin-video/utils/email"
	"time"
)

var emailCodePrefix = "emailCode"
var captchaCodePrefix = "captchaCode"

func SaveEmailCode(emailStr string) error {
	key := fmt.Sprintf("%s:%s", emailCodePrefix, emailStr)
	if redisClient.Get(context.TODO(), key).Val() != "" {
		return errors.New("不要重复发送验证码")
	}
	code := email.RandCode()
	if !email.EmailSendMsg(emailStr, code) {
		return errors.New("验证码发送失败")
	}
	return redisClient.Set(context.TODO(), key, code, time.Second*120).Err()
}

func VerificationEmailCode(email string, code string) bool {
	key := fmt.Sprintf("%s:%s", emailCodePrefix, email)
	val := redisClient.Get(context.TODO(), key).Val()
	if val == code {
		redisClient.Del(context.TODO(), key)
		return true
	}
	return false
}

func SaveCaptchaCode(email string, x int32) {
	key := fmt.Sprintf("%s:%s", captchaCodePrefix, email)
	redisClient.Set(context.TODO(), key, x, time.Second*60)
}

func VerificationCaptchaCode(email string, x int) bool {
	key := fmt.Sprintf("%s:%s", captchaCodePrefix, email)
	xValue, _ := redisClient.Get(context.TODO(), key).Int()
	if xValue == x {
		return true
	}
	return false
}

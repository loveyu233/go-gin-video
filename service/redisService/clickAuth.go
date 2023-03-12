package redisService

import (
	"context"
	"time"
)

const clickPrefix = "clickPrefix"

// 判断用户是否重复点击,返回值false为没有重复点击
func IsFrequentClicks(userId string) bool {
	key := clickPrefix + ":" + userId
	val := redisClient.Get(context.TODO(), key).Val()
	if val == "" {
		// 第一次点击就把用户id存储,下次点击先验证是否有该值有则表示为重复点击
		redisClient.Set(context.TODO(), key, "1", time.Second*30)
		return false
	}
	return true
}

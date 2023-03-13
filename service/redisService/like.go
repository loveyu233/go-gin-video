package redisService

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

const videoPrefix = "videoLike:"

func IsLike(vid, uid string) bool {
	if redisClient.ZScore(context.TODO(), videoPrefix+vid, uid).Val() != 0 {
		return true
	}
	return false
}

func AddLike(vid, uid string) {
	redisClient.ZAdd(context.TODO(), videoPrefix+vid, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: uid,
	})
}

func DelLike(vid, uid string) {
	redisClient.ZRem(context.TODO(), videoPrefix+vid, uid)
}

func LikeCount(vid string) int64 {
	return redisClient.ZCount(context.TODO(), videoPrefix+vid, "0", strconv.FormatInt(time.Now().Unix(), 10)).Val()
}

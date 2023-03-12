package redisService

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var redisClient *redis.Client

func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Username: viper.GetString("redis.username"),
		Password: viper.GetString("redis.password"),
	})
	err := redisClient.Ping(context.TODO()).Err()
	if err != nil {
		return err
	}
	return nil
}

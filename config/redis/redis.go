package redis

import (
	"context"
	"easy-go-admin/config"
	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
)

// InitRedis 初始化redis
func InitRedis() error {
	var ctx = context.Background()
	var dbConfig = config.Config.Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     dbConfig.Address,
		Password: dbConfig.Password,
		DB:       dbConfig.Db,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

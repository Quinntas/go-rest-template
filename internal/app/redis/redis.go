package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisClient *redis.Client

func ConnectRedis(redisUrl string) {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(opt)
}

func Set[T string](key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return redisClient.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	ctx := context.Background()
	return redisClient.Get(ctx, key).Result()
}

package util

import (
	"context"
	"redu/global"
	"time"
)

func RedisStrGet(key string) (val string, err error) {
	val, err = global.RedisClient.Get(context.Background(), key).Result()
	return
}

func RedisStrSet(key, val string, ex time.Duration) (err error) {
	return global.RedisClient.Set(context.Background(), key, val, ex).Err()
}

func RedisLPush(key string, values []string) (err error) {
	return global.RedisClient.LPush(context.Background(), key, values).Err()
}

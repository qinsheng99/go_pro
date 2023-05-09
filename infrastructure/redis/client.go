package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var red *redis.Client

func Init(cfg *Config) (err error) {
	red = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:       0,
		Password: "",
	})
	_, err = red.Ping(context.Background()).Result()
	if err != nil {
		return
	}
	return nil
}

func GetRedis() *redis.Client {
	return red
}

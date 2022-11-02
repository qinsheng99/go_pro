package app

import (
	"context"
	"github.com/qinsheng99/go-domain-web/domain/redis"
	"time"
)

type redisService struct {
	r redis.Redis
}

type RedisService interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
}

func NewRedisService(r redis.Redis) RedisService {
	return &redisService{
		r: r,
	}
}

func (r *redisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.r.Set(ctx, key, value, expiration)
}

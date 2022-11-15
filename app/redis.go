package app

import (
	"context"
	v8 "github.com/go-redis/redis/v8"
	"github.com/qinsheng99/go-domain-web/domain/redis"
	"time"
)

type redisService struct {
	r redis.Redis
}

type RedisServiceImpl interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	Zadd(ctx context.Context, key string, m ...*v8.Z) (int64, error)
	ZRevrange(ctx context.Context, key string, start, stop int64) ([]string, error)
	Del(ctx context.Context, key string) (int64, error)
}

func NewRedisService(r redis.Redis) RedisServiceImpl {
	return &redisService{r: r}
}

func (r *redisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.r.Set(ctx, key, value, expiration)
}

func (r *redisService) Zadd(ctx context.Context, key string, m ...*v8.Z) (int64, error) {
	return r.r.Zadd(ctx, key, m...)
}

func (r *redisService) ZRevrange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.r.ZRevrange(ctx, key, start, stop)
}

func (r *redisService) Del(ctx context.Context, key string) (int64, error) {
	return r.r.Del(ctx, key)
}

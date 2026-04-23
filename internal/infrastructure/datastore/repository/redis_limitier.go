package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisLimiter struct {
	client *redis.Client
	limit  int
	window time.Duration
}

func NewRedisLimiter(client *redis.Client, limit int, window time.Duration) *RedisLimiter {
	return &RedisLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

func (r *RedisLimiter) Allow(ctx context.Context, key string) (bool, error) {
	redisKey := fmt.Sprintf("rl:%s", key)

	count, err := r.client.Incr(ctx, redisKey).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		r.client.Expire(ctx, redisKey, r.window)
	}

	if count > int64(r.limit) {
		return false, nil
	}

	return true, nil
}

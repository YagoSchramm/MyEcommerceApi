package foundation

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewClient(addr string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		DB:           1,
		PoolSize:     10,
		MinIdleConns: 2,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return rdb, nil
}

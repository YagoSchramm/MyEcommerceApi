package foundation

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewClient(addr string) (*redis.Client, error) {
	var (
		rdb *redis.Client
		err error
	)

	if strings.Contains(addr, "://") {
		var opt *redis.Options
		opt, err = redis.ParseURL(addr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse redis url: %w", err)
		}

		parsed, parseErr := url.Parse(addr)
		if parseErr != nil {
			return nil, fmt.Errorf("failed to inspect redis url: %w", parseErr)
		}
		if strings.Trim(parsed.Path, "/") == "" {
			opt.DB = 1
		}
		opt.PoolSize = 10
		opt.MinIdleConns = 2

		rdb = redis.NewClient(opt)
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:         addr,
			DB:           1,
			PoolSize:     10,
			MinIdleConns: 2,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return rdb, nil
}

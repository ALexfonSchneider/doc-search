package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Repository struct {
	client          *redis.Client
	cacheExpireTime time.Duration
}

func NewRepository(ctx context.Context, addr string, password string, cacheExpireTime time.Duration) (*Repository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	status := client.Ping(ctx)
	err := status.Err()
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		client:          client,
		cacheExpireTime: cacheExpireTime,
	}

	return repo, nil
}

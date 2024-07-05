package net

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	CheckConnection(ctx context.Context) (string, error)
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, val string, exp time.Duration) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(rClient *redis.Client) *RedisCache {
	return &RedisCache{
		client: rClient,
	}
}

func (r *RedisCache) CheckConnection(ctx context.Context) (string, error) {
	return r.client.Ping(ctx).Result()
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, err
}

func (r *RedisCache) Set(ctx context.Context, key string, val string, exp time.Duration) error {
	if err := r.client.Set(ctx, key, val, exp).Err(); err != nil {
		return err
	}
	return nil
}

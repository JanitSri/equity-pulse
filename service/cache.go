package service

import (
	"context"
	"strings"
	"time"

	"github.com/JanitSri/equity-pulse/net"
	"github.com/JanitSri/equity-pulse/util"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

type CacheService struct {
	cache net.Cache
}

func NewCacheService(c net.Cache) *CacheService {
	return &CacheService{
		cache: c,
	}
}

func (c *CacheService) Get(key string, target interface{}) error {
	zap.L().Sugar().Infof("checking cache for key %s", key)

	val, err := c.cache.Get(ctx, key)
	if err != nil {
		if err == redis.Nil {
			zap.L().Sugar().Infof("cache miss for key %s", key)
		} else {
			zap.L().Sugar().Errorf("value cannot be retrieved for key %s: %s", key, err)
		}

		return err
	}

	if err := util.DecodeJSONContentType(strings.NewReader(val), target); err != nil {
		zap.L().Sugar().Errorf("cannot decode val %s: %s", val, err)
		return err
	}

	return nil
}

func (c *CacheService) Set(key string, val interface{}, exp time.Duration) error {
	zap.L().Sugar().Infof("setting key %s value %s", key, val)

	sVal, err := util.ToJSONString(val)
	if err != nil {
		zap.L().Sugar().Errorf("cannot marshal value %s to string: %s", val, err)
		return err
	}

	if err := c.cache.Set(ctx, key, sVal, exp); err != nil {
		zap.L().Sugar().Errorf("cannot set key %s with value %s: %s", key, val, err)
		return err
	}

	return nil
}

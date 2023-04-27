package cache

import (
	"context"
	"errors"
	redisCache "github.com/coherentopensource/go-service-framework/cache"
	"github.com/coherentopensource/go-service-framework/util"
	"strconv"
)

type cache struct {
	redis *redisCache.Cache
}

type Cache interface {
	GetCurrentBlockNumber(ctx context.Context, blockChainInfoKey string) (uint64, error)
	SetCurrentBlockNumber(ctx context.Context, blockChainInfoKey string, blockNumber uint64) error
}

func NewCache(cfg *Config, logger util.Logger) *cache {
	client := redisCache.NewRedisClient(
		&redisCache.RedisConfig{
			Host:     cfg.Host,
			Username: cfg.Username,
			Password: cfg.Password,
			DB:       cfg.DB,
		},
		logger,
	)

	return &cache{
		redis: client,
	}
}

func (c *cache) GetCurrentBlockNumber(ctx context.Context, blockChainInfoKey string) (uint64, error) {
	val, err := c.redis.Get(ctx, blockChainInfoKey)
	if err != nil {
		return 0, err
	}

	blockNumber, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		return 0, errors.New("could not convert blocknumber val from redis to uint64")
	}
	return blockNumber, nil

}

func (c *cache) SetCurrentBlockNumber(ctx context.Context, blockChainInfoKey string, blockNumber uint64) error {
	err := c.redis.Set(ctx, blockChainInfoKey, blockNumber)

	if err != nil {
		return err
	}

	return nil
}

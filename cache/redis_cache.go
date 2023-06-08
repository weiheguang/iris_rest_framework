package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisCache(host string, password string, db int) (ICache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	c := CacheRedis{rdb: rdb}
	_, err := c.Ping()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type CacheRedis struct {
	rdb *redis.Client // redis client
}

func (c *CacheRedis) Ping() (string, error) {
	ctx := context.Background()
	res, err := c.rdb.Ping(ctx).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

// 设置缓存
func (c *CacheRedis) Set(key string, val interface{}, expiration time.Duration) error {
	ctx := context.Background()
	err := c.rdb.Set(ctx, key, val, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// 读取缓存
func (c *CacheRedis) Get(key string) (string, error) {
	ctx := context.Background()
	return c.rdb.Get(ctx, key).Result()
}

func (c *CacheRedis) GetBytes(key string) ([]byte, error) {
	ctx := context.Background()
	return c.rdb.Get(ctx, key).Bytes()
}

func (c *CacheRedis) Close() error {
	if err := c.rdb.Close(); err != nil {
		return err
	}
	return nil
}

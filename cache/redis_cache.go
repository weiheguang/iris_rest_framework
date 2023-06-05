package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisCache(host string, password string, db int) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	c := CacheRedis{rdb: rdb}
	_, err := c.Ping()
	if err != nil {
		return err
	}
	// myLogger.Info("初始化 redis 成功: ", host)
	// return &
	cache = &c
	return nil
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
		// myLogger.Error("设置 cache 出错: ", err)
		// fmt.Println("设置 cache 出错: %v \n", err)
		return err
	}
	return nil
}

// 读取缓存
func (c *CacheRedis) Get(key string) (string, error) {
	ctx := context.Background()
	return c.rdb.Get(ctx, key).Result()
	// if err != nil {
	// 	myLogger.Error("获取 cache key 出错: ", err)
	// 	return "", err
	// }
	// return val, nil
}

func (c *CacheRedis) GetBytes(key string) ([]byte, error) {
	ctx := context.Background()
	return c.rdb.Get(ctx, key).Bytes()
	// if err != nil {
	// 	myLogger.Error("获取 cache key 出错: ", err)
	// 	return "", err
	// }
	// return val, nil
}

func (c *CacheRedis) Close() error {
	if err := c.rdb.Close(); err != nil {
		return err
	}
	return nil
}

// 生成唯一key值
// func (c *CacheRedis) GetKey() string {
// 	return "ccl.order.clost_order"
// }

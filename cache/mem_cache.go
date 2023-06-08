package cache

import (
	"time"

	mem "github.com/patrickmn/go-cache"
)

type CacheMem struct {
	// c *redis.Client // redis client
	c *mem.Cache
	// cache.New(5*time.Minute, 10*time.Minute)

}

func NewMemCache() ICache {
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     host,
	// 	Password: password, // no password set
	// 	DB:       db,       // use default DB
	// })
	mem := mem.New(5*time.Minute, 10*time.Minute)
	c := CacheMem{c: mem}
	// _, err := c.Ping()
	// if err != nil {
	// 	return err
	// }
	return &c
}

func (s *CacheMem) Ping() (string, error) {
	// ctx := context.Background()
	// res, err := self.c.Ping(ctx).Result()
	// if err != nil {
	// 	return "", err
	// }
	return "ok", nil
}

// 设置缓存
func (s *CacheMem) Set(key string, val interface{}, expiration time.Duration) error {
	// ctx := context.Background()
	s.c.Set(key, val, expiration)
	// if err != nil {
	// 	myLogger.Error("设置 cache 出错: ", err)
	// 	return err
	// }
	return nil
}

// 读取缓存
func (s *CacheMem) Get(key string) (string, error) {
	// ctx := context.Background()
	// result , err := self.c.Get(key)

	var res string
	if x, found := s.c.Get(key); found {
		res = x.(string)
		return res, nil
	}
	return res, nil

	// if err != nil {
	// 	myLogger.Error("获取 cache key 出错: ", err)
	// 	return "", err
	// }
	// return val, nil
}

func (s *CacheMem) GetBytes(key string) ([]byte, error) {
	// ctx := context.Background()
	// return c.rdb.Get(ctx, key).Bytes()
	// if err != nil {
	// 	myLogger.Error("获取 cache key 出错: ", err)
	// 	return "", err
	// }
	// return val, nil
	return nil, nil
}

func (s *CacheMem) Close() error {
	// if err := c.rdb.Close(); err != nil {
	// 	return err
	// }
	return nil
}

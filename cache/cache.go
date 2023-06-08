package cache

import (
	"time"
)

var cache ICache

const (
	CacheTypeMem   = "mem"
	CacheTypeRedis = "redis"
)

// 抽象 cache 类
type ICache interface {
	Set(key string, val interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	GetBytes(key string) ([]byte, error) // 获取缓存的 bytes
	// GetKey() string                      // 生成key值
	Ping() (string, error)
	Close() error // 关闭连接
}

// 全局唯一
func GetCache() ICache {
	if cache == nil {
		panic("cache 没有初始化")
	}
	return cache
}

func InitRedis(host string, password string, db int) {
	ca, err := NewRedisCache(host, password, db)
	if err != nil {
		panic(err)
	}
	_, err = ca.Ping()
	if err != nil {
		panic(err)
	}
	cache = ca
}

func InitMem() {
	ca := NewMemCache()
	cache = ca
}

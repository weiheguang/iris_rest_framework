package cache

import (
	"time"
)

var cache ICache

// var myLogger = logging.GetLogger()

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
	return cache
}

package settings

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Settings struct {
	*viper.Viper
}

// var myLogger = logging.GetLogger()

var s *Settings

// 导出函数, 如果需要使用其他函数. 参考: https://github.com/spf13/viper/blob/master/viper.go
// var (
// 	GetString   = viper.GetString   // 获取字符串
// 	GetBool     = viper.GetBool     // 获取 bool
// 	GetDuration = viper.GetDuration // 获取时间长度
// 	GetInt      = viper.GetInt      // 获取整形
// 	GetInt64    = viper.GetInt64    // 获取64位整形
// )

// 快捷获取字符串
func GetString(key string) string {
	return s.GetString(key)
}

// 快捷获取 bool
func GetBool(key string) bool {
	return s.GetBool(key)
}

// 快捷获取时间长度
func GetDuration(key string) time.Duration {
	return s.GetDuration(key)
}

// 快捷获取整形
func GetInt(key string) int {
	return s.GetInt(key)
}

// 快捷获取64位整形
func GetInt64(key string) int64 {
	return s.GetInt64(key)
}

func Init(fileName string) {
	if fileName != "" {
		viper.SetConfigName(fileName) // name of config file (without extension)
		viper.AddConfigPath(".")
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			// panic(fmt.Errorf("配置文件不存在: %s", err))
			fmt.Printf("配置文件不存在: %s", err)
			// myLogger.Error("当前文件夹没有找到 配置文件")
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	s = &Settings{viper.GetViper()}
}

// 获取 Settings
func GetSettings() *Settings {
	return s
}

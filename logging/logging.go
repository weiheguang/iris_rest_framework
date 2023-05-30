package logging

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	// logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetFormatter(&logrus.TextFormatter{
		// DisableColors:   true,
		TimestampFormat: "2006-01-02T15:04:05",
		FullTimestamp:   true,
		// ForceFormatting: true,
	})
	// 设置日志函数名字输出
	logger.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logrus.DebugLevel)
	// logger.SetReportCaller(true)

}

func SetLogLevel(level logrus.Level) {
	if logger == nil {
		fmt.Printf("日志没有初始化")
	}
	logger.SetLevel(level)
}

// GetLogger return a logger
func GetLogger() *logrus.Logger {
	return logger
}

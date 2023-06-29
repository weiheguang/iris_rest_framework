package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *IRFLogger

type IRFLogger = logrus.Logger

func NewLogger(debug bool) *IRFLogger {
	// var logger *Logger
	logger := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	// logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetFormatter(&logrus.TextFormatter{
		// DisableColors:   true,
		TimestampFormat: "2006-01-02T15:04:05",
		FullTimestamp:   true,
		// ForceFormatting: true,
	})
	// 设置日志函数名字输出
	// logger.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	if debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}

func Init(debug bool) {
	logger = NewLogger(true)
}

func GetLogger() *IRFLogger {
	if logger == nil {
		logger = NewLogger(true)
	}
	return logger
}

package iris_rest_framework

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/weiheguang/iris_rest_framework/settings"
	// "gorm.io/gorm/logger"
)

/*
	1. 创建 iris app
	2. 设置日志级别
	3. 设置 recover 中间件
	4. 设置 logger 中间件
*/
func NewIrisApp() *iris.Application {
	settings.Init("test_settings")
	app := iris.New()
	app.Logger().SetLevel("info")
	app.Use(recover.New())
	// app.Use(logger.New())
	return app
}

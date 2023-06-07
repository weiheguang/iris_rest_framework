package iris_rest_framework

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/weiheguang/iris_rest_framework/cache"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/settings"
	// "gorm.io/gorm/logger"
)

type IrisAppConfig struct {
	// setting 名字, 默认值 settings.yaml
	SettingsName string
	// 默认值: notFound
	NotFoundHandler iris.Handler
	// 默认值: internalServerError
	InternalServerErrorHandler iris.Handler
	// 是否启用swagger, 默认值: false
	EnableSwagger bool
	// 是否启用redis, 默认值: false
	EnableRedis bool
	// 是否启用gorm日志, 默认值: false
	EnableGormLog bool
	// 是否初始化数据库, 默认值: false
	EnableInitDb bool
}

/*
1. 创建 iris app
2. 设置日志级别
3. 设置 recover 中间件
4. 设置 logger 中间件
@params settingsName 配置文件名称
*/
func NewIrisApp(c *IrisAppConfig) *iris.Application {

	settings.Init()
	initDb()
	app := iris.New()
	app.Logger().SetLevel("info")
	app.Use(recover.New())
	// 允许跨域
	crs := cors.AllowAll()
	app.UseRouter(crs)
	// 接口允许options方法
	app.AllowMethods(iris.MethodOptions)
	return app
}

func initSettings(settingsName string) {
	settings.Init(settingsName)
}

func initDb() {
	dbUser := settings.GetString("DATABASE_USER")
	dbPwd := settings.GetString("DATABASE_PASSWORD")
	dbHost := settings.GetString("DATABASE_HOST")
	dbPort := settings.GetString("DATABASE_PORT")
	dbName := settings.GetString("DATABASE_DBNAME")
	database.Init(dbUser, dbPwd, dbHost, dbPort, dbName)
}

func initCache() {
	// 初始化redis，不需要可删除
	host := settings.GetString("REDIS_HOST")
	password := settings.GetString("REDIS_PASSWORD")
	db := settings.GetInt("REDIS_DB")
	cache.NewRedisCache(host, password, db)
	ca := cache.GetCache()
	_, err := ca.Ping()
	if err != nil {
		panic(err)
	}
}

func notFound(ctx iris.Context) {
	ctx.StatusCode(iris.StatusNotFound)
	ctx.JSON(iris.Map{
		"code":    -1,
		"message": "请求的资源不存在",
	})
}

func internalServerError(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code":    -1,
		"message": "error",
	})
}

func initSwagger(app *iris.Application) {
	config := &swagger.Config{
		// The url pointing to API definition.
		URL: settings.GetString("SWAGGER_URL"),
		// DeepLinking: true,
	}
	swaggerUI := swagger.CustomWrapHandler(config, swaggerFiles.Handler)
	app.Get("/swagger", swaggerUI)
	// app.Get("/swagger", func(ctx iris.Context) {
	// 	ctx.JSON(iris.Map{"message": "xxxx"})
	// })
	app.Get("/swagger/{any:path}", swaggerUI)
}

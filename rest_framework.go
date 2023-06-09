package iris_rest_framework

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/weiheguang/iris_rest_framework/auth"
	"github.com/weiheguang/iris_rest_framework/cache"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/logging"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
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
	// 是否启用swagger, 默认值: false, swagger 无法做到完全与main函数分离,不在这里初始化
	// EnableSwagger bool
	// 启用cache, 选项: cache.CacheTypeMem, cache.CacheTypeRedis, 默认值: cache.CacheTypeMem
	CacheType string
	// 是否初始化数据库, 默认值: false
	EnableDb bool
	// Auth处理中间件, 默认值: nil
	Auth auth.IAuth
	// 启用jwt中间件
	EnableJwt bool
}

func GetLogger() *logging.IRFLogger {
	return logging.GetLogger()
}

// 获取全局唯一的cache, 默认返回memcache
func GetCache() cache.ICache {
	return cache.GetCache()
}

// 获取全局唯一的settings
func GetSettings() *settings.Settings {
	return settings.GetSettings()
}

/*
1. 创建 iris app
2. 设置日志级别
3. 设置 recover 中间件
4. 设置 logger 中间件
@params settingsName 配置文件名称
*/
func NewIrisApp(c *IrisAppConfig) *iris.Application {
	if c == nil {
		c = &IrisAppConfig{}
	}
	// 初始化配置文件
	if c.SettingsName == "" {
		c.SettingsName = "settings.yaml"
	}
	settings.Init(c.SettingsName)

	// 初始化数据库
	if c.EnableDb {
		dbUser := settings.GetString("DATABASE_USER")
		dbPwd := settings.GetString("DATABASE_PASSWORD")
		dbHost := settings.GetString("DATABASE_HOST")
		dbPort := settings.GetString("DATABASE_PORT")
		dbName := settings.GetString("DATABASE_DBNAME")
		sqlDebug := settings.GetBool("SQL_DEBUG")
		database.Init(dbUser, dbPwd, dbHost, dbPort, dbName, sqlDebug)
	}
	// 初始化缓存
	if c.CacheType == cache.CacheTypeRedis {
		host := settings.GetString("REDIS_HOST")
		password := settings.GetString("REDIS_PASSWORD")
		db := settings.GetInt("REDIS_DB")
		cache.InitRedis(host, password, db)
	} else {
		cache.InitMem()
	}
	// 初始化app
	app := iris.New()
	app.Logger().SetLevel("info")
	app.Use(recover.New())
	// 允许跨域
	crs := cors.AllowAll()
	app.UseRouter(crs)
	// 接口允许options方法
	app.AllowMethods(iris.MethodOptions)
	// 初始化swagger
	// if c.EnableSwagger {
	// 	config := &swagger.Config{
	// 		// The url pointing to API definition.
	// 		URL: c.settings.GetString("SWAGGER_URL"),
	// 		// DeepLinking: true,
	// 	}
	// 	swaggerUI := swagger.CustomWrapHandler(config, swaggerFiles.Handler)
	// 	app.Get("/swagger", swaggerUI)
	// 	app.Get("/swagger/{any:path}", swaggerUI)
	// }
	if c.EnableJwt {
		secret := settings.GetString("JWT_SECRET")
		jwtMiddleware := jwt.GetJwtMiddleware(secret)
		app.Use(jwtMiddleware.Serve)
	}
	// 初始化auth
	if c.Auth != nil {
		app.Use(c.Auth.Auth)
	}
	// 默认404
	if c.NotFoundHandler == nil {
		app.OnErrorCode(iris.StatusNotFound, notFound)
	}
	// 默认500
	if c.InternalServerErrorHandler == nil {
		app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	}
	return app
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

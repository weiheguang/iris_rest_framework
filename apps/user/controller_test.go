package user

import (
	"fmt"
	"testing"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework/alias"
	"github.com/weiheguang/iris_rest_framework/cache"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/irisapp"
	"github.com/weiheguang/iris_rest_framework/logging"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
	"github.com/weiheguang/iris_rest_framework/rftests"
	"github.com/weiheguang/iris_rest_framework/settings"
)

// var logger = logging.GetLogger()

// 测试参考
// https://github.com/kataras/iris/blob/master/_examples/testing/httptest/main_test.go

/*
参数:

	token: 有效

结果

	http status: httptest.StatusOk
	remote_user: "testid"
*/
const (
	CREATE_USER_TABLE = `CREATE TABLE auth_user (
		id varchar(40) COMMENT "用户ID",
		username varchar(30) NOT NULL COMMENT "用户名",
		password varchar(128) NULL COMMENT "用户密码",
		is_superuser tinyint(1) DEFAULT 0 COMMENT "是否超级用户",
		phone varchar(11) NULL COMMENT "用户手机号",
		is_active tinyint(1) DEFAULT 1 COMMENT "用户状态",
		is_del tinyint(1) DEFAULT 1 COMMENT "软删除",
		created_at datetime(6) NULL COMMENT "创建时间",
		PRIMARY KEY (id),
		Unique KEY auth_user_username_jmjhg(username) USING BTREE,
		Unique KEY auth_user_phone_dljfe(phone) USING BTREE
	);`
	INSERT_USER_SQL = `insert into auth_user (id, username, password, is_superuser, phone, is_active, is_del, created_at) 
	values (?, ?, ?, ?, ?, ?, ?, ?);`
)

/*
测试登录接口, 该接口无需token
*/
func TestUser0(t *testing.T) {
	// 初始化系统
	c := irisapp.IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     true,
		AuthFunc:     UserAuth,
		EnableJwt:    true,
	}
	app := irisapp.NewIrisApp(&c)
	// 初始化测试数据
	dbName := rftests.GetTestDbName()
	var userId = "testid"
	var username = "testname"
	var password = "test"
	var enyPwd = MakePassword(password)
	var isSuperuser = 1
	var isActive = 1
	var isDel = 0
	var phone = "13812345678"
	var createdAt = time.Now()
	db := database.GetDb()
	rftests.CreateDatabase(db, dbName)
	rftests.CreateTables(db, CREATE_USER_TABLE)
	rftests.InsertData(db, INSERT_USER_SQL, userId, username, enyPwd, isSuperuser, phone, isActive, isDel, createdAt)

	// 注册应用
	Register(app, "/api/user")
	e := httptest.New(t, app)
	// 发送登录请求
	data := alias.Map{
		"phone":    phone,
		"password": password,
	}
	e.POST("/api/user/loginbypwd").WithJSON(data).Expect().Status(httptest.StatusOK).JSON().Object().IsValueEqual("code", 0)
	// 清理数据库
	rftests.DropDatabase(db, dbName)
}

/*
测试获取用户信息接口:
不传token场景, 返回 401
*/
func TestUser1(t *testing.T) {
	// 初始化系统
	c := irisapp.IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     true,
		AuthFunc:     UserAuth,
		EnableJwt:    true,
	}
	app := irisapp.NewIrisApp(&c)
	// 初始化测试数据
	dbName := rftests.GetTestDbName()
	var userId = "testid"
	var username = "testname"
	var password = "test"
	var enyPwd = MakePassword(password)
	var isSuperuser = 1
	var isActive = 1
	var isDel = 0
	var phone = "13812345678"
	var createdAt = time.Now()
	db := database.GetDb()
	rftests.CreateDatabase(db, dbName)
	rftests.CreateTables(db, CREATE_USER_TABLE)
	rftests.InsertData(db, INSERT_USER_SQL, userId, username, enyPwd, isSuperuser, phone, isActive, isDel, createdAt)

	// 注册应用
	Register(app, "/api/user")
	e := httptest.New(t, app)
	e.GET("/api/user").Expect().Status(httptest.StatusUnauthorized)
}

/*
测试获取用户信息接口:
传token场景, 正常user信息
*/
func TestUser2(t *testing.T) {
	// 初始化系统
	c := irisapp.IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     true,
		AuthFunc:     UserAuth,
		EnableJwt:    true,
	}
	app := irisapp.NewIrisApp(&c)
	// 初始化测试数据
	dbName := rftests.GetTestDbName()
	var userId = "testid"
	var username = "testname"
	var password = "test"
	var enyPwd = MakePassword(password)
	var isSuperuser = 1
	var isActive = 1
	var isDel = 0
	var phone = "13812345678"
	var createdAt = time.Now()
	db := database.GetDb()
	rftests.CreateDatabase(db, dbName)
	rftests.CreateTables(db, CREATE_USER_TABLE)
	rftests.InsertData(db, INSERT_USER_SQL, userId, username, enyPwd, isSuperuser, phone, isActive, isDel, createdAt)
	token := genToken(userId)
	tokenmsg := fmt.Sprintf("Bearer %s", token)

	// 注册应用
	Register(app, "/api/user")
	PrintRouter(app)
	e := httptest.New(t, app)
	result := alias.Map{
		"code": 0,
		"data": alias.Map{
			"id":            userId,
			"username":      username,
			"phone":         phone,
			"is_authorized": true,
		},
	}
	e.GET("/api/user").WithHeader("Authorization", tokenmsg).Expect().Status(httptest.StatusOK).JSON().IsEqual(result)
}

// 生成token
func genToken(userId string) string {
	secret := settings.GetString("JWT_SECRET")
	expireIn := time.Duration(3600) * time.Second
	issuer := ""
	token := jwt.GenTokenHS256(secret, userId, expireIn, issuer)
	return token
}

func PrintRouter(app *iris.Application) {
	for _, only := range app.GetRoutes() {
		if (only.Method == "POST") || (only.Method == "GET") {
			logging.GetLogger().Debugf("name:%s, path:%s, Method:%s", only.Name, only.Path, only.Method)
		}
	}
}

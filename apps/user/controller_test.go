package user

import (
	"testing"
	"time"

	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework/cache"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/irisapp"
	"github.com/weiheguang/iris_rest_framework/rftests"
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

func TestUser(t *testing.T) {
	// 初始化系统
	c := irisapp.IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     true,
		AuthFunc:     UserIDAuth,
		EnableJwt:    true,
	}
	app := irisapp.NewIrisApp(&c)
	// 注册应用
	Register(app, "/api/user")
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

	e := httptest.New(t, app)
	// 发送登录请求
	data := map[string]interface{}{
		"phone":    phone,
		"password": password,
	}
	// logger.Info(e.POST("/api/user/loginbypwd").WithJSON(data).Expect().Status(httptest.StatusOK).JSON())
	// 测试结果的 code 值为 0

	e.POST("/api/user/loginbypwd").WithJSON(data).Expect().Status(httptest.StatusOK).JSON().Object().IsValueEqual("code", 0)
	// 清理数据库
	rftests.DropDatabase(db, dbName)
}

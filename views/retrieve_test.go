package views

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/irisapp"
	"github.com/weiheguang/iris_rest_framework/rftests"
)

func setUpRetrieve(dbName string) {
	// settings.Init("test_settings")
	// dbUser := settings.GetString("DATABASE_USER")
	// dbPwd := settings.GetString("DATABASE_PASSWORD")
	// dbHost := settings.GetString("DATABASE_HOST")
	// dbPort := settings.GetString("DATABASE_PORT")
	// // 这个时候db还不存在, 设置为空
	// database.Init(dbUser, dbPwd, dbHost, dbPort, "", true)

	// 创建测试数据
	db := database.GetDb()
	rftests.CreateDatabase(db, dbName)
	rftests.CreateTables(db, rftests.TEST_TABLE_USER_SQL)
	// 插入数据
	rftests.InsertData(db, rftests.INSERT_USER_SQL, "test1", 18)
	rftests.InsertData(db, rftests.INSERT_USER_SQL, "test2", 19)
}

// 清理测试数据
func clearUpRetrieve(dbName string) {
	db := database.GetDb()
	rftests.DropDatabase(db, dbName)
}

// 购物车流程测试
/*
	1 空购物车
	2 添加一个商品
*/
func TestGetByPk(t *testing.T) {
	c := irisapp.IrisAppConfig{
		SettingsName: "test_settings",
		// CacheType:    cache.CacheTypeMem,
		// AuthFunc:     testUserIDNil,
		EnableDb: true,
		// EnableJwt: true,
	}
	app := irisapp.NewIrisApp(&c)
	// 初始化数据
	dbName := rftests.GetTestDbName()
	setUpRetrieve(dbName)
	app.Logger().SetLevel("info")
	app.Use(recover.New())
	app.Use(logger.New())

	// appParty := app.Party("/api/user")
	// {
	// 	mApp := mvc.New(appParty)
	// 	c := NewUserController()
	// 	mApp.Handle(c)
	// }

	app.Get("/api/user/{id:int}", GetBy)
	e := httptest.New(t, app)

	id := 1
	fmt.Println("----- 测试开始: id=", id, " -----")
	result := e.GET("/api/user/{pk:int}", id).Expect().Status(iris.StatusOK).JSON() //.Object().Value("code").IsEqual(rferrors.CodeSuc)
	fmt.Println(result)
	fmt.Println("----- 测试结束: id= ", id, " -----")
	// 不存在
	// id = 30855
	// e.GET("/api/user/{pk:int}", id).Expect().Status(iris.StatusOK).JSON().Object().Value("code").IsEqual(rferrors.Code404)

	clearUpRetrieve(dbName)

}

func GetBy(ctx iris.Context) {
	pk := ctx.Params().GetUint64Default("id", 0)
	conf := &RetrieveAPIViewConf{
		Model: rftests.User{},
		// db:    database.GetDb(),
	}
	// logger := ctx.Application().Logger()
	// logger.Infof("pk: %d", pk)
	rv := NewRetrieveAPIView(conf)
	data := rv.GetBy(ctx, pk)
	ctx.JSON(data)
}

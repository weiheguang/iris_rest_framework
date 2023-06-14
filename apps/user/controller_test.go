package user

import (
	"fmt"
	"testing"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework/auth"
	"github.com/weiheguang/iris_rest_framework/cache"
	"github.com/weiheguang/iris_rest_framework/irisapp"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
	"github.com/weiheguang/iris_rest_framework/settings"
)

// 测试参考
// https://github.com/kataras/iris/blob/master/_examples/testing/httptest/main_test.go

/*
参数:

	token: 有效

结果

	http status: httptest.StatusOk
	remote_user: "testid"
*/
func TestUser(t *testing.T) {
	c := irisapp.IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     false,
		AuthFunc:     UserIDAuth,
		EnableJwt:    true,
	}
	app := irisapp.NewIrisApp(&c)
	secret := settings.GetString("JWT_SECRET")
	expireIn := time.Duration(3600) * time.Second
	issuer := ""
	app.Get("/api/ping", func(ctx iris.Context) {
		// remoteUser := ctx.GetHeader(jwt.DefayktUserIDKey)
		// fmt.Println("remoteUser=", remoteUser)
		user := auth.GetUser(ctx)
		userID, _ := user.GetID()
		ctx.JSON(iris.Map{
			"message":     "ok",
			"code":        0,
			"remote_user": userID,
		})
	})
	e := httptest.New(t, app)
	token := jwt.GenTokenHS256(secret, USER_ID, expireIn, issuer)
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	// fmt.Println(tokenmsg)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().
		Status(httptest.StatusOK).JSON().Object().IsValueEqual("remote_user", USER_ID)
}

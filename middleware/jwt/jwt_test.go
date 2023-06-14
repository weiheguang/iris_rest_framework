package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	// "github.com/weiheguang/iris_rest_framework/middleware/jwt"
)

// 测试参考
// https://github.com/kataras/iris/blob/master/_examples/testing/httptest/main_test.go

const (
	USER_ID = "testid"
	SECRET  = "123456"
)

/*
参数:

	token: 有效

结果

	http status: httptest.StatusOk
	remote_user: "testid"
*/
func TestJwt(t *testing.T) {

	app := iris.New()
	expireIn := time.Duration(3600) * time.Second
	issuer := ""
	var jwtMiddleware = GetJwtMiddleware(SECRET)
	app.Use(jwtMiddleware.Serve)
	app.Get("/api/ping", func(ctx iris.Context) {
		remoteUser := ctx.GetHeader(DefaultUserIDKey)
		fmt.Println("remoteUser=", remoteUser)
		ctx.JSON(iris.Map{
			"message":     "ok",
			"code":        0,
			"remote_user": remoteUser,
		})
	})
	e := httptest.New(t, app)
	token := GenTokenHS256(SECRET, USER_ID, expireIn, issuer)
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	// fmt.Println(tokenmsg)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().Status(httptest.StatusOK).JSON().
		Object().IsValueEqual("remote_user", USER_ID)
}

/*
参数:

	token: 无效

结果

	http status: httptest.StatusUnauthorized
	remote_user: ""
*/
func TestJwtInvalidToken(t *testing.T) {
	app := iris.New()
	expireIn := time.Duration(3600) * time.Second
	issuer := ""
	var jwtMiddleware = GetJwtMiddleware(SECRET)
	app.Use(jwtMiddleware.Serve)
	app.Get("/api/ping", func(ctx iris.Context) {
		remoteUser := ctx.GetHeader(DefaultUserIDKey)
		fmt.Println("remoteUser=", remoteUser)
		// a := assert.New(t)
		// a.Equal(remoteUser, "1")
		ctx.JSON(iris.Map{
			"message":     "ok",
			"code":        0,
			"remote_user": remoteUser,
		})
	})
	e := httptest.New(t, app)
	// 无效token
	token := GenTokenHS256("SECRET", USER_ID, expireIn, issuer)
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	// fmt.Println(tokenmsg)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().Status(httptest.StatusUnauthorized)
}

/*
参数:

	token: 不传

结果

	remote_user: ""
*/
func TestJwtNoToken(t *testing.T) {

	app := iris.New()
	var jwtMiddleware = GetJwtMiddleware(SECRET)
	app.Use(jwtMiddleware.Serve)
	app.Get("/api/ping", func(ctx iris.Context) {
		remoteUser := ctx.GetHeader(DefaultUserIDKey)
		fmt.Println("remoteUser=", remoteUser)
		ctx.JSON(iris.Map{
			"message":     "ok",
			"code":        0,
			"remote_user": remoteUser,
		})
	})
	e := httptest.New(t, app)
	e.GET("/api/ping").Expect().Status(httptest.StatusOK).JSON().
		Object().IsValueEqual("remote_user", "")
}

/*
参数:

	token: 不传
	remote_user: "testid",

结果

	remote_user: ""
*/
func TestJwtUserId(t *testing.T) {
	app := iris.New()
	var jwtMiddleware = GetJwtMiddleware(SECRET)
	app.Use(jwtMiddleware.Serve)
	app.Get("/api/ping", func(ctx iris.Context) {
		remoteUser := ctx.GetHeader(DefaultUserIDKey)
		fmt.Println("remoteUser=", remoteUser)
		ctx.JSON(iris.Map{
			"message":     "ok",
			"code":        0,
			"remote_user": remoteUser,
		})
	})
	e := httptest.New(t, app)
	e.GET("/api/ping").WithHeader(DefaultUserIDKey, USER_ID).Expect().
		Status(httptest.StatusOK).JSON().Object().IsValueEqual("remote_user", "")
}

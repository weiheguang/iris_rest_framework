package irisapp

import (
	"fmt"
	"testing"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework/auth"
	"github.com/weiheguang/iris_rest_framework/cache"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
	"github.com/weiheguang/iris_rest_framework/settings"
)

// 测试参考
// https://github.com/kataras/iris/blob/master/_examples/testing/httptest/main_test.go

const (
	USER_ID = "testid"
)

// type MyUserModel struct {
// 	Id string
// }

// func (m *MyUserModel) GetID() string {
// 	return m.Id
// }

// func (m *MyUserModel) GetUsername() string {
// 	return m.Id
// }

// func (m *MyUserModel) IsAuthorized() bool {
// 	return true
// }

// func (m *MyUserModel) GetPhone() string {
// 	return ""
// }

// func testUserID(ctx iris.Context) *auth.User {
// 	ctx.User()
// 	userId := ctx.GetHeader(jwt.DefaultUserIDKey)
// 	um := MyUserModel{
// 		Id: userId,
// 	}
// 	user := &auth.User{}
// 	user.SetUserModel(&um)
// 	return user
// }

// 生成token
func genToken() string {
	secret := settings.GetString("JWT_SECRET")
	expireIn := time.Duration(3600) * time.Second
	issuer := ""
	token := jwt.GenTokenHS256(secret, USER_ID, expireIn, issuer)
	return token
}

/*
参数:

	token: 有效
	AuthFunc: 返回正常的 auth.User

结果

	http status: httptest.StatusOk
	remote_user: "testid"
	is_authorized: true
*/
func TestApp1(t *testing.T) {
	c := IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     false,
		AuthFunc:     auth.UserIDAuth,
		EnableJwt:    true,
	}
	app := NewIrisApp(&c)
	app.Get("/api/ping", func(ctx iris.Context) {
		user := auth.GetUser(ctx)
		userID, _ := user.GetID()
		ctx.JSON(iris.Map{
			"message":       "ok",
			"code":          0,
			"remote_user":   userID,
			"is_authorized": user.IsAuthorized,
		})
	})
	e := httptest.New(t, app)
	// token := jwt.GenTokenHS256(secret, USER_ID, expireIn, issuer)
	token := genToken()
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().
		Status(httptest.StatusOK).JSON().Object().IsValueEqual("remote_user", USER_ID).IsValueEqual("is_authorized", true)
}

/*
参数:

	token: 有效
	AuthFunc: nil

结果

	http status: httptest.StatusOk
	remote_user: "testid"
*/
func TestApp2(t *testing.T) {
	c := IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     false,
		EnableJwt:    true,
	}
	app := NewIrisApp(&c)

	app.Get("/api/ping", func(ctx iris.Context) {
		// remoteUser := ctx.GetHeader(jwt.DefayktUserIDKey)
		// fmt.Println("remoteUser=", remoteUser)
		user := auth.GetUser(ctx)
		userID, _ := user.GetID()
		ctx.JSON(iris.Map{
			"message":       "ok",
			"code":          0,
			"remote_user":   userID,
			"is_authorized": user.IsAuthorized,
		})
	})
	e := httptest.New(t, app)
	// token := jwt.GenTokenHS256(secret, USER_ID, expireIn, issuer)
	token := genToken()
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	// fmt.Println(tokenmsg)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().
		Status(httptest.StatusOK).JSON().Object().IsValueEqual("remote_user", "").IsValueEqual("is_authorized", false)
}

/*
参数:

	token: 有效
	AuthFunc: 返回 nil

结果

	http status: httptest.StatusOk
	remote_user: "testid"
*/

func testUserIDNil(ctx iris.Context) *auth.User {
	return nil
}

func TestApp3(t *testing.T) {
	c := IrisAppConfig{
		SettingsName: "test_settings",
		CacheType:    cache.CacheTypeMem,
		AuthFunc:     testUserIDNil,
		EnableDb:     false,
		EnableJwt:    true,
	}
	app := NewIrisApp(&c)
	app.Get("/api/ping", func(ctx iris.Context) {
		user := auth.GetUser(ctx)
		userID, _ := user.GetID()
		ctx.JSON(iris.Map{
			"message":       "ok",
			"code":          0,
			"remote_user":   userID,
			"is_authorized": user.IsAuthorized,
		})
	})
	e := httptest.New(t, app)
	token := genToken()
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().
		Status(httptest.StatusOK).JSON().Object().IsValueEqual("remote_user", "").IsValueEqual("is_authorized", false)
}

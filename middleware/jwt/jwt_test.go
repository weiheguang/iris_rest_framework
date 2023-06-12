package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework/app/auth"
	// "github.com/weiheguang/iris_rest_framework/middleware/jwt"
)

// 测试参考
// https://github.com/kataras/iris/blob/master/_examples/testing/httptest/main_test.go
func TestListView(t *testing.T) {
	// conf := iris_rest_framework.IrisAppConfig{
	// 	SettingsName: "test_settings",
	// 	EnableDb:     false,
	// 	Auth:         auth.NewUserIDAuth(),
	// 	EnableJwt:    true,
	// }
	// app := iris_rest_framework.NewIrisApp(&conf)
	app := iris.New()
	secret := "123456"
	id := "testid"
	expireIn := time.Duration(3600) * time.Second
	issuer := ""
	var jwtMiddleware = GetJwtMiddleware(secret)
	app.Use(jwtMiddleware.Serve)
	auth := auth.NewUserIDAuth()
	app.Use(auth.Auth)
	app.Get("/api/ping", func(ctx iris.Context) {
		userID, _ := ctx.User().GetID()
		ctx.JSON(iris.Map{
			"message": "ok",
			"code":    0,
			"user_id": userID,
		})
	})
	e := httptest.New(t, app)
	// token := jwtMiddleware.GetToken(claims)

	token := GenTokenHS256(secret, id, expireIn, issuer)
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	// fmt.Println(tokenmsg)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().Status(httptest.StatusOK).JSON().Object().IsValueEqual("user_id", id)
	// fmt.Println(e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().Headers())
}

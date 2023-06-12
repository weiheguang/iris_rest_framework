package iris_app

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework/app/auth"
	"github.com/weiheguang/iris_rest_framework/cache"
)

// 测试参考
// https://github.com/kataras/iris/blob/master/_examples/testing/httptest/main_test.go
func TestListView(t *testing.T) {
	c := IrisAppConfig{
		SettingsName: "",
		CacheType:    cache.CacheTypeMem,
		EnableDb:     false,
		Auth:         auth.NewUserIDAuth(),
		EnableJwt:    true,
	}

	app := NewIrisApp(&c)
	// secret := settings.GetString("JWT_SECRET")
	id := "testid"
	// expireIn := time.Duration(3600) * time.Second
	// issuer := ""
	app.Get("/api/ping", func(ctx iris.Context) {
		userID, _ := ctx.User().GetID()
		ctx.JSON(iris.Map{
			"message": "ok",
			"code":    0,
			"user_id": userID,
		})
	})
	e := httptest.New(t, app)
	// token := jwt.GenTokenHS256(secret, id, expireIn, issuer)
	token := "1"
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	// fmt.Println(tokenmsg)
	e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().
		Status(httptest.StatusOK).JSON().Object().IsValueEqual("user_id", id)
}

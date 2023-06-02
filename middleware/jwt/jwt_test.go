package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/weiheguang/iris_rest_framework"
	// "github.com/weiheguang/iris_rest_framework/middleware/jwt"
)

func TestListView(t *testing.T) {
	app := iris_rest_framework.NewIrisApp("")
	secret := "123456"
	id := "testid"
	expireIn := time.Duration(3600) * time.Second
	issuer := ""
	var jwtMiddleware = GetJwtMiddleware(secret)
	app.Use(jwtMiddleware.Serve)
	app.Get("/api/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "ok"})
	})
	e := httptest.New(t, app)
	// token := jwtMiddleware.GetToken(claims)

	token := GenTokenHS256(secret, id, expireIn, issuer)
	tokenmsg := fmt.Sprintf("Bearer %s", token)
	// fmt.Println(tokenmsg)
	// e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().Header(DefayktUserIDKey).IsEqual(id)
	fmt.Println(e.GET("/api/ping").WithHeader("Authorization", tokenmsg).Expect().Headers())
}

package jwt

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/kataras/iris/v12"
// 	"github.com/kataras/iris/v12/httptest"
// 	"github.com/weiheguang/iris_rest_framework"
// 	// "github.com/weiheguang/iris_rest_framework/middleware/jwt"
// )

// func TestListView(t *testing.T) {
// 	app := iris_rest_framework.NewIrisApp("")
// 	claims := RegisteredClaims{
// 		Issuer:   "test",
// 		Subject:  "t",
// 		Audience: []string{},
// 		ID:       "testid",
// 	}
// 	var jwtMiddleware = GetJwtMiddleware(claims)
// 	app.Use(jwtMiddleware.Serve)
// 	app.Get("api/ping", func(ctx iris.Context) {
// 		ctx.JSON(iris.Map{"message": "ok"})
// 	})
// 	e := httptest.New(t, app)
// 	token := jwtMiddleware.GetToken(claims)
// 	// e.GET("/api/ping").Expect().Header(DefayktUserIDKey).IsEqual("testid")
// 	fmt.Println(e.GET("/api/ping").Expect().Headers())
// }

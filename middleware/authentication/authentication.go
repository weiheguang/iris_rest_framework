package authentication

import "github.com/kataras/iris/v12"

type IAuth interface {
	Auth(ctx iris.Context) *User
}

// 生成 auth middleware
func AuthMiddleware(a IAuth) iris.Handler {
	return func(ctx iris.Context) {
		user := a.Auth(ctx)
		ctx.Values().Set("user", user)
		ctx.Next()
	}
}

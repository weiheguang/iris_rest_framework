package auth

import "github.com/kataras/iris/v12"

type AuthFunc = func(ctx iris.Context) *User

func AuthMiddlewareFunc(af AuthFunc) iris.Handler {
	return func(ctx iris.Context) {
		user := af(ctx)
		ctx.SetUser(user)
		ctx.Next()
	}
}

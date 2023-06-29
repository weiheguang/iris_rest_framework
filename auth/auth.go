package auth

import (
	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/alias"
	"github.com/weiheguang/iris_rest_framework/logging"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
)

type AuthFunc = func(ctx iris.Context) *User

/*
如果授权函数返回有效的用户, 则设置到ctx.User()中
如果授权还是没提供或者没有返回有效用户, 则设置一个空的用户到ctx.User()中
空的User默认为未授权状态
*/
func AuthMiddlewareFunc(af AuthFunc) iris.Handler {
	return func(ctx iris.Context) {
		var user *User
		if af != nil {
			user = af(ctx)
		}
		if user == nil {
			user = &User{
				UserModel:    nil,
				Id:           "",
				Username:     "",
				Phone:        "",
				IsAuthorized: false,
			}
		}
		ctx.SetUser(user)
		ctx.Next()
	}
}

/*
默认提供授权函数
*/
func UserIDAuth(ctx iris.Context) *User {
	userID := ctx.GetHeader(jwt.DefaultUserIDKey)
	user := &User{
		Id:           userID,
		IsAuthorized: true,
	}
	return user
}

// 登录请求中间件
func LoginRequire(ctx iris.Context) {
	user := GetUser(ctx)
	logging.GetLogger().Debug("user----------: ", user)
	if !user.IsAuthorized {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(alias.Map{
			"code":    -1,
			"message": "未登录(-6)",
		})
		ctx.StopExecution()
		return
	}
	ctx.Next()
}

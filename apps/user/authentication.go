package user

import (
	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/auth"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
)

func UserIDAuth(ctx iris.Context) *auth.User {
	userID := ctx.GetHeader(jwt.DefaultUserIDKey)
	// 查询user model 信息
	um := AuthUser{
		ID: userID,
	}
	user := &auth.User{}
	user.SetUserModel(&um)
	return user
}

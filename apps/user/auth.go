package user

import (
	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/auth"
	"github.com/weiheguang/iris_rest_framework/logging"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
)

func UserAuth(ctx iris.Context) *auth.User {
	userID := ctx.GetHeader(jwt.DefaultUserIDKey)
	logging.GetLogger().Debug("userID: ", userID)
	// 根据 userId查询出来 user model
	um := &AuthUser{}
	user := &auth.User{
		Id:           userID,
		IsAuthorized: false,
		Username:     "",
		Phone:        "",
		Model:        nil,
	}
	// 如果没有 userId 则返回
	if userID == "" {
		return user
	}
	// 根据 userId查询出来 user model
	err := um.GetByID(userID)
	if err != nil {
		logging.GetLogger().Error(err.Error())
		return user
	}
	user.Model = um
	user.IsAuthorized = true
	user.Username = um.Username
	user.Phone = um.Phone
	return user
}

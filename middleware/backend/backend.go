package backend

import (
	"github.com/kataras/iris/v12"
)

const DefaultUserHeaderKey = "REMOTE_USER"

type IAuth interface {
	Auth(ctx iris.Context)
}

// 使用header里面的 字段 remote_user 作为用户ID进行授权
// type UserIDAuth struct {
// }

// func NewUserIDAuth() *UserIDAuth {
// 	return &UserIDAuth{}
// }

// func (a *UserIDAuth) Auth(ctx iris.Context) {
// 	userID := ctx.GetHeader(DefaultUserHeaderKey)
// 	// myLogger.Info("user_id=", userID)
// 	user := NewUser(userID)
// 	// 无需验证 UserID的有效性.有效验证在上层
// 	// 注意这里使用的是user指针
// 	ctx.SetUser(user)
// 	ctx.Next()
// }

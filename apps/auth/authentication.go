package auth

import (
	"github.com/kataras/iris/v12"
)

const DefaultUserHeaderKey = "REMOTE_USER"

// 实现 authentication.IAuth 接口
type UserIDAuth struct {
}

func NewUserIDAuth() *UserIDAuth {
	return &UserIDAuth{}
}

func (a *UserIDAuth) Auth(ctx iris.Context) *User {
	userID := ctx.GetHeader(DefaultUserHeaderKey)
	user := NewUser(userID)
	return user
}

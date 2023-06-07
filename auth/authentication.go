package auth

import (
	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/logging"
	"github.com/weiheguang/iris_rest_framework/settings"
)

const DefaultUserHeaderKey = "REMOTE_USER"
const ENCRYPTION_TIMES = 10000

var myLogger = logging.GetLogger()

// 授权接口定义
type IAuth interface {
	Auth(ctx iris.Context)
}

// 使用header里面的 字段 remote_user 作为用户ID进行授权
type UserIDAuth struct {
}

func NewUserIDAuth() *UserIDAuth {
	return &UserIDAuth{}
}

func (a *UserIDAuth) Auth(ctx iris.Context) {
	// 获取配置文件里面的 UserIDHeaderKey
	userIdKey := settings.GetString("USER_ID_HEADER_KEY")
	if userIdKey == "" {
		userIdKey = DefaultUserHeaderKey
	}
	userID := ctx.GetHeader(userIdKey)

	myLogger.Info("user_id=", userID)
	user := NewUser(userID)
	// 无需验证 UserID的有效性.有效验证在上层
	// 注意这里使用的是user指针
	ctx.SetUser(user)
	ctx.Next()
}

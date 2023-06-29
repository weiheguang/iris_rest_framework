package user

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Register(app *iris.Application, apiPrefix string) {
	userAppParty := app.Party(apiPrefix)
	{
		userApp := mvc.New(userAppParty)
		userController := NewUserController()
		userApp.Handle(userController)
		userApp.Controllers[0].GetRoute("PostLoginbypwd").Name = "手机密码登录"
		userApp.Controllers[0].GetRoute("Get").Name = "获取登录用户的信息"
		userApp.Controllers[0].GetRoute("GetBy").Name = "根据用户id获取用户信息"
	}
}

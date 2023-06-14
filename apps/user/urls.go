package user

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Register(app *iris.Application, apiPrefix string) {
	authAppParty := app.Party(apiPrefix)
	{
		authApp := mvc.New(authAppParty)
		authController := NewAuthController()
		authApp.Handle(authController)

	}
}

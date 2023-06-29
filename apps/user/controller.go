package user

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/weiheguang/iris_rest_framework/auth"
	"github.com/weiheguang/iris_rest_framework/logging"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
	"github.com/weiheguang/iris_rest_framework/response"
	"github.com/weiheguang/iris_rest_framework/settings"
)

type IUserController interface {
	PostLoginbypwd(ctx iris.Context) response.IRFResult   // 手机密码登录
	Get(ctx iris.Context) response.IRFResult              // 获取登录用户的信息
	GetBy(ctx iris.Context, id string) response.IRFResult // 根据用户id获取用户信息
}

type UserController struct {
}

func NewUserController() IUserController {
	return &UserController{}
}

// @Summary		根据用户手机号和密码获取 token
// @Tags			user
// @Accept			json
// @Description	### 返回：200
// @Description	```
// @Description	{
// @Description	"code": 0
// @Description	"data": {
// @Description	"token": "eameeljfaldkfaldkf.alkfjelfkelfe.alekfjelfjkelfjelfj"
// @Description	},
// @Description	}
// @Description	```
// @Router			/user [post]
func (c *UserController) PostLoginbypwd(ctx iris.Context) response.IRFResult {
	us := LoginByPwdSerializer{}
	if err := ctx.ReadJSON(&us); err != nil {
		return response.ResponseResult(nil, 1, err)
	}
	user := AuthUser{}
	pwd := MakePassword(us.Password)
	err := user.GetByPwd(us.Phone, pwd)
	if err != nil {
		return response.ResponseResult(nil, 1, err)
	}
	secret := settings.GetString("JWT_SECRET")
	expireIn := settings.GetDuration("JWT_EXPIRE_IN")
	issuer := settings.GetString("APP_NAME")
	token := jwt.GenTokenHS256(secret, user.Id, expireIn, issuer)
	data := iris.Map{
		"token": token,
	}
	return response.ResponseResult(data, 0, nil)
}

// @Summary		根据用户id获取用户信息
// @Tags			auth
// @Accept			json
// @Description	### 返回：200
// @Router			/user [get]
func (c *UserController) Get(ctx iris.Context) response.IRFResult {
	logging.GetLogger().Debug("Get")
	user := auth.GetUser(ctx)
	if !user.IsAuthorized {
		return response.ResponseResult(nil, -401, nil)
	}
	data := iris.Map{
		"id":            user.Id,
		"username":      user.Username,
		"phone":         user.Phone,
		"is_authorized": user.IsAuthorized,
	}
	return response.ResponseResult(data, 0, nil)
}

func (c *UserController) GetBy(ctx iris.Context, id string) response.IRFResult {
	return response.ResponseResult(nil, 0, nil)
}

// 添加 LoginRequire 中间件
func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	// logging.GetLogger().Debug("BeforeActivation")
	// 相对路径的根路径开始
	b.Handle("GET", "/", "Get", auth.LoginRequire)
	// 或者甚至可以添加基于这个控制器路由的全局中间件，
	// 比如在这个例子里面的根路由 "/" :
	// b.Router().Use(myMiddleware)
}

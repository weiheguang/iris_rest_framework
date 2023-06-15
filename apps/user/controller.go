package user

import (
	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/middleware/jwt"
	"github.com/weiheguang/iris_rest_framework/response"
	"github.com/weiheguang/iris_rest_framework/settings"
)

type IAuthController interface {
	PostLoginbypwd(ctx iris.Context) response.IRFResult // 手机密码登录
}

type AuthController struct {
}

func NewAuthController() IAuthController {
	return &AuthController{}
}

// @Summary		根据用户手机号和密码获取 token
// @Tags			auth
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
// @Router			/auth [post]
func (c *AuthController) PostLoginbypwd(ctx iris.Context) response.IRFResult {
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
	token := jwt.GenTokenHS256(secret, user.ID, expireIn, issuer)
	data := iris.Map{
		"token": token,
	}
	return response.ResponseResult(data, 0, nil)
}

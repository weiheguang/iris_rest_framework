package user

// 手机号登录
type LoginByPwdSerializer struct {
	Phone    string `json:"phone" validate:"number,len=11"`
	Password string `json:"password" validate:"min=6,max=12"`
}

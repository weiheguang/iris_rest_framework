package auth

import (
	"errors"
	"time"

	"github.com/kataras/iris/v12"
)

var ErrNotSupported = errors.New("not supported")

// 数据库model
type IUserModel interface {
	GetID() string       // 用户id
	GetUsername() string // 用户名
	IsAuthorized() bool  // 是否授权
}

// 实现 iris 的 User 接口: https://github.com/kataras/iris/blob/master/context/context_user.go
type User struct {
	um IUserModel // User Model
}

// 获取 UserModel
func (u *User) GetUserModel() IUserModel {
	return u.um
}

// 设置用户 UserModel
func (u *User) SetUserModel(au IUserModel) {
	u.um = au
}

// GetRaw should return the raw instance of the user, if supported.
func (u *User) GetRaw() (interface{}, error) {
	return u.um, ErrNotSupported
}

// GetAuthorization should return the authorization method,
// e.g. Basic Authentication.
func (u *User) GetAuthorization() (string, error) {
	return "jwt", nil
}

func (u *User) GetAuthorizedAt() (time.Time, error) {
	return time.Now(), nil
}

// GetID should return the ID of the User.

func (u *User) GetID() (string, error) {
	return u.um.GetID(), nil
}

// GetUsername should return the name of the User.

func (u *User) GetUsername() (string, error) {
	return u.um.GetUsername(), nil
}

// GetPassword should return the encoded or raw password
// (depends on the implementation) of the User.

func (u *User) GetPassword() (string, error) {
	return "", ErrNotSupported
}

// GetEmail should return the e-mail of the User.

func (u *User) GetEmail() (string, error) {
	return "", ErrNotSupported
}

// GetRoles should optionally return the specific user's roles.
// Returns `ErrNotSupported` if this method is not
// implemented by the User implementation.

func (u *User) GetRoles() ([]string, error) {
	return nil, ErrNotSupported
}

// GetToken should optionally return a token used
// to authorize this User.

func (u *User) GetToken() ([]byte, error) {
	return nil, ErrNotSupported
}

// GetField should optionally return a dynamic field
// based on its key. Useful for custom user fields.
// Keep in mind that these fields are encoded as a separate JSON key.

func (u *User) GetField(key string) (interface{}, error) {
	return "", ErrNotSupported
}

func (u *User) IsAuthorized() bool {
	return u.um.IsAuthorized()
}

func GetUser(ctx iris.Context) *User {
	return ctx.Values().Get("user").(*User)
}

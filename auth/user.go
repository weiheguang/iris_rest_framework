package auth

import (
	"errors"
	"reflect"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/logging"
)

var ErrNotSupported = errors.New("not supported")

// 数据库model
type IUserModel interface {
	GetID() string
}

// 实现 iris 的 User 接口: https://github.com/kataras/iris/blob/master/context/context_user.go
type User struct {
	UserModel    IUserModel `json:"fields,omitempty"` // User Model
	Id           string     `json:"id"`
	Username     string     `json:"username"`
	Phone        string     `json:"phone"`
	IsAuthorized bool       `json:"is_authorized"`
}

// 获取 UserModel
func (u *User) GetUserModel() IUserModel {
	return u.UserModel
}

// 设置用户 UserModel
func (u *User) SetUserModel(au IUserModel) {
	u.UserModel = au
}

// GetRaw should return the raw instance of the user, if supported.
func (u *User) GetRaw() (interface{}, error) {
	return u.UserModel, ErrNotSupported
}

// GetAuthorization should return the authorization method,
// e.g. Basic Authentication.
func (u *User) GetAuthorization() (string, error) {
	return "", ErrNotSupported
}

func (u *User) GetAuthorizedAt() (time.Time, error) {
	return time.Now(), nil
}

// GetID should return the ID of the User.

func (u *User) GetID() (string, error) {
	return u.Id, nil
}

// GetUsername should return the name of the User.

func (u *User) GetUsername() (string, error) {
	return u.Username, nil
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
	if u.UserModel == nil {
		return nil, ErrNotSupported
	}
	if reflect.TypeOf(u.UserModel).String() == "struct" {
		// 根据字段名字返回字段值
		return reflect.ValueOf(u.UserModel).FieldByName(key), nil
	} else {
		return nil, errors.New("UserModel is not struct")
	}
}

// 新建用户
func NewUser(id string, phone string, username string, isAuthorized bool, um IUserModel) *User {
	return &User{
		UserModel:    um,
		Id:           id,
		Phone:        phone,
		Username:     username,
		IsAuthorized: isAuthorized,
	}
}

/*
 * 获取系统用户
 */
func GetUser(ctx iris.Context) *User {
	user, ok := ctx.User().(*User)
	if ok {
		return user
	}
	logging.GetLogger().Error("user is not *User type")
	// 返回空 未授权的 User
	user = &User{
		UserModel:    nil,
		Id:           "",
		Username:     "",
		Phone:        "",
		IsAuthorized: false,
	}
	return nil
}

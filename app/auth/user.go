package auth

import (
	"errors"
	"time"
)

var ErrNotSupported = errors.New("not supported")

// 实现 iris 的 User 接口: https://github.com/kataras/iris/blob/master/context/context_user.go
type User struct {
	au *AuthUser // User Model
}

func NewUser(user_id string) *User {
	u := &User{}
	userModel := AuthUser{ID: user_id}
	u.SetUserModel(&userModel)
	return u
}

// 获取 UserModel
func (u *User) GetUserModel() *AuthUser {
	return u.au
}

// 设置用户 UserModel
func (u *User) SetUserModel(au *AuthUser) {
	u.au = au
}

// GetRaw should return the raw instance of the user, if supported.
func (u *User) GetRaw() (interface{}, error) {
	return u.au, ErrNotSupported
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
	return u.au.ID, nil
}

// GetUsername should return the name of the User.

func (u *User) GetUsername() (string, error) {
	return u.au.Username, nil
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
	// myLogger.Debug("self=", u)
	return u.au.ID != ""
}

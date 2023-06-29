package user

import (

	// "time"

	"errors"

	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/datatypes"
	"gorm.io/plugin/soft_delete"
)

// 用户表
// 实现 auth.IUserModel
type AuthUser struct {
	Id          string                `json:"id" gorm:"primaryKey"`
	Password    string                `json:"password"`
	Username    string                `json:"username" gorm:"unique"`
	IsSuperuser bool                  `json:"is_superuser"`
	Phone       string                `json:"phone"`
	IsActive    bool                  `json:"is_active"`
	CreatedAt   datatypes.IRFTime     `json:"created_at" gorm:"autoCreateTime"`
	IsDel       soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag"`
}

// 实现 auth.IUserModel
func (u *AuthUser) GetID() string {
	return u.Id
}

// 根据 id查询出来 user model
func (u *AuthUser) GetByID(id string) error {
	db := database.GetDb()
	result := db.Where("id = ? and is_active= ?", id, true).First(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据用户名的密码获取User对象
func (u *AuthUser) GetByPwd(phone string, pwd string) error {
	db := database.GetDb()
	result := db.Where("phone = ? and password = ? and is_active= ?", phone, pwd, true).First(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据用户名的密码获取User对象
func (u *AuthUser) Save() error {
	if u.Id == "" {
		return errors.New("用户ID不能为空")
	}
	if u.Username == "" {
		return errors.New("用户名不能为空")
	}
	if u.Phone == "" {
		return errors.New("用户手机号不能为空")
	}
	db := database.GetDb()
	result := db.Save(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

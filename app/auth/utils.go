package auth

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/datatypes"
)

const ENCRYPTION_TIMES = 10000

func MakePassword(password string) string {
	/*
		生成密码, 使用明文密码加密密码: 默认生成数据
		bjc_md5$10000$encrypted_string:
		bjc: 前缀
		MD5: 加密方法
		10000: 循环次数
		encrypted_string: 加密字符串
	*/
	var hash string = password
	var hashTmp [16]byte
	// var hashTmp [16]byte
	for i := 0; i < ENCRYPTION_TIMES; i++ {
		hashTmp = md5.Sum([]byte(hash))
		hash = hex.EncodeToString(hashTmp[:])
	}
	pwd := fmt.Sprintf("%s_%s$%d$%s", "bjc", "md5", ENCRYPTION_TIMES, hash)
	return pwd
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateUsername(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 从 ctx获取 User
func GetUser(ctx iris.Context) (*User, error) {
	// myLogger.Debug("从ctx获取 user")
	// TODO: 是否需要将 user存入cache
	xx := ctx.User()
	// xx 是 User指针, 类型转换的时候需要使用 *User
	user, ok := xx.(*User)
	if ok {
		// myLogger.Debug("ok=", ok)
		return user, nil
	}
	// myLogger.Debug("user=", user)
	return nil, errors.New("获取 User 出错")
}

// 从 ctx获取 User
func LoginRequire(ctx iris.Context) {
	// 查找header里面是否有已经解码的 jwt
	user, err := GetUser(ctx)
	if err != nil {
		// myLogger.Debug("err")
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"code":    -1,
			"message": "未登录(-6)",
		})
	}
	if user.IsAuthorized() {
		ctx.Next()
	} else {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"code":    -1,
			"message": "未登录(-4)",
		})
	}
}

func CreateAuthUser(phone string, password string, isAdmin bool) (*AuthUser, error) {
	// 创建用户
	uuid := uuid.New()
	key := uuid.String()
	key = strings.ReplaceAll(key, "-", "")
	pwd := MakePassword(password)
	authUser := AuthUser{
		ID:          key,
		Password:    pwd,
		Username:    GenerateUsername(10),
		IsSuperuser: isAdmin,
		Phone:       phone,
		IsActive:    true,
		CreatedAt:   datatypes.IRFTime{},
		IsDel:       0,
	}
	if err := authUser.Save(); err != nil {
		return nil, err
	}
	return &authUser, nil
}

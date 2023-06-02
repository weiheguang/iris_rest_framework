package views

import (
	"errors"
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/response"

	// "github.com/weiheguang/iris_rest_framework/rferrors"

	"gorm.io/gorm"
)

// type ListAPIView struct {
// 	model interface{} // 设置model
// }

// func (v *ListAPIView) SetModel(mi interface{}) {
// 	v.model = mi
// }

// ------------------- 根据主键获取单个对象 -------------------
type RetrieveAPIView struct {
	conf *RetrieveAPIViewConf
}

type RetrieveAPIViewConf struct {
	Model interface{} // 设置 model
	// db    *database.Db
	// 暂时无需制定 pk, 因为 gorm 会自动根据主键查询
	// PkName string // 设置主键名称, 默认为 id
}

func NewRetrieveAPIView(conf *RetrieveAPIViewConf) *RetrieveAPIView {
	return &RetrieveAPIView{
		conf: conf,
	}
}

func (v *RetrieveAPIView) GetBy(ctx iris.Context, pk interface{}) response.IRFResult {
	// if v.config.PkName == "" {
	// 	v.config.PkName = ""
	// }
	var err error
	if v.conf.Model == nil {
		// panic("APIView: model 未设置")
		err = errors.New("APIView: model 未设置")
		return response.ResponseResult(nil, 0, err)
	}
	db := database.GetDb()
	// if db == nil {
	// 	err = errors.New("APIView: db 未设置")
	// 	return response.ResponseResult(nil, 0, err)

	// }
	// if v.config.PkName == "" {
	// 	panic("APIView: 主见名称未设置 pkName")
	// }
	// pk = ctx.Params().GetStringDefault(v.config.PkName, "0")
	reflectType := reflect.TypeOf(v.conf.Model)
	// fmt.Println("type of reflectType:", reflectType)
	// 如果传入的是指针
	isPtr := reflectType.Kind() == reflect.Ptr
	if isPtr {
		reflectType = reflectType.Elem()
	}
	if reflectType.Kind() != reflect.Struct {
		panic("RetrieveAPIView: 传入的 model 必须是 Struct 或者 Struct 指针")
	}
	returnValue := reflect.New(reflectType)
	mi := returnValue.Interface()
	// db := database.GetDb()
	// gorm会自动根据主键查询
	result := db.First(mi, pk)
	// var err *rferrors.ApiError

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return response.ResponseResult(mi, 0, err)
	}
	if result.Error != nil {
		err = errors.New(result.Error.Error())
		return response.ResponseResult(mi, 0, err)
	}
	return response.ResponseResult(mi, 0, err)
}

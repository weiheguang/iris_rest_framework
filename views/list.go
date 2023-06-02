package views

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/response"

	"golang.org/x/exp/slices"
)

// ------------------- 查询列表 -------------------
type ListAPIView struct {
	conf *ListAPIViewConf
}

type ListAPIViewConf struct {
	Model interface{} // 设置 model
	// db    *gorm.DB    // 设置 db
	// 分页放到 view里面解析
	// Page         int      // 设置分页, 默认为 1
	// PageSize     int      // 设置分页大小, 默认为 10
	FilterFields []string // 设置过滤字段
}

func NewListAPIView(conf *ListAPIViewConf) *ListAPIView {
	return &ListAPIView{
		conf: conf,
	}
}

func (v *ListAPIView) List(ctx iris.Context) response.IRFResult {
	// var err *rferrors.ApiError
	var err error
	if v.conf.Model == nil {
		err = errors.New("ListAPIView: model 未设置")
		return response.ResponseResult(nil, 0, err)
	}
	// 判断 配置的
	// 处理分页 FilterFields 是否在 model 的字段中
	// 获取 model 的字段
	t := reflect.TypeOf(v.conf.Model)
	// fmt.Println("type of t:", t)
	// 获取字段数量
	fieldNum := t.NumField()
	// fmt.Println("fieldNum:", fieldNum)
	fieldTagJsonList := make([]string, 0)
	// 遍历字段
	for i := 0; i < fieldNum; i++ {
		// 获取字段
		field := t.Field(i)
		// fmt.Println("field:", field)
		// 获取字段名称
		// fieldName := field.Name
		// fmt.Println("fieldName:", fieldName)
		// 获取字段 tag
		fieldTag := field.Tag
		// fmt.Println("fieldTag:", fieldTag)
		// 获取字段 tag 中的 json
		fieldTagJson := fieldTag.Get("json")
		// fmt.Println("fieldTagJson:", fieldTagJson)
		// 获取字段 tag 中的 json 中的名称
		// fieldTagJsonName := strings.Split(fieldTagJson, ",")[0]
		// fmt.Println("fieldTagJsonName:", fieldTagJsonName)
		// 判断是否在过滤字段中
		// fmt.Printf("fieldName=%s , type=%v , fieldTagJson=%s \n", fieldName, field.Type, fieldTagJson)
		fieldTagJsonList = append(fieldTagJsonList, fieldTagJson)
	}
	fmt.Println("fieldTagJsonList:", fieldTagJsonList)
	// 判断过滤条件是否为空, 如果过滤条件为空, 则不进行过滤
	db := database.GetDb()
	if len(v.conf.FilterFields) > 0 {
		// 判断过滤字段是否在 model 的字段中
		for _, filterField := range v.conf.FilterFields {
			// 判断是否在 model 的字段中
			if !slices.Contains(fieldTagJsonList, filterField) {
				fmt.Println("filterField:", filterField)
				err = errors.New("ListAPIView: 过滤字段不在 model 的字段中")
				return response.ResponseResult(nil, 0, err)
			}
		}
		// 获取参数,生成查询条件
		for key, value := range ctx.URLParams() {
			fmt.Println("key:", key, "value:", value)
			// 判断是否为过滤字段, split key, 如果 key 中包含 __, 则按照 __ 分割, 否则默认为 exact
			fieldName, fieldExpr := "", ""
			if strings.Contains(key, "__") {
				fieldName = strings.Split(key, "__")[0]
				fieldExpr = strings.Split(key, "__")[1]
				// field.value = value
			} else {
				fieldName = key
				fieldExpr = EXACT
				// field.value = value
			}
			// fmt.Println("fieldName:", fieldName, "fieldExpr:", fieldExpr)
			// 判断fieldName v.config.FilterFields 中, 判断fieldExpr 是否在允许的范围内
			if slices.Contains(v.conf.FilterFields, fieldName) {
				// 判断 fieldExpr 是否在允许的范围内
				switch fieldExpr {
				case "exact":
					db = db.Where(fieldName+" = ?", value)
				case "iexact":
					db = db.Where(fieldName+" = ?", value)
				case "contains":
					db = db.Where(fieldName+" like ?", "%"+value+"%")
				case "icontains":
					db = db.Where(fieldName+" ilike ?", "%"+value+"%")
				case "gt":
					db = db.Where(fieldName+" > ?", value)
				case "gte":
					db = db.Where(fieldName+" >= ?", value)
				case "lt":
					db = db.Where(fieldName+" < ?", value)
				case "lte":
					db = db.Where(fieldName+" <= ?", value)
				case "in":
					// 需要将 value 转换为 []interface{},
					// TODO: 未完成, 需要根据model的字段类型, 转换为对应的类型. 例如: id 字段为 int, 需要转换为 int
					valueList := strings.Split(value, ",")
					valueListInterface := make([]interface{}, 0)
					for _, v := range valueList {
						valueListInterface = append(valueListInterface, v)
					}
					db = db.Where(fieldName+" in (?)", valueListInterface)
					fmt.Println("case in , fieldName:", fieldName, "valueListInterface:", valueListInterface)
				case "startswith":
					db = db.Where(fieldName+" like ?", value+"%")
				case "istartswith":
					db = db.Where(fieldName+" ilike ?", value+"%")
				case "endswith":
					db = db.Where(fieldName+" like ?", "%"+value)
				case "iendswith":
					db = db.Where(fieldName+" ilike ?", "%"+value)
				case "range":
					db = db.Where(fieldName+" between ? and ?", value)
				case "year":
					db = db.Where(fieldName+" = ?", value)
				case "month":
					db = db.Where(fieldName+" = ?", value)
				case "day":
					db = db.Where(fieldName+" = ?", value)
				case "isnull":
					db = db.Where(fieldName + " is null")
				}
			}
		}
	}

	// 处理分页参数
	page, _ := strconv.Atoi(ctx.URLParamDefault("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.URLParamDefault("page_size", "10"))
	// fmt.Println("page:", page, "pageSize:", pageSize)
	// 根据类型生成 slice
	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	// modelList := reflect.Zero(arrayType)
	// fmt.Println("type of sliceModel:", sliceModel)
	// 生成 slice 的接口
	smi := slice.Interface()
	// gorm 不定条件查询
	// for k, v := range filterFields {
	// 	db = db.Where(k+" = ?", v)
	// }
	// 查询数据

	if err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&smi).Error; err != nil {
		// err = rferrors.New(err.Error(), -1)
		return response.ResponseResult(nil, 0, err)
	}
	return response.ResponseResult(smi, 0, err)
}

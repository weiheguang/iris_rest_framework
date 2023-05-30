package response

import "github.com/kataras/iris/v12"

type IRFResult interface{}

// 封装http请求统一返回的数据结构
func ResponseResult(data interface{}, code int, err error) IRFResult {

	if err != nil {
		return iris.Map{
			"code":    code,
			"message": err.Error(),
		}
	} else {
		return iris.Map{
			"code": code,
			"data": data,
		}
	}
}

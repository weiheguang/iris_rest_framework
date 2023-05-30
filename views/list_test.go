package views

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/weiheguang/iris_rest_framework/database"
	"github.com/weiheguang/iris_rest_framework/rftests"
	"github.com/weiheguang/iris_rest_framework/settings"
)

/*
	查询条件:
		exact: 精确查询
		contains: 包含查询
		lt: 小于查询
		lte: 小于等于查询
		gt: 大于查询
		gte: 大于等于查询
		in: 包含查询
		startswith: 开头查询
		endswith: 结尾查询
		year: 年查询
		month: 月查询
		day: 日查询
		hour: 时查询
		minute: 分查询
		second: 秒查询
		isnull: 是否为空查询
		search: 搜索查询
		regex: 正则查询
		iregex: 不区分大小写正则查询

		django rest framework 查询条件:
			https://www.django-rest-framework.org/api-guide/filtering/#searchfilter
			__exact 精确等于 like ‘aaa’
			__iexact 精确等于 忽略大小写 ilike ‘aaa’
			__contains 包含 like ‘%aaa%’
			__icontains 包含 忽略大小写 ilike ‘%aaa%’，但是对于sqlite来说，contains的作用效果等同于icontains。
			__gt 大于
			__gte 大于等于
			__lt 小于
			__lte 小于等于
			__in 存在于一个list范围内
			__startswith 以…开头
			__istartswith 以…开头 忽略大小写
			__endswith 以…结尾
			__iendswith 以…结尾，忽略大小写
			__range 在…范围内
			__year 日期字段的年份
			__month 日期字段的月份
			__day 日期字段的日
			__isnull=True/False

		https://github.com/miki725/django-url-filter
		# get user with id 5
		example.com/users/?id=5

		# get user with id either 5, 10 or 15
		example.com/users/?id__in=5,10,15

		# get user with id between 5 and 10
		example.com/users/?id__range=5,10

		# get user with username "foo"
		example.com/users/?username=foo

		# get user with username containing case insensitive "foo"
		example.com/users/?username__icontains=foo

		# get user where username does NOT contain "foo"
		example.com/users/?username__icontains!=foo
*/

func setUpList(dbName string) {
	settings.Init("test_settings")
	dbUser := settings.GetString("DATABASE_USER")
	dbPwd := settings.GetString("DATABASE_PASSWORD")
	dbHost := settings.GetString("DATABASE_HOST")
	dbPort := settings.GetString("DATABASE_PORT")
	// dbName := ""
	database.Init(dbUser, dbPwd, dbHost, dbPort, "")

	// 创建测试数据
	db := database.GetDb()
	rftests.CreateDatabase(db, dbName)
	rftests.CreateTables(db, rftests.TEST_TABLE_USER_SQL)
	// 循环插入数据
	for i := 1; i <= 10; i++ {
		xxx := fmt.Sprintf("test%d", i)
		rftests.InsertData(db, rftests.INSERT_USER_SQL, xxx, 18+i)
	}
}

// 清理测试数据
func clearUpList(dbName string) {
	db := database.GetDb()
	rftests.DropDatabase(db, dbName)
}

func TestListView(t *testing.T) {
	dbName := rftests.GetTestDbName()
	setUpList(dbName)
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/api/pdt", List)
	// app.Get("/api/pdt2", List2)
	e := httptest.New(t, app)

	query_param := "page=2&page_size=2&age__gt=18"
	fmt.Println("----- 测试: querey_param: ", query_param)
	ok_result := map[string]interface{}{
		"code": 0,
		// "message": "",
		"data": []interface{}{
			map[string]interface{}{
				"id":   3,
				"name": "test3",
				"age":  21,
			},
			map[string]interface{}{
				"id":   4,
				"name": "test4",
				"age":  22,
			},
		},
	}
	// 测试: __gt: 小于等于查询: "page=2&page_size=2&age__gt=18"
	e.GET("/api/pdt").WithQueryString(query_param).Expect().Status(iris.StatusOK).JSON().IsEqual(ok_result)
	fmt.Println("---- 测试结束: querey_param: ", query_param)
	// 测试 ?name__exact=xx 精确等于查询
	query_param = "name__exact=test3"
	fmt.Println("----- 测试: querey_param: ", query_param)
	ok_exact_result := map[string]interface{}{
		"code": 0,
		// "message": "",
		"data": []interface{}{
			map[string]interface{}{
				"id":   3,
				"name": "test3",
				"age":  21,
			},
		},
	}
	e.GET("/api/pdt").WithQueryString(query_param).Expect().Status(iris.StatusOK).JSON().IsEqual(ok_exact_result)
	// 测试 ?name=xxx
	query_param = "name=test3"
	fmt.Println("----- 测试开始: querey_param: ", query_param)
	e.GET("/api/pdt").WithQueryString(query_param).Expect().Status(iris.StatusOK).JSON().IsEqual(ok_exact_result)
	fmt.Println("---- 测试结束: querey_param: ", query_param)
	// 测试 ?name__contains=xxx
	query_param = "name__contains=test"
	fmt.Println("----- 测试开始: querey_param: ", query_param)
	ok_contains_result := map[string]interface{}{
		"code": 0,
		// "message": "",
		"data": []interface{}{
			map[string]interface{}{
				"age":  19,
				"id":   1,
				"name": "test1",
			},
			map[string]interface{}{
				"age":  20,
				"id":   2,
				"name": "test2",
			},
			map[string]interface{}{
				"age":  21,
				"id":   3,
				"name": "test3",
			},
			map[string]interface{}{
				"age":  22,
				"id":   4,
				"name": "test4",
			},
			map[string]interface{}{
				"age":  23,
				"id":   5,
				"name": "test5",
			},
			map[string]interface{}{
				"age":  24,
				"id":   6,
				"name": "test6",
			},
			map[string]interface{}{
				"age":  25,
				"id":   7,
				"name": "test7",
			},
			map[string]interface{}{
				"age":  26,
				"id":   8,
				"name": "test8",
			},
			map[string]interface{}{
				"age":  27,
				"id":   9,
				"name": "test9",
			},
			map[string]interface{}{
				"age":  28,
				"id":   10,
				"name": "test10",
			},
		}}
	e.GET("/api/pdt").WithQueryString(query_param).Expect().Status(iris.StatusOK).JSON().IsEqual(ok_contains_result)
	fmt.Println("---- 测试结束: querey_param: ", query_param)

	// 测试 ?name__startswith=xxx
	query_param = "name__startswith=test"
	fmt.Println("----- 测试开始: querey_param: ", query_param)
	ok_startswith_result := ok_contains_result
	e.GET("/api/pdt").WithQueryString(query_param).Expect().Status(iris.StatusOK).JSON().IsEqual(ok_startswith_result)
	fmt.Println("---- 测试结束: querey_param: ", query_param)

	// 测试 ?name__endswith=xxx
	query_param = "name__endswith=test"
	fmt.Println("----- 测试开始: querey_param: ", query_param)
	ok_endswith_result := map[string]interface{}{
		"code": 0,
		// "message": "",
		"data": []interface{}{}}
	e.GET("/api/pdt").WithQueryString("name__endswith=test").Expect().Status(iris.StatusOK).JSON().IsEqual(ok_endswith_result)
	fmt.Println("---- 测试结束: querey_param: ", query_param)

	// 测试 ?id__in=xxx,xxx
	query_param = "id__in=3,4"
	fmt.Println("----- 测试开始: querey_param: ", query_param)
	ok_in_result := map[string]interface{}{
		"code": 0,
		// "message": "",
		"data": []interface{}{
			map[string]interface{}{
				"id":   3,
				"name": "test3",
				"age":  21,
			},
			map[string]interface{}{
				"id":   4,
				"name": "test4",
				"age":  22,
			},
		},
	}
	e.GET("/api/pdt").WithQueryString(query_param).Expect().Status(iris.StatusOK).JSON().IsEqual(ok_in_result)
	fmt.Println("---- 测试结束: querey_param: ", query_param)

	clearUpList(dbName)

}

// 测试
func List(ctx iris.Context) {
	config := &ListAPIViewConf{
		Model: rftests.User{},
		db:    database.GetDb(),
		// 设置查询字段
		FilterFields: []string{"name", "age", "id"},
	}
	logger := ctx.Application().Logger()
	// logger.Infof("pk: %d", pk)
	rv := NewListAPIView(config)
	data := rv.List(ctx)
	logger.Info(data)
	// result := response.ResponseResult(data, 0, err)
	ctx.JSON(data)
}

// 测试无效配置, 过滤字段配置错误情况
func List2(ctx iris.Context) {
	config := &ListAPIViewConf{
		Model: rftests.User{},
		db:    database.GetDb(),
		// 设置查询字段
		FilterFields: []string{"name1", "age"},
	}
	logger := ctx.Application().Logger()
	// logger.Infof("pk: %d", pk)
	rv := NewListAPIView(config)
	data := rv.List(ctx)
	logger.Info(data)
	ctx.JSON(data)
}

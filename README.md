# iris_rest_framework

## 介绍

基于iris和gorm的 api框架

## 安装教程

1. xxxx
2. xxxx
3. xxxx

## Settings 配置

使用例子1 :

```go
    settings.Init("test_settings")  // 初始化配置, 全局仅执行一次
    // 直接获取配置
    s := settings.GetSettings()
    debug := s.GetBool("DEBUG")
    fmt.Println(debug)
```

使用例子2:

```go
    settings.Init("test_settings")  // 初始化配置, 全局仅执行一次
    // 通过 Settings 模块的快捷方法获取配置
    debug := settings.GetBool("DEBUG")
    fmt.Println(c.Debug)
```

## db 数据库

使用例子:

```go
    // 初始化数据库连接
    dbUser := settings.GetString("DATABASE_USER")
    dbPwd := settings.GetString("DATABASE_PASSWORD")
    dbHost := settings.GetString("DATABASE_HOST")
    dbPort := settings.GetString("DATABASE_PORT")
    dbName := settings.GetString("DATABASE_NAME")
    database.InitDb(dbUser, dbPwd, dbHost, dbPort, dbName)
    // 使用db
    type User struct {
        Id   uint64
        Name string
        Age  int
    }
    u := User{Id: 1}
    db.Find(&u)
```

## view

### ListAPIView

获取表的列表数据, 使用例子如下:

```go
conf := &ListAPIViewConf{
    Model: User{},           // 要查询的表
    db:    *gorm.Db,         // 数据库连接
    // 设置查询字段
    FilterFields: []string{"name", "age", "id"},
}
rv := NewListAPIView(conf)
data := rv.List(ctx)
```

过滤字段说明:
过滤字段遵守django的过滤字段规则. 过滤字段必须被预先配置到 FilterFields 里面,否则过滤字段不起作用
分页参数不用配置, 会自动解析
过滤字段规则: 字段名字__过滤方法, 如: name__contains=weiheguang, "__" 为双横线分隔符
通过url查询字段, 如: /api/user?name=weiheguang&age=18
支持的过滤字段方法:

* 分页: page=1&page_size=10
* 精确等于: name=weiheguang 或者 name__exact=weiheguang.
* 大于: age__gt=18
* 大于等于: age__gte=18
* 包含: name__contains=weiheguang
* 字符串以xxx开头: name__startswith=xxx
* 字符串以xxx结尾: name__endswith=xxx
* in: age__in=18,19,20

### RetrieveAPIView

根据pk获取单条数据, 使用例子如下:

```go
pk := ctx.Params().GetUint64Default("id", 0) // 获取url中的id
conf := &RetrieveAPIViewConf{
    Model: rftests.User{},      // 要查询的表
    db:    database.GetDb(),    // 数据库连接
}
rv := NewRetrieveAPIView(conf)
data := rv.GetBy(ctx, pk)
```

## 测试

* 运行所有测试: go test -v ./...
* 运行一个包内测试: go test -v ./views
* 运行一个包内测试: go test -v ./middleware/jwt
* 运行一个单独的测试: go test github.com/weiheguang/iris_rest_framework/views -run TestGetByPk

## 参与贡献

1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request

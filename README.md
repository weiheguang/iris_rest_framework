# iris_rest_framework

## 介绍

基于iris和gorm的 api框架

## 安装

```shell
    go get github.com/weiheguang/iris_rest_framework
```

## NewIrisApp

返回iris的app对象, 使用例子如下:

```go
    c := IrisAppConfig{
        SettingsName string
        // 默认值: notFound
        NotFoundHandler iris.Handler
        // 默认值: internalServerError
        InternalServerErrorHandler iris.Handler
        // 是否启用swagger, 默认值: false, swagger 无法做到完全与main函数分离,不在这里初始化
        // EnableSwagger bool
        // 启用cache, 选项: cache.CacheTypeMem, cache.CacheTypeRedis, 默认值: cache.CacheTypeMem
        CacheType string
        // 是否初始化数据库, 默认值: false
        EnableDb bool
        // Auth处理函数, 返回user指针, 默认值: nil
        AuthFunc UserIdAuth
        // 启用jwt中间件, jwt中间件依赖环境变量: JWT_SECRET
        EnableJwt bool
    }
    app := NewIrisApp(&c)
    app.Run(iris.Addr(":8080"))
```

## Settings 使用

使用例子2:

```go
    // 通过 Settings 模块的快捷方法获取配置
    debug := settings.GetBool("DEBUG")
```

### 配置项

```yaml
    # 数据库配置
    DATABASE_USER: "user"         # 数据库用户名
    DATABASE_PASSWORD: "password" # 数据库密码
    DATABASE_HOST: ""             # 数据库链接地址
    DATABASE_PORT: 3306           # 数据库端口
    DATABASE_DBNAME: "db"         # 数据库名字
    SQL_DEBUG: false              # 是否开启sql语句打印
    DEBUG: false                  # 是否开启debug模式

    # Redis缓存配置
    REDIS_HOST: ""             # redis链接地址
    REDIS_PASSWORD: "password" # redis密码
    REDIS_DB: 0                # redis数据库

    JWT_SECRET: "123456"       # jwt secret
    JWT_EXPIRE_IN: "7200"      # jwt 过期时间, 单位: 秒

    APP_NAME: "app"            # app 名字

```

## db 数据库

使用例子:

```go
    // 初始化数据库连接
    db := database.GetDb()
    u := User{Id: 1}
    db.Find(&u)
```

## view

### ListAPIView

获取表的列表数据, 使用例子如下:

```go
conf := &ListAPIViewConf{
    // 要查询的表
    Model: User{},           
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
* 如有其他需求,欢迎提issue

### RetrieveAPIView

根据pk获取单条数据, 使用例子如下:

```go
    pk := ctx.Params().GetUint64Default("id", 0) // 获取url中的id
    conf := &RetrieveAPIViewConf{
        Model: rftests.User{},      // 要查询的表
    }
    rv := NewRetrieveAPIView(conf)
    data := rv.GetBy(ctx, pk)
```

## Cache使用说明

### 初始化

```go
    // 初始化redis连接
    ca := cache.GetCache()
    ca.Set("test", "test", 10)
```

### 访问控制

1. 权限入库: R1, R2, R3
2. func checkPre(userId string , r string) bool{}

## 测试

* 运行所有测试: go test -v ./... -run=".*" -cover
* 运行一个包内测试: go test -v ./middleware/jwt -run=".*" -cover
* 运行一个单独的测试: go test -v ./middleware/jwt -run="TestJwtInvalidToken" -cover
* go test -v ./apps/user -run=".*" -cover
* go test -v ./irisapp -run=".*" -cover

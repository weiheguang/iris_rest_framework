package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

var jwtParser = jwt.Parser{}

type Middleware struct {
	Config Config
}

var (
	// ErrTokenMissing is the error value that it's returned when
	// a token is not found based on the token extractor.
	ErrTokenMissing = errors.New("required authorization token not found")

	// ErrTokenInvalid is the error value that it's returned when
	// a token is not valid.
	ErrTokenInvalid = errors.New("token is invalid")

	// // ErrTokenExpired is the error value that it's returned when
	// // a token value is found and it's valid but it's expired.
	ErrTokenExpired = errors.New("token is expired")
)

type (
	// Token for JWT. Different fields will be used depending on whether you're
	// creating or parsing/verifying a token.
	//
	// A type alias for jwt.Token.
	Token = jwt.Token
	// MapClaims type that uses the map[string]interface{} for JSON decoding
	// This is the default claims type if you don't supply one
	//
	// A type alias for jwt.MapClaims.
	MapClaims = jwt.MapClaims
	// Claims must just have a Valid method that determines
	// if the token is invalid for any supported reason.
	//
	// A type alias for jwt.Claims.
	Claims = jwt.Claims
	//
	RegisteredClaims = jwt.RegisteredClaims
)

// Shortcuts to create a new Token.
var (
	// NewToken           = jwt.New
	NewTokenWithClaims = jwt.NewWithClaims
	NewNumericDate     = jwt.NewNumericDate
)

// HS256 and company.
var (
	SigningMethodHS256 = jwt.SigningMethodHS256
	SigningMethodHS384 = jwt.SigningMethodHS384
	SigningMethodHS512 = jwt.SigningMethodHS512
)

// ECDSA - EC256 and company.
var (
	SigningMethodES256 = jwt.SigningMethodES256
	SigningMethodES384 = jwt.SigningMethodES384
	SigningMethodES512 = jwt.SigningMethodES512
)

// A function called whenever an error is encountered
// 错误处理函数
type errorHandler func(iris.Context, error)

// Get returns the user (&token) information for this client/request
// func (m *Middleware) Get(ctx iris.Context) *jwt.Token {
// 	v := ctx.Values().Get(m.Config.ContextKey)
// 	if v == nil {
// 		return nil
// 	}
// 	jwtToken := jwtInfo.(*jwt.Token)
// 	myCliams := jwtToken.Claims.(*jwt.MyClaims)
// 	return v.(*jwt.Token)
// }

// New constructs a new Secure instance with supplied options.
//
//	新建一个中间件实例
func New(cfg ...Config) *Middleware {
	var c Config
	if len(cfg) == 0 {
		c = Config{}
	} else {
		c = cfg[0]
	}

	if c.ContextKey == "" {
		c.ContextKey = DefaultContextKey
	}
	if c.UserIDKey == "" {
		c.UserIDKey = DefayktUserIDKey
	}

	if c.ErrorHandler == nil {
		c.ErrorHandler = OnError
	}

	if c.Extractor == nil {
		c.Extractor = FromAuthHeader
	}

	return &Middleware{Config: c}
}

// iris中间件处理
func (m *Middleware) Serve(ctx iris.Context) {
	if err := m.CheckJWT(ctx); err != nil {
		m.Config.ErrorHandler(ctx, err)
		return
	}
	// If everything ok then call next.
	ctx.Next()
}

// OnError is the default error handler.
// Use it to change the behavior for each error.
// See `Config.ErrorHandler`.
func OnError(ctx iris.Context, err error) {
	if err == nil {
		return
	}

	ctx.StopExecution()
	ctx.StatusCode(iris.StatusUnauthorized)
	// ctx.WriteString(err.Error())
	ctx.JSON(iris.Map{
		"code":    -1,
		"message": err.Error(),
	})
}

func logf(ctx iris.Context, format string, args ...interface{}) {
	ctx.Application().Logger().Debugf(format, args...)
}

// 检查token是否有效
func (m *Middleware) CheckJWT(ctx iris.Context) error {
	// 删掉 header中 jwt key值
	// ctx.Values().Set(m.Config.ContextKey, "")
	// 获取token
	token, err := m.Config.Extractor(ctx)
	fmt.Printf("extracting token: %s\n", token)
	if err != nil {
		logf(ctx, "Error extracting JWT: %v", err)
		return err
	}
	// logf(ctx, "Token extracted: %s", token)
	if token == "" {
		// Check if it was required
		// if m.Config.CredentialsOptional {
		// 	logf(ctx, "No credentials found (CredentialsOptional=true)")
		// 	// No error, just no token (and that is ok given that CredentialsOptional is true)
		// 	return nil
		// }
		// If we get here, the required token is missing
		// logf(ctx, "Error: No credentials found (CredentialsOptional=false)")
		// return ErrTokenMissing
		// token 字段置空
		ctx.Values().Set(m.Config.ContextKey, "")
		// 设置 remote_user 为空
		ctx.Values().Set(m.Config.UserIDKey, "")
		ctx.Request().Header.Set(m.Config.UserIDKey, "")
		return nil
	}
	// 解析token
	parsedToken, err := jwtParser.ParseWithClaims(token, &RegisteredClaims{}, m.Config.ValidationKeyGetter)
	// fmt.Printf("Error parsing token: %v\n", err)
	// fmt.Printf("parsedToken.Valid: %v\n", parsedToken.Valid)
	if err != nil {
		// fmt.Printf("Error parsing token: %v\n", err)
		// logf(ctx, "Error parsing token: %v", err)
		return ErrTokenInvalid
	}
	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		err := fmt.Errorf("expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		fmt.Printf("Error validating token algorithm: %v\n", err)
		// logf(ctx, "Error validating token algorithm: %v", err)
		return ErrTokenInvalid
	}
	// 解析token数据, 类型断言使用指针
	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)
	// fmt.Printf("ok: %v\n", ok)
	// fmt.Printf("claims: %v\n", claims)
	if ok && parsedToken.Valid {
		fmt.Printf("token有效, %v %v\n", claims.ID, claims.ExpiresAt)
		// logf(ctx, "%v, %v", claims.ID, claims.ExpiresAt)
		// 把UserID放到header里面
		// ctx.Header(m.Config.UserIDKey, claims.ID)
		ctx.Request().Header.Set(m.Config.UserIDKey, claims.ID)
		// 解析后的token,暂时放到Values里面,如果有需求,以后可以放到header里面
		ctx.Values().Set(m.Config.ContextKey, claims)
	} else {
		// fmt.Printf("token无效\n")
		// logf(ctx, "token无效")
		// ctx.Header(m.Config.ContextKey, "")
		ctx.Request().Header.Set(m.Config.UserIDKey, "")
		ctx.Values().Set(m.Config.UserIDKey, "")
		return ErrTokenInvalid
	}
	// If we get here, everything worked and we can set the
	// user property in context.
	// 把解析后的token放到context中
	// logf(ctx, "%v", claims)
	// ctx.Values().Set(m.Config.ContextKey, claims)
	// cType := reflect.TypeOf(m.Config.Claims).Kind()
	// id := parsedToken.Claims.(jwt.MapClaims)["id"].(string)
	// ctx.Values().Set(m.Config.UserIDKey, claims.ID)
	// ctx.Values().Set(m.Config.ContextKey, claims)
	return nil
}

// TokenExtractor is a function that takes a context as input and returns
// either a token or an error.  An error should only be returned if an attempt
// to specify a token was found, but the information was somehow incorrectly
// formed.  In the case where a token is simply not present, this should not
// be treated as an error.  An empty string should be returned in that case.
// 定义一个抽取器，接受一个上下文，返回token或者错误
type TokenExtractor func(iris.Context) (string, error)

// 按顺序查找多个token提取器，并返回第一个找到的token
func FromFirst(extractors ...TokenExtractor) TokenExtractor {
	return func(ctx iris.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}

// FromAuthHeader 是 TokenExtractor 的一个实现，它接受一个上下文，并从 Authorization header 中提取 JWT token。
func FromAuthHeader(ctx iris.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}
	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("authorization header format must be Bearer {token}")
	}
	return authHeaderParts[1], nil
}

// 从查询参数中提取token, 参数为查询参数的名称
func FromParameter(param string) TokenExtractor {
	return func(ctx iris.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

func GetJwtMiddleware(secret string) *Middleware {
	config := Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Extractor:     FromFirst(FromAuthHeader, FromParameter("token")),
		// Claims:        claims,
	}
	return New(config)
}

/*
生成HS256 token
@scret: 密钥
@id: 用户id
@expireIn: 过期时间(秒)
@issuer: 签发者
*/
func GenTokenHS256(secret string, id string, expireIn time.Duration, issuer string) string {
	now := time.Now()
	// expiresAt := NewNumericDate(now.Add(expireIn * time.Second))
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   issuer,
		Audience:  []string{},
		ExpiresAt: NewNumericDate(now.Add(expireIn * time.Second)),
		NotBefore: &jwt.NumericDate{},
		IssuedAt:  NewNumericDate(time.Now()),
		ID:        id,
	}
	token := NewTokenWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

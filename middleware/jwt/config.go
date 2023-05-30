package jwt

import "github.com/golang-jwt/jwt/v4"

const (
	// DefaultContextKey jwt
	DefaultContextKey = "jwt"
)

type Config struct {
	// The function that will return the Key to validate the JWT.
	// It can be either a shared secret or a public key.
	// Default value: nil
	ValidationKeyGetter jwt.Keyfunc
	// The function that will be called when there's an error validating the token
	// Default value:
	ErrorHandler errorHandler
	// 默认token抽取器
	Extractor TokenExtractor
	// The name of the property in the request where the user (&token) information
	// from the JWT will be stored.
	// 默认的存储在iris Context中的key, 默认值为"jwt"
	ContextKey string
	// A boolean indicating if the credentials are required or not
	// Default value: false
	CredentialsOptional bool
	// When set, the middelware verifies that tokens are signed with the specific signing algorithm
	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	// Default: nil
	SigningMethod jwt.SigningMethod
}

// wanxiang的token payload
type MyClaims struct {
	*jwt.RegisteredClaims
	// Username string `json:"username"`
	Uid int `json:"uid"`
}

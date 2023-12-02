package jwt

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"go-project-frame/pkg/globalError"
)

var JwtToken jwtToken

type jwtToken struct {
	secret string
}

// InitJwt init jwt
func (j *jwtToken) InitJwt(secret string) {
	j.secret = secret
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          int
	Username    string
	NickName    string
	AuthorityId uint
}

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

// ParseToken parse token to custom claims
func (j *jwtToken) ParseToken(tokenStr string) (claims *CustomClaims, err error) {
	// 使用jwt.ParseWithClaims方法解析token，将前端传过来的tokenStr转成一个*Token类型的对象
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	// 如果发生错误，则先判断是否为 token过期，不是的话就返回 通用的错误信息
	if err != nil {
		if ve, ok := err.(jwt.ValidationError); ok {
			if ve.Errors == jwt.ValidationErrorExpired {
				return nil, errors.New(globalError.GetCodeTag(globalError.AuthorizationExpiredError))
			} else {
				return nil, errors.New(globalError.GetCodeTag(globalError.AuthorizationError))
			}
		}
	}
	// 如果token 不是我们自定义类型 或 不合法，返回 token解析错误
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New(globalError.GetCodeTag(globalError.AuthorizationParseError))
	}
	// token解析成功，返回claims
	return claims, nil
}

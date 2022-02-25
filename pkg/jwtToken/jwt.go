package jwtToken

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration 设置Token的过期时间
const TokenExpireDuration = 2 * time.Hour

// MySecret 加盐
var MySecret = []byte("苏某人大傻子")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(UserID int64, Username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserID:   UserID,
		Username: Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).UnixNano(), // 过期时间
			Issuer:    "Login",                                        // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 检验Token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

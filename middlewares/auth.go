package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"webapp/controllers"
	"webapp/controllers/encode"
	"webapp/pkg/jwtToken"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端请求Token有三种方式 1. 放在请求头 2. 放在请求体 3.放在URL
		// 这里假设放在请求头，根据事情情况实现
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, encode.CodeAuthHeaderNotExist)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, encode.CodeAuthHeaderFailed)
			c.Abort()
			return
		}
		// parts[1] 使我们取到的tokenString
		mc, err := jwtToken.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, encode.CodeTokenFailed)
			c.Abort()
			return
		}
		// 将当前请求userid保存到上下文
		c.Set(controllers.ContextUserIDKey, mc.UserID)
		c.Next()
	}
}

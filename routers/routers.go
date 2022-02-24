package routers

import (
	"github.com/gin-gonic/gin"
	"webapp/controllers"
	"webapp/logger"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册接口
	r.POST("/register", controllers.Register)
	// 登录接口
	r.POST("/login", controllers.Login)
	return r
}

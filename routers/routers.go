package routers

import (
	"github.com/gin-gonic/gin"
	"webapp/controllers"
	"webapp/logger"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 登录接口
	r.POST("/register", controllers.Register)
	return r
}

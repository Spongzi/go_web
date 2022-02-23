package routers

import (
	"github.com/gin-gonic/gin"
	"webapp/logger"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	return r
}

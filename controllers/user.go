package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"webapp/logic"
	"webapp/models"
	"webapp/pkg/translator"
)

func Register(c *gin.Context) {
	// 获取参数并且校验参数
	p := new(models.ParamsUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(translator.Trans),
		})
		return
	}
	// 打印接收到的数据
	fmt.Println(p)
	// 业务处理
	if err := logic.Register(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": c.Writer.Status(),
			"msg":    err.Error(),
		})
		return
	}
	// 返回JSON信息
	c.JSON(http.StatusOK, gin.H{
		"status": c.Writer.Status(),
		"msg":    "注册成功",
	})
}

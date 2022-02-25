package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"webapp/controllers/encode"
	"webapp/logic"
	"webapp/models"
	"webapp/pkg/translator"
)

// Register 注册
func Register(c *gin.Context) {
	// 获取参数并且校验参数
	p := new(models.ParamsUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, encode.CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, encode.CodeInvalidPassword, errs.Translate(translator.Trans))
		return
	}
	// 打印接收到的数据
	fmt.Println(p)
	// 业务处理
	code, err := logic.Register(p)
	if err != nil {
		ResponseError(c, code)
		return
	}
	// 返回JSON信息
	ResponseSuccess(c, code)
}

// Login 登录
func Login(c *gin.Context) {
	// 1. 获取提交的数据
	p := new(models.ParamsLoginUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, encode.CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, encode.CodeInvalidPassword, errs.Translate(translator.Trans))
		return
	}
	fmt.Println(p)
	// 业务逻辑处理
	token, code, err := logic.Login(p)
	if err != nil {
		ResponseErrorWithMsg(c, code, error.Error(err))
		return
	}
	// 返回数据
	ResponseSuccess(c, token)
}

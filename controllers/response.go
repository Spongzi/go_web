package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"webapp/controllers/encode"
)

func ResponseError(c *gin.Context, code encode.ResCode) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  code.GetMsg(),
		"data": nil,
	})
}

type MsgType interface {
	string | validator.ValidationErrorsTranslations
}

func ResponseErrorWithMsg[msgType MsgType](c *gin.Context, code encode.ResCode, msg msgType) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": encode.CodeSuccess,
		"msg":  encode.CodeSuccess.GetMsg(),
		"data": data,
	})
}

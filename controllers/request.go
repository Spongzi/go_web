package controllers

import (
	"github.com/gin-gonic/gin"
	"webapp/controllers/errcode"
)

const ContextUserIDKey = "user_id"

// 获取当前用户的Id
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = errcode.ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = errcode.ErrorUserNotLogin
		return
	}
	return
}

package errcode

import "errors"

var (
	ErrorUserNotLogin   = errors.New("用户没有登录，请登录")
	ErrorUserIsExist    = errors.New("用户名已存在，请更换用户名")
	ErrorUserNotExist   = errors.New("用户不存在")
	ErrorPasswordFailed = errors.New("密码错误")
)

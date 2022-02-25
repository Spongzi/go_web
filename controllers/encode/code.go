package encode

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeAuthHeaderNotExist
	CodeAuthHeaderFailed
	CodeTokenFailed
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:            "Success",
	CodeInvalidParam:       "请求参数错误",
	CodeUserExist:          "用户名已存在",
	CodeUserNotExist:       "用户名不存在",
	CodeInvalidPassword:    "用户名或密码错误",
	CodeServerBusy:         "服务繁忙",
	CodeAuthHeaderNotExist: "请求auth不存在",
	CodeAuthHeaderFailed:   "请求auth格式错误",
	CodeTokenFailed:        "token错误",
}

func (c ResCode) GetMsg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

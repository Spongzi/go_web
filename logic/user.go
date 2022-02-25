package logic

import (
	"webapp/controllers/encode"
	"webapp/controllers/errcode"
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/jwtToken"
	"webapp/pkg/sf"
)

func Register(p *models.ParamsUser) (code encode.ResCode, err error) {
	// 校验用户是否存在
	isExist := mysql.CheckUserExist(p.Username)
	if isExist {
		return encode.CodeUserExist, errcode.ErrorUserIsExist
	}
	// 生成UID
	userId := sf.GenID()
	// 构造一个User实例
	u := models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 插入用户
	err = mysql.InsertUser(&u)
	return encode.CodeSuccess, nil
}

func Login(p *models.ParamsLoginUser) (token string, code encode.ResCode, err error) {
	// 验证数据库里面的人名和密码是否正确
	u := models.User{
		Username: p.Username,
		Password: p.Password,
	}
	code, err = mysql.CheckingPassword(&u)
	if err != nil {
		return "", code, err
	}
	token, err = jwtToken.GenToken(u.UserId, u.Username)
	if err != nil {
		return "", encode.CodeServerBusy, err
	}
	return token, encode.CodeSuccess, nil
}

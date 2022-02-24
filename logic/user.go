package logic

import (
	"errors"
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/sf"
)

func Register(p *models.ParamsUser) (err error) {
	// 校验用户是否存在
	isExist := mysql.CheckUserExist(p.Username)
	if isExist {
		return errors.New("用户名已存在，请更换用户名")
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
	return
}

func Login(p *models.ParamsLoginUser) (err error) {
	// 验证数据库里面的人名和密码是否正确
	u := models.User{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.CheckingPassword(&u)
	if err != nil {
		return err
	}
	return
}

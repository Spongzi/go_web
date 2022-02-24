package logic

import (
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/sf"
)

func Register(p *models.ParamsUser) (err error) {
	// 校验用户是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		// 数据库查询失败
		return
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
	// 密码加密
	return
}

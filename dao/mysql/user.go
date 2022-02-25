package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"webapp/controllers/encode"
	"webapp/models"
)

// CheckUserExist 检查指定用户是否存在
func CheckUserExist(username string) bool {
	sqlStr := "select count(user_id) from user where username = ?;"
	var count int
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return false
	}
	return count > 0
}

// InsertUser 向用户中插入一条新的用户信息
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 插入用户
	sqlStr := "insert into USER (user_id, username, password) values (?, ?, ?);"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	return
}

// CheckingPassword 验证密码
func CheckingPassword(p *models.User) (code encode.ResCode, err error) {
	if !CheckUserExist(p.Username) {
		zap.L().Error("用户名不存在")
		return encode.CodeUserNotExist, errors.New("用户不存在")
	}
	sqlStr := "select password from USER where username = ?;"
	var password string
	err = db.Get(&password, sqlStr, p.Username)
	if err != nil {
		fmt.Println("get data failed!", err)
		return encode.CodeServerBusy, err
	}
	if password != encryptPassword(p.Password) {
		return encode.CodeInvalidPassword, errors.New("密码错误")
	}
	return encode.CodeSuccess, nil
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("suxu.com"))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

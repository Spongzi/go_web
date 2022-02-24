package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"webapp/models"
)

// CheckUserExist 检查指定用户是否存在
func CheckUserExist(username string) error {
	// 查询用户是否重复
	sqlStr := "select count(user_id) from user where username = ?;"
	var count int
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
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

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("suxu.com"))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

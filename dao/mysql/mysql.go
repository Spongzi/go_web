package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"webapp/settings"
)

var db *sqlx.DB

func Init(conf *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Dbname,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect db failed!", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(conf.MaxOpenConn)
	db.SetMaxIdleConns(conf.MaxIdleConn)
	if err != nil {
		return err
	}
	return
}

func Close() {
	err := db.Close()
	if err != nil {
		zap.L().Error("mysql关闭错误", zap.Error(err))
		return
	}
}

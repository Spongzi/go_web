package main

import (
	"fmt"
	"go.uber.org/zap"
	"webapp/dao/mysql"
	"webapp/dao/redis"
	"webapp/logger"
	"webapp/pkg/sf"
	"webapp/pkg/translator"
	"webapp/routers"
	"webapp/settings"
)

func main() {
	// 1. 读取配置
	if err := settings.Init(); err != nil {
		fmt.Println("settings init failed", err)
	}
	// 2. 设置日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("Logger init failed", err)
	}
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			fmt.Println("zap logger sync failed!", err)
		}
	}(zap.L())
	// 3. 连接数据库
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Println("Mysql init failed", err)
	}
	defer mysql.Close()
	// 4. 连接redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("redis init failed", err)
	}
	defer redis.Close()
	// 雪花生成ID
	if err := sf.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("snowflake init failed", err)
	}
	fmt.Println("测试", sf.GenID())
	if err := translator.InitTrans("zh"); err != nil {
		fmt.Println("translator failed", err)
	}
	// 5. 路由管理
	r := routers.SetUp()
	// 6. 启动服务
	panic(r.Run(":8080"))
}

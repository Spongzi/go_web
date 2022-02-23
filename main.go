package main

import (
	"fmt"
	"go.uber.org/zap"
	"webapp/dao/mysql"
	"webapp/dao/redis"
	"webapp/logger"
	"webapp/routers"
	"webapp/settings"
)

func main() {
	// 1. 读取配置
	if err := settings.Init(); err != nil {
		fmt.Println("settings init failed", err)
	}
	// 2. 设置日志
	if err := logger.Init(settings.Conf); err != nil {
		fmt.Println("Logger init failed", err)
	}
	defer zap.L().Sync()
	// 3. 连接数据库
	if err := mysql.Init(settings.Conf); err != nil {
		fmt.Println("Mysql init failed", err)
	}
	defer mysql.Close()
	// 4. 连接redis
	if err := redis.Init(settings.Conf); err != nil {
		fmt.Println("redis init failed", err)
	}
	defer redis.Close()
	// 5. 路由管理
	r := routers.SetUp()
	// 6. 启动服务
	panic(r.Run(":8080"))
}

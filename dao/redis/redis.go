package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"webapp/settings"
)

var rdb *redis.Client

func Init(conf *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v",
			conf.Host,
			conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
		PoolSize: conf.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Error("redis关闭错误", zap.Error(err))
		return
	}
}

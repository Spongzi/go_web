package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("../config.yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig() // 读取配置信息
	viper.WatchConfig()        //检测
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
	})
	return
}

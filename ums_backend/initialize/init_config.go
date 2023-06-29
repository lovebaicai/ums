package initialize

import (
	"fmt"
	"log"

	"ums_backend/config"
	"ums_backend/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(fileConfig string) *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	fmt.Printf("使用配置文件路径为%s\n", fileConfig)
	v.SetConfigFile(fileConfig)
	if err := v.ReadInConfig(); err != nil {
		log.Panicf("配置文件读取失败..")
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	serverConfig := config.Server{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		log.Panicf("配置文件解析失败..")
	}
	global.GVA_CONFIG = serverConfig
	return v
}

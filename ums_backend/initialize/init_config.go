package initialize

import (
	"log"

	"ums_backend/config"
	"ums_backend/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Panicf("配置文件读取失败..")
	}
	serverConfig := config.Server{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		log.Panicf("配置文件解析失败..")
	}
	global.GVA_CONFIG = serverConfig
}

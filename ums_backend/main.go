package main

import (
	"flag"
	"fmt"

	"ums_backend/global"
	"ums_backend/initialize"
	"ums_backend/router"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	router.InitUserRouter(superRoute)
	router.InitLdapRouter(superRoute)
}

func main() {
	var fileConfig string
	flag.StringVar(&fileConfig, "c", "", "choose config file.")
	flag.Parse()
	if flag.NFlag() > 0 {
		initialize.InitConfig(fileConfig)
	} else {
		initialize.InitConfig("./config.yaml")
	}
	global.GVA_DB = initialize.Gorm()
	err := initialize.InitAdminUser()
	if err != nil {
		return
	}
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	mainrouter := app.Group("/api/v1")
	AddRoutes(mainrouter)
	runPort := global.GVA_CONFIG.UmsInfo.RunPort
	appPort := fmt.Sprintf(":%d", runPort)
	if err := app.Run(appPort); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}

package main

import (
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
	initialize.InitConfig()
	global.GVA_DB = initialize.Gorm()
	app := gin.New()
	mainrouter := app.Group("/api/v1")
	AddRoutes(mainrouter)
	if err := app.Run(":5500"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}

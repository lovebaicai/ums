package router

import (
	"ums_backend/api"

	"github.com/gin-gonic/gin"
)

// type UserRouter struct{}

// 同步ldap账号写入数据库
func InitLdapRouter(Router *gin.RouterGroup) {
	ldapRouter := Router.Group("ldap")
	{
		ldapRouter.GET("sync", api.SyncLdap)
	}
}

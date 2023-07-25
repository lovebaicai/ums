package router

import (
	"ums_backend/api"

	"github.com/gin-gonic/gin"
)

// type ldapRouter struct{}

// 同步ldap账号写入数据库
func InitLdapRouter(Router *gin.RouterGroup) {
	ldapRouter := Router.Group("ldap")
	{
		ldapRouter.GET("sync", api.SyncLdap)
		ldapRouter.POST("getUserList", api.GetUserList)
		ldapRouter.POST("addLdapUser", api.AddLdapUser)
		ldapRouter.POST("changeUserStatus", api.ChangeUserStatus)
		ldapRouter.POST("getExistUser", api.GetExistUser)
		ldapRouter.POST("resetPassword", api.ResetPassword)
		ldapRouter.GET("getUserTotal", api.GetUserTotal)
	}
}

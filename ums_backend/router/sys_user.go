package router

import (
	"ums_backend/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("login", api.LoginUser)
		userRouter.GET("getUserInfo", api.GetUserInfo)
		userRouter.POST("logout", api.LogOut)
		userRouter.POST("getUserList", api.GetSysUserList)
		userRouter.POST("AddSysUser", api.AddSysUser)
		userRouter.POST("getExistUser", api.GetSysExistUser)
		userRouter.POST("ResetSysPassword", api.ResetSysPassword)
	}
}

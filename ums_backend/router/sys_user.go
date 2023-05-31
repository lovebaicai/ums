package router

import (
	"ums_backend/api"

	"github.com/gin-gonic/gin"
)

// type UserRouter struct{}

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("login", api.LoginUser)
		userRouter.POST("getUserList", api.GetUserList)
		userRouter.GET("getUserInfo", api.GetUserInfo)
		userRouter.POST("addUser", api.AddUser)
		userRouter.POST("changeUserStatus", api.ChangeUserStatus)
		userRouter.POST("getExistUser", api.GetExistUser)
		userRouter.POST("resetPassword", api.ResetPassword)
	}
}

package api

import (
	"ums_backend/global"
	system "ums_backend/model"
	"ums_backend/service"
	"ums_backend/utils"
	"ums_backend/utils/request"
	"ums_backend/utils/response"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var user request.Login
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		u := &system.SysUser{Username: user.Username, Password: user.Password}
		if user, err := service.Login(u); err != nil {
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			if user.Status != 1 {
				response.FailWithMessage("用户被禁止登录", c)
				return
			}
			TokenNext(c, *user)
		}
	}
}

func TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(request.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	} else {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt,
		}, "登录成功", c)
		return
	}
}

func GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	if ReqUser, err := service.GetUserInfo(uuid); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

func LogOut(c *gin.Context) {
	response.OkWithMessage("ok", c)
}

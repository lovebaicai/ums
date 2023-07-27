package api

import (
	"ums_backend/global"
	system "ums_backend/model"
	"ums_backend/service"
	"ums_backend/utils"
	"ums_backend/utils/request"
	"ums_backend/utils/response"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
			if user.Status != 2 {
				response.FailWithMessage("用户被禁止登录", c)
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

// AddUser
// @Description 增加系统用户
func AddSysUser(c *gin.Context) {
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{Username: r.Username, CNname: r.CNname, Email: r.Email}
	user.UUID = uuid.NewV4()
	user.Password = utils.BcryptHash("ums@123")
	_, err := service.AddSysUser(*user)
	if err != nil {
		response.FailWithMessage("系统用户增加失败", c)
	} else {
		response.OkWithDetailed(response.UserInfoResult{
			Username: r.Username,
			CNname:   r.CNname,
			Email:    r.Email,
		}, "系统用户增加成功", c)
	}
}

func GetSysExistUser(c *gin.Context) {
	var reqInfo request.UserName
	_ = c.ShouldBindJSON(&reqInfo)
	if err := service.GetSysExistUser(reqInfo.Username); err != nil {
		response.OkWithMessage("0", c)
	} else {
		response.OkWithMessage("1", c)
	}
}

func GetSysUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := service.GetSysUserList(pageInfo); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func ResetSysPassword(c *gin.Context) {
	var reqInfo request.UserName
	_ = c.ShouldBindJSON(&reqInfo)
	if err := service.ResetSysPassword(reqInfo.Username); err != nil {
		response.OkWithMessage("error", c)
	} else {
		response.OkWithMessage("ok", c)
	}
}

func ChangeSysUserStatus(c *gin.Context) {
	var reqInfo request.GetById
	_ = c.ShouldBindJSON(&reqInfo)
	if err := service.ChangeSysUserStatus(reqInfo.ID); err != nil {
		response.FailWithMessage("更新用户状态失败", c)
	} else {
		response.OkWithMessage("更新用户状态成功", c)
	}
}

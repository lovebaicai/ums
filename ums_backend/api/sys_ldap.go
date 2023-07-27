package api

import (
	system "ums_backend/model"
	"ums_backend/service"
	"ums_backend/utils"
	"ums_backend/utils/request"
	"ums_backend/utils/response"

	"github.com/gin-gonic/gin"
)

func SyncLdap(c *gin.Context) {
	// service.SyncLdap()
	if err := service.SyncLdap(); err != nil {
		response.OkWithMessage("同步用户失败", c)
	} else {
		response.OkWithMessage("同步用户成功", c)
	}
}

func ChangeUserStatus(c *gin.Context) {
	var reqInfo request.GetById
	_ = c.ShouldBindJSON(&reqInfo)
	// if err := utils.Verify(reqInfo, utils.IdVerify); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	if err := service.ChangeUserStatus(reqInfo.ID); err != nil {
		response.FailWithMessage("更新用户状态失败", c)
	} else {
		response.OkWithMessage("更新用户状态成功", c)
	}
}

// GetUserList
// @Description 获取用户列表
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := service.GetUserInfoList(pageInfo); err != nil {
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

// AddUser
// @Description 增加LDAP用户
func AddLdapUser(c *gin.Context) {
	var r request.LdapUser
	user := &system.LdapUser{Username: r.Username, CNname: r.CNname, Email: r.Email}
	_, err := service.AddLdapUser(*user, false)
	if err != nil {
		response.FailWithMessage("ldap用户增加失败", c)
	} else {
		response.OkWithDetailed(response.UserInfoResult{
			Username: r.Username,
			CNname:   r.CNname,
			Email:    r.Email,
		}, "ldap用户增加成功", c)
	}
}

func GetExistUser(c *gin.Context) {
	var reqInfo request.UserName
	_ = c.ShouldBindJSON(&reqInfo)
	if err := service.GetExistUser(reqInfo.Username); err != nil {
		response.OkWithMessage("0", c)
	} else {
		response.OkWithMessage("1", c)
	}
}

func ResetPassword(c *gin.Context) {
	var reqInfo request.UserName
	_ = c.ShouldBindJSON(&reqInfo)
	if err := service.ModifyPassword(reqInfo.Username); err != nil {
		response.OkWithMessage("error", c)
	} else {
		response.OkWithMessage("ok", c)
	}
}

func GetUserTotal(c *gin.Context) {
	info, err := service.GetUserTotal()
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.UserTotalResponse{
			Total:        info.Total,
			EnableTotal:  info.EnableTotal,
			DisableTotal: info.DisableTotal,
			GroupTotal:   info.GroupTotal,
		}, "获取成功", c)
	}
}

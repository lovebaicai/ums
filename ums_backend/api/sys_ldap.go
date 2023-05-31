package api

import (
	"ums_backend/service"
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

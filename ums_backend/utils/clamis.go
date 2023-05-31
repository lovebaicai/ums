package utils

import (
	"log"
	"ums_backend/global"
	"ums_backend/utils/request"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := c.Request.Header.Get("X-Token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		log.Println(err)
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
// func GetUserID(c *gin.Context) uint {

// 	if claims, exists := c.Get("claims"); !exists {
// 		if cl, err := GetClaims(c); err != nil {
// 			return 0
// 		} else {
// 			return cl.ID
// 		}
// 	} else {
// 		waitUse := claims.(*request.CustomClaims)
// 		return waitUse.ID
// 	}
// }

func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID
	}
}

package response

import (
	system "ums_backend/model"

	"github.com/golang-jwt/jwt/v4"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type UserInfoResult struct {
	Username string `json:"username"`
	CNname   string `json:"cnname"`
	Email    string `json:"Email"`
}

type LoginResponse struct {
	User      system.SysUser   `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt *jwt.NumericDate `json:"expiresAt"`
}

type UserTotalResponse struct {
	Total        int64 `json:"total"`
	GroupTotal   int64 `json:"grouptotal"`
	EnableTotal  int64 `json:"enabletotal"`
	DisableTotal int64 `json:"disabletotal"`
}

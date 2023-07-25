package system

import (
	"ums_backend/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.BasicModel
	UUID     uuid.UUID `json:"uuid" gorm:"not null;comment:用户UUID"`
	Username string    `json:"username" gorm:"ot null;comment:用户名"`
	Password string    `json:"_" gorm:"comment:用户密码"`
	CNname   string    `json:"cnname" gorm:"comment:用户中文名"`
	Email    string    `json:"email" grom:"comment:用户邮箱"`
	Status   int       `json:"status" gorm:"用户状态;default:2"`
	// ServiceID   string        `json:"serviceid" gorm:"commend:用户关联系统ID"`
	// ServiceType []ServiceType `json:"servicetype" gorm:"foreignKey:ServiceID;comment:关联系统"`
}

func (SysUser) TableName() string {
	return "sys_users"
}

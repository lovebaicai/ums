package system

import (
	"ums_backend/global"
)

type ServiceType struct {
	global.BasicModel
	ServiceID   uint   `json:"serviceid" gorm:"not null;unique;primary_key;comment:系统ID;size:90"`
	ServiceName string `json:"servicename" gorm:"comment:系统名称"`
}

func (ServiceType) TableName() string {
	return "service_types"
}

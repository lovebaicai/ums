package initialize

import (
	"errors"
	"log"

	"ums_backend/global"
	system "ums_backend/model"
	"ums_backend/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func InitAdminUser() (err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", "admin").First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		log.Println("admin已存在.")
		return
	} else {
		user.UUID = uuid.NewV4()
		user.Username = "admin"
		user.Password = utils.BcryptHash("ums@123")
		if err = global.GVA_DB.Create(&user).Error; err != nil {
			log.Println("admin创建成功.")
		}
	}
	return err
}

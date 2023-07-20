package service

import (
	"errors"

	"ums_backend/global"
	system "ums_backend/model"
	"ums_backend/utils"

	uuid "github.com/satori/go.uuid"
)

func Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	var user system.SysUser
	err = global.GVA_DB.Model(system.SysUser{}).Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

func GetUserInfo(UUID uuid.UUID) (user system.SysUser, err error) {
	var userInfo system.SysUser
	err = global.GVA_DB.First(&userInfo, "UUID = ?", UUID).Error
	if err != nil {
		return
	}
	return userInfo, err
}

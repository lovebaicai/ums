package service

import (
	"errors"

	"ums_backend/global"
	system "ums_backend/model"
	"ums_backend/utils"
	"ums_backend/utils/request"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
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
	err = global.GVA_DB.Model(&userInfo).First(&userInfo, "UUID = ?", UUID).Error
	if err != nil {
		return
	}
	return userInfo, err
}

func AddSysUser(sysInfo system.SysUser) (userInter system.SysUser, err error) {
	err = global.GVA_DB.Model(&sysInfo).Create(&sysInfo).Error
	return sysInfo, err
}

func GetSysUserList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// 忽略admin
	err = db.Limit(limit).Offset(offset).Not("username = ?", "admin").Find(&userList).Error
	return userList, total, err
}

func GetSysExistUser(username string) (err error) {
	var user system.SysUser
	var users []system.SysUser
	if !errors.Is(global.GVA_DB.Model(&user).Where("username = ?", username).First(&users).Error, gorm.ErrRecordNotFound) { // 判断用户名是否存在
		return errors.New("用户名已存在")
	}
	return
}

func ResetSysPassword(username string) (err error) {
	var user system.SysUser
	user.Password = utils.BcryptHash("ums@123")
	global.GVA_DB.Model(&user).Where("username = ?", username).Update("password", user.Password)
	if err != nil {
		return err
	}
	return
}

func ChangeSysUserStatus(id int) (err error) {
	var userInfo system.SysUser
	var status int
	if userInfo.Status == 1 {
		status = 2
	} else {
		status = 1
	}
	err = global.GVA_DB.Model(&userInfo).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return
}

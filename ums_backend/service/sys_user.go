package service

import (
	"errors"

	"ums_backend/global"
	"ums_backend/model/system"
	"ums_backend/utils/request"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	var user system.SysUser
	err = global.GVA_DB.Model(system.SysUser{}).Where("username = ?", u.Username).First(&user).Error
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

func GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Not("username = ?", "admin").Find(&userList).Error
	return userList, total, err
}

func GetExistUser(username string) (err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册")
	}
	return
}

func AddUser(u system.SysUser, ldad_sync bool) (userInter system.SysUser, err error) {
	// 判断增加用户类型
	if !ldad_sync {
		// 如果不是同步ldap数据, 则查找最后一个增加的ldap用户, 下一个用户的ldapid + 1
		var userInfo system.SysUser
		_ = global.GVA_DB.Last(&userInfo)
		AddLDAPUser(u.Username, u.Email, userInfo.LADPUID+1)
		u.LADPUID = userInfo.LADPUID + 1
	}
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

func ChangeUserStatus(id int) (err error) {
	var userInfo system.SysUser
	var status int
	err = global.GVA_DB.First(&userInfo, "id = ?", id).Error
	if err != nil {
		return
	}
	if userInfo.Status == 0 {
		status = 1
	} else {
		status = 0
	}
	if err := ChangeLDAPStatus(userInfo.Username, status); err != nil {
		return err
	}
	err = global.GVA_DB.Model(system.SysUser{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return
}

func GetUserTotal() (total int64, err error) {
	// var userInfo system.SysUser
	err = global.GVA_DB.Model(system.SysUser{}).Where("status = ?", 1).Count(&total).Error
	if err != nil {
		return
	}
	return
}

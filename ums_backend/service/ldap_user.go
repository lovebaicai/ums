package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"ums_backend/global"
	"ums_backend/initialize"
	system "ums_backend/model"
	"ums_backend/utils/request"

	ldap "github.com/go-ldap/ldap/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CheckUserStatus(ldapuid int) (status int) {
	user_active := 1
	ld := initialize.ConnectLDAP()
	filter := fmt.Sprintf("(uidNumber=%d)", ldapuid)
	searchRequest := ldap.NewSearchRequest(
		"ou=Disable,dc=greentown-tech,dc=com", // OU的DN
		ldap.ScopeWholeSubtree,                // 搜索整个子树
		ldap.NeverDerefAliases,                // 不解析别名
		0,                                     // 没有限制条目数
		0,                                     // 不超时
		false,                                 // 不返回属性值
		filter,                                // 过滤器
		nil,                                   // 返回所有属性
		nil,
	)
	searchResult, err := ld.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	if len(searchResult.Entries) == 0 {
		user_active = 2
	}
	return user_active
}

func SyncLdap() (err error) {
	ld := initialize.ConnectLDAP()
	ldapgroup := global.GVA_CONFIG.LdapInfo.LGroup
	defer ld.Close()
	searchRequest := ldap.NewSearchRequest(
		ldapgroup,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))",
		[]string{"dn", "cn", "gidNumber", "displayName"},
		nil,
	)
	search, err := ld.Search(searchRequest)
	if err != nil {
		fmt.Println("ldap查找失败", err)
	}
	var user system.LdapUser
	var users []system.LdapUser
	for _, entry := range search.Entries {
		ldapuid, _ := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
		user_active := CheckUserStatus(ldapuid)
		// 查询用户是否存在, 若存在则判断状态是否更新
		result := global.GVA_DB.Model(&user).Where("ladp_uid = ?", ldapuid).First(&users)
		if result.RowsAffected == 1 {
			global.GVA_DB.Model(&user).Where("ladp_uid = ?", ldapuid).Update("status", user_active)
			continue
		}
		userinfo := &system.LdapUser{Username: entry.GetAttributeValue("cn"), CNname: entry.GetAttributeValue("displayName"), LADPUID: ldapuid, Status: user_active}
		_, err = AddLdapUser(*userinfo, true)
		if err != nil {
			return err
		}
	}
	return
}

func AddLdapUser(ldapInfo system.LdapUser, ldadSync bool) (userInter system.LdapUser, err error) {
	// 判断增加用户类型
	if !ldadSync {
		// 如果不是同步ldap数据, 则查找最后一个增加的ldap用户, 下一个用户的ldapid + 1
		_ = global.GVA_DB.Last(&ldapInfo)
		LdapAdd(ldapInfo.Username, ldapInfo.Email, ldapInfo.LADPUID+1)
		ldapInfo.LADPUID = ldapInfo.LADPUID + 1
	}
	// u.UUID = uuid.NewV4()
	err = global.GVA_DB.Model(&ldapInfo).Create(&ldapInfo).Error
	return ldapInfo, err
}

func LdapAdd(username, email string, ldapid int) {
	ldapgroup := global.GVA_CONFIG.LdapInfo.LGroup
	ld := initialize.ConnectLDAP()
	defer ld.Close()
	addResponse := ldap.NewAddRequest(fmt.Sprintf("uid=%s,ou=People,%s", username, ldapgroup), []ldap.Control{})
	addResponse.Attribute("cn", []string{username})
	addResponse.Attribute("sn", []string{username})
	addResponse.Attribute("uid", []string{username})
	addResponse.Attribute("homeDirectory", []string{fmt.Sprintf("/home/%s", username)})
	addResponse.Attribute("mail", []string{email})
	addResponse.Attribute("loginShell", []string{"/bin/bash"})
	addResponse.Attribute("gidNumber", []string{fmt.Sprintf("%d", ldapid)})
	addResponse.Attribute("uidNumber", []string{fmt.Sprintf("%d", ldapid)})
	addResponse.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "top", "inetOrgPerson"})
	if err := ld.Add(addResponse); err != nil {
		log.Fatal("error adding service:", addResponse, err)
	}
	if err := ModifyPassword(username); err != nil {
		log.Fatal("密码修改失败:", err)
		return
	}
}

func ModifyPassword(username string) (err error) {
	ldapgroup := global.GVA_CONFIG.LdapInfo.LGroup
	ld := initialize.ConnectLDAP()
	defer ld.Close()
	passwordModifyRequest2 := ldap.NewPasswordModifyRequest(fmt.Sprintf("uid=%s,%s,dc=com", username, ldapgroup), "", "passwd@123")
	_, err = ld.PasswordModify(passwordModifyRequest2)
	if err != nil {
		return err
	}
	return
}

func ChangeLDAPStatus(username string, status int) (err error) {
	ldapgroup := global.GVA_CONFIG.LdapInfo.LGroup
	ld := initialize.ConnectLDAP()
	defer ld.Close()
	var old_rn, new_rn string
	if status == 2 {
		old_rn = "Disable"
		new_rn = "People"
	} else {
		old_rn = "People"
		new_rn = "Disable"
	}
	changeResponse := ldap.NewModifyDNRequest(
		fmt.Sprintf("uid=%s,ou=%s,%s", username, old_rn, ldapgroup), fmt.Sprintf("uid=%s", username), true, fmt.Sprintf("ou=%s,%s", new_rn, ldapgroup))
	if err := ld.ModifyDN(changeResponse); err != nil {
		global.GVA_LOG.Error("Failed to modify userDN: %s\n", zap.Error(err))
	}
	return
}

func GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.LdapUser{})
	var userList []system.LdapUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Not("username = ?", "admin").Find(&userList).Error
	return userList, total, err
}

func GetExistUser(username string) (err error) {
	var user system.LdapUser
	var users []system.LdapUser
	if !errors.Is(global.GVA_DB.Model(&user).Where("username = ?", username).First(&users).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册")
	}
	return
}

func ChangeUserStatus(id int) (err error) {
	var userInfo system.LdapUser
	var status int
	err = global.GVA_DB.First(&userInfo, "id = ?", id).Error
	if err != nil {
		return
	}
	if userInfo.Status == 1 {
		status = 2
	} else {
		status = 1
	}
	if err := ChangeLDAPStatus(userInfo.Username, status); err != nil {
		return err
	}
	err = global.GVA_DB.Model(&userInfo).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return
}

func GetUserTotal() (info request.UserTotal, err error) {
	err = global.GVA_DB.Model(system.LdapUser{}).Count(&info.Total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(system.LdapUser{}).Where("status = ?", 1).Count(&info.EnableTotal).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(system.LdapUser{}).Not("status = ?", 1).Count(&info.DisableTotal).Error
	if err != nil {
		return
	}
	info.GroupTotal = 2
	return info, err
}

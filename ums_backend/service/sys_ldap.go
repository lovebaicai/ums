package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"ums_backend/global"
	"ums_backend/initialize"
	"ums_backend/model/system"

	ldap "github.com/go-ldap/ldap/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

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
	for _, entry := range search.Entries {
		ldapuid, _ := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
		var user system.SysUser
		if !errors.Is(global.GVA_DB.Where("username = ?", entry.GetAttributeValue("cn")).First(&user).Error, gorm.ErrRecordNotFound) {
			// log.Println("用户已存在")
			continue
		}
		userinfo := &system.SysUser{Username: entry.GetAttributeValue("cn"), CNname: entry.GetAttributeValue("displayName"), LADPUID: ldapuid}
		_, err = AddUser(*userinfo, true)
		if err != nil {
			return
		}
	}
	return
}

func AddLDAPUser(username, email string, ldapid int) {
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
	passwordModifyRequest2 := ldap.NewPasswordModifyRequest(fmt.Sprintf("uid=%s,%s,dc=com", username, ldapgroup), "", "Uama@123")
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
	if status == 1 {
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

package initialize

import (
	"fmt"

	"ums_backend/global"

	ldap "github.com/go-ldap/ldap/v3"
)

func ConnectLDAP() (LD *ldap.Conn) {
	ldapuser := global.GVA_CONFIG.LdapInfo.LUser
	ldappass := global.GVA_CONFIG.LdapInfo.LPasswd
	ldhost := global.GVA_CONFIG.LdapInfo.LHost
	ldport := global.GVA_CONFIG.LdapInfo.LPort
	LD, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldhost, ldport))
	if err != nil {
		fmt.Println("ldap连接失败", err)
	}
	err = LD.Bind(ldapuser, ldappass)
	if err != nil {
		fmt.Println("管理员认证失败", err)
	}
	return LD
}

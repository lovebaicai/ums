package config

type Server struct {
	Name      string      `mapstructure:"name" json:"name" yaml:"name"`
	JWT       JWT         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysqlinfo MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	LdapInfo  LdapConfig  `mapstructure:"ldap" json:"ldap" yaml:"ldap"`
	UmsInfo   UmsConfig   `mapstructure:"ums" json:"ums" yaml:"ums"`
}

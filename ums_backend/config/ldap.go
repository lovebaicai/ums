package config

type LdapConfig struct {
	LHost   string `mapstructure:"ldaphost"`
	LPort   int    `mapstructure:"ldapport"`
	LUser   string `mapstructure:"ldapuser"`
	LPasswd string `mapstructure:"ldappass"`
	LGroup  string `mapstructure:"ldapgroup"`
}

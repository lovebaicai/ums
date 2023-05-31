package config

type Server struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	JWT  JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	// Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	// System  System  `mapstructure:"system" json:"system" yaml:"system"`
	// Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// // auto
	// AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysqlinfo MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	LdapInfo  LdapConfig  `mapstructure:"ldap" json:"ldap" yaml:"ldap"`
	// Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	// DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	// Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	// Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// // 跨域配置
	// Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}

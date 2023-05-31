package config

type MysqlConfig struct {
	DBHost     string `mapstructure:"dbhost"`
	DBPort     int    `mapstructure:"dbport"`
	DBUser     string `mapstructure:"dbuser"`
	DBPassword string `mapstructure:"dbpassword"`
	DBName     string `mapstructure:"dbname"`
}

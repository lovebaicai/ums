package initialize

import (
	"gorm.io/gorm"
	"log"
	"os"
	system2 "ums_backend/model/system"
)

func Gorm() *gorm.DB {
	return ConnectDB()
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system2.SysUser{},
		system2.ServiceType{},
	)
	if err != nil {
		log.Println("register table failed")
		os.Exit(0)
	}
	log.Println("register table success")
}

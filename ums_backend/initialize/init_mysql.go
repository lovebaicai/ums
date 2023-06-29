package initialize

import (
	"fmt"
	"log"
	"os"

	"ums_backend/global"
	system2 "ums_backend/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func ConnectDB() (db *gorm.DB) {
	dbUser := global.GVA_CONFIG.Mysqlinfo.DBUser
	dbPass := global.GVA_CONFIG.Mysqlinfo.DBPassword
	dbHost := global.GVA_CONFIG.Mysqlinfo.DBHost
	dbPort := global.GVA_CONFIG.Mysqlinfo.DBPort
	dbName := global.GVA_CONFIG.Mysqlinfo.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	RegisterTables(db)
	return db
}

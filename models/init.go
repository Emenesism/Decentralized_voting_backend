package models

import (
    "fmt"
	"github.com/emenesism/Decentralized-voting-backend/config"
    "github.com/charmbracelet/log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error 

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	config.AppConfig.DB_User,
	config.AppConfig.DB_Passwd,
	config.AppConfig.DB_Host,
	config.AppConfig.DB_Port,
	config.AppConfig.DB_Name,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database", "error", err.Error())
	}

	DB.AutoMigrate(&User{})

	log.Info("Database connected and migrated successfully")
}




package main

import (
	"log"
	"otp-authentication-system/database"
	"otp-authentication-system/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.ConnectToDatabase()
}

func main() {
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Could not get database connection: ", err)
		}
		sqlDB.Close()
	}()

	if err := db.AutoMigrate(&models.User{}, &models.OTPSession{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}

package database

import (
	"log-me-in/model"
	"log-me-in/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB

func GetConnection() *gorm.DB {
	if instance == nil {
		dsn, err := utils.GetFromEnv("DB_DSN")
		if err != nil {
			panic("Error making connection to database")
		}
		db, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = db
	}
	return instance
}

func AutoMigrate() {
	db := GetConnection()
	db.AutoMigrate(&model.User{}, &model.Note{})
}

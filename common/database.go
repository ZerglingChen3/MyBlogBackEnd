package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"myBlog/model"
)

var DB *gorm.DB

func InitDB() {
	database := "blog.db"

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("Create Migrate User error, err: " + err.Error())
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

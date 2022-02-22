package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDb() {
	database, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Url{})
	DB = database
}

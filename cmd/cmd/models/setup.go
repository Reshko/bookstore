package models

import (
	_ "github.com/miekg/pkcs11"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}

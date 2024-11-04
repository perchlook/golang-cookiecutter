package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database_url string) *gorm.DB {
	log.Println("Connecting to database", database_url)
	d, err := gorm.Open(sqlite.Open(database_url), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect database", database_url)
	}

	return d
}
package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetDatabaseConnection() *gorm.DB {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	return db
}

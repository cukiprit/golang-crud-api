package config

import (
	"learn/api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Print("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Todo{})

	return db, nil
}

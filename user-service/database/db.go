package database

import (
	"log"
	"user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "host=localhost user=postgres password=123 dbname=quickbook port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Successfully connected to the database!")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}

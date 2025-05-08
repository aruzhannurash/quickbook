package main

import (
	"user-service/handlers"
	"user-service/models"
	"user-service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=postgres user=postgres password=123 dbname=quickbook port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{})

	h := &handlers.Handler{DB: db}
	r := gin.Default()
	routes.RegisterUserRoutes(r, h)

	r.Run(":8081")
}

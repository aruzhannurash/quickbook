package main

import (
	"log"

	"github.com/aruzhannurash/quickbook/handlers"
	"github.com/aruzhannurash/quickbook/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=quickbook port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	db.AutoMigrate(&models.Specialist{}, &models.Client{}, &models.Appointment{})

	h := &handlers.Handler{DB: db}
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/specialists", h.GetSpecialists)
	api.GET("/specialists/:id", h.GetSpecialistByID)
	api.POST("/specialists", h.CreateSpecialist)
	api.PUT("/specialists/:id", h.UpdateSpecialist)
	api.DELETE("/specialists/:id", h.DeleteSpecialist)

	api.GET("/clients", h.GetClients)
	api.GET("/clients/:id", h.GetClientByID)
	api.POST("/clients", h.CreateClient)
	api.PUT("/clients/:id", h.UpdateClient)
	api.DELETE("/clients/:id", h.DeleteClient)

	api.GET("/appointments", h.GetAppointments)
	api.GET("/appointments/:id", h.GetAppointmentByID)
	api.POST("/appointments", h.CreateAppointment)
	api.PUT("/appointments/:id", h.UpdateAppointment)
	api.DELETE("/appointments/:id", h.DeleteAppointment)

	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}

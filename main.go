package main

import (
	"log"

	"github.com/aruzhannurash/quickbook/handlers"
	"github.com/aruzhannurash/quickbook/middlewares"

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

	h := &handlers.Handler{DB: db}
	r := gin.Default()

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	auth := r.Group("/api")
	auth.Use(middlewares.JWTAuthMiddleware())

	auth.GET("/specialists", h.GetSpecialists)
	auth.GET("/specialists/:id", h.GetSpecialistByID)
	auth.POST("/specialists", h.CreateSpecialist)
	auth.PUT("/specialists/:id", h.UpdateSpecialist)
	auth.DELETE("/specialists/:id", h.DeleteSpecialist)

	auth.GET("/clients", h.GetClients)
	auth.GET("/clients/:id", h.GetClientByID)
	auth.POST("/clients", h.CreateClient)
	auth.PUT("/clients/:id", h.UpdateClient)
	auth.DELETE("/clients/:id", h.DeleteClient)

	auth.GET("/appointments", h.GetAppointments)
	auth.GET("/appointments/:id", h.GetAppointmentByID)
	auth.POST("/appointments", h.CreateAppointment)
	auth.PUT("/appointments/:id", h.UpdateAppointment)
	auth.DELETE("/appointments/:id", h.DeleteAppointment)
	auth.GET("/appointments/client/:client_id", h.GetAppointmentsByClient)
	auth.GET("/appointments/specialist/:specialist_id", h.GetAppointmentsBySpecialist)

	auth.POST("/reviews", h.CreateReview)
	auth.GET("/reviews/specialist/:specialist_id", h.GetReviewsBySpecialist)
	auth.GET("/reviews/client/:client_id", h.GetReviewsByClient)

	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}

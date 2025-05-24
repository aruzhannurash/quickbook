package main

import (
	"appointments-service/database"
	"appointments-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	h := handlers.Handler{DB: db}

	r := gin.Default()

	r.GET("/appointments", h.GetAppointments)
	r.POST("/appointments", h.CreateAppointment)

	r.Run(":8082")
}

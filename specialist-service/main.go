package main

import (
	"specialist-service/handlers"
	"specialist-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	h := handlers.NewHandler()
	routes.RegisterSpecialistRoutes(r, h)

	r.Run(":8081")
}

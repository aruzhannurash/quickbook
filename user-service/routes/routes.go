package routes

import (
	"user-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, h *handlers.Handler) {
	r.GET("/users/:id", h.GetUserByID)
}

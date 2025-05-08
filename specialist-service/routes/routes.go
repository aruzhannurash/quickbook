package routes

import (
	"specialist-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSpecialistRoutes(r *gin.Engine, h *handlers.Handler) {
	r.GET("/specialists/:id", h.GetUserFromUserService)
}

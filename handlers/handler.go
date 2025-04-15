package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetSpecialists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get specialists"})
}
func (h *Handler) GetSpecialistByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get specialist by ID"})
}
func (h *Handler) CreateSpecialist(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create specialist"})
}
func (h *Handler) UpdateSpecialist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update specialist"})
}
func (h *Handler) DeleteSpecialist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete specialist"})
}

func (h *Handler) GetClients(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Get clients"}) }
func (h *Handler) GetClientByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get client by ID"})
}
func (h *Handler) CreateClient(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create client"})
}
func (h *Handler) UpdateClient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update client"})
}
func (h *Handler) DeleteClient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete client"})
}

func (h *Handler) GetAppointments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get appointments"})
}
func (h *Handler) GetAppointmentByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get appointment by ID"})
}
func (h *Handler) CreateAppointment(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create appointment"})
}
func (h *Handler) UpdateAppointment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update appointment"})
}
func (h *Handler) DeleteAppointment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete appointment"})
}

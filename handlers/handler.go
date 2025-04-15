package handlers

import (
	"fmt"
	"net/http"

	"github.com/aruzhannurash/quickbook/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetSpecialists(c *gin.Context) {
	var specialists []models.Specialist
	if err := h.DB.Find(&specialists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch specialists"})
		return
	}
	c.JSON(http.StatusOK, specialists)
}

func (h *Handler) GetSpecialistByID(c *gin.Context) {
	id := c.Param("id")
	var specialist models.Specialist
	if err := h.DB.First(&specialist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Specialist not found"})
		return
	}
	c.JSON(http.StatusOK, specialist)
}

func (h *Handler) CreateSpecialist(c *gin.Context) {
	fmt.Println("🟢 Новый хендлер вызван!") // 👈
	var specialist models.Specialist
	if err := c.ShouldBindJSON(&specialist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&specialist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create specialist"})
		return
	}
	c.JSON(http.StatusCreated, specialist)
}

func (h *Handler) UpdateSpecialist(c *gin.Context) {
	id := c.Param("id")
	var specialist models.Specialist
	if err := h.DB.First(&specialist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Specialist not found"})
		return
	}
	if err := c.ShouldBindJSON(&specialist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&specialist)
	c.JSON(http.StatusOK, specialist)
}

func (h *Handler) DeleteSpecialist(c *gin.Context) {
	id := c.Param("id")
	var specialist models.Specialist
	if err := h.DB.First(&specialist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Specialist not found"})
		return
	}
	h.DB.Delete(&specialist)
	c.JSON(http.StatusOK, gin.H{"message": "Specialist deleted"})
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []models.Client
	if err := h.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (h *Handler) GetClientByID(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := h.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *Handler) CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&client)
	c.JSON(http.StatusCreated, client)
}

func (h *Handler) UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := h.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&client)
	c.JSON(http.StatusOK, client)
}

func (h *Handler) DeleteClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := h.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	h.DB.Delete(&client)
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
}

func (h *Handler) GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	if err := h.DB.Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func (h *Handler) GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment
	if err := h.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}
	c.JSON(http.StatusOK, appointment)
}

func (h *Handler) CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&appointment)
	c.JSON(http.StatusCreated, appointment)
}

func (h *Handler) UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment
	if err := h.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&appointment)
	c.JSON(http.StatusOK, appointment)
}

func (h *Handler) DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment
	if err := h.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}
	h.DB.Delete(&appointment)
	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
}

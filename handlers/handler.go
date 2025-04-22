package handlers

import (
	"fmt"
	"log"
	"net/http"

	"strconv"
	"strings"

	"time"

	"github.com/aruzhannurash/quickbook/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := h.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hash)

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *Handler) GetSpecialists(c *gin.Context) {
	var specialists []models.Specialist

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	name := c.Query("name")
	position := c.Query("position")

	pageInt := 1
	limitInt := 10
	fmt.Sscanf(page, "%d", &pageInt)
	fmt.Sscanf(limit, "%d", &limitInt)
	offset := (pageInt - 1) * limitInt

	query := h.DB.Model(&models.Specialist{})

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if position != "" {
		query = query.Where("position ILIKE ?", "%"+position+"%")
	}

	var total int64
	query.Count(&total)

	if err := query.Offset(offset).Limit(limitInt).Find(&specialists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch specialists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":        pageInt,
		"limit":       limitInt,
		"total":       total,
		"specialists": specialists,
	})
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
	fmt.Println("Новый хендлер вызван!")
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

	name := c.Query("name")
	phone := c.Query("phone")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query := h.DB.Model(&models.Client{})

	if name != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%")
	}

	if phone != "" {
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}

	if err := query.Offset(offset).Limit(limit).Find(&clients).Error; err != nil {
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
func (h *Handler) CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if review.CreatedAt.IsZero() {
		review.CreatedAt = time.Now()
	}
	if review.UpdatedAt.IsZero() {
		review.UpdatedAt = time.Now()
	}

	log.Printf("Creating review: %+v", review)

	if err := h.DB.Create(&review).Error; err != nil {
		log.Printf("Error creating review: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}
	c.JSON(http.StatusCreated, review)
}

func (h *Handler) GetReviewsBySpecialist(c *gin.Context) {
	specialistID := c.Param("specialist_id")
	var reviews []models.Review

	if err := h.DB.Where("specialist_id = ?", specialistID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *Handler) GetReviewsByClient(c *gin.Context) {
	clientID := c.Param("client_id")
	var reviews []models.Review

	if err := h.DB.Where("client_id = ?", clientID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
func (h *Handler) GetAppointmentsByClient(c *gin.Context) {
	clientID := c.Param("client_id")
	var appointments []models.Appointment

	if err := h.DB.Where("client_id = ?", clientID).Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}
func (h *Handler) GetAppointmentsBySpecialist(c *gin.Context) {
	specialistID := c.Param("specialist_id")
	var appointments []models.Appointment

	if err := h.DB.Where("specialist_id = ?", specialistID).Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

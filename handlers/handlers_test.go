package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/aruzhannurash/quickbook/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockHandler struct {
	*Handler
	specialists []models.Specialist
}

func (h *mockHandler) GetSpecialists(c *gin.Context) {
	c.JSON(http.StatusOK, h.specialists)
}

func TestGetSpecialists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockSpecialists := []models.Specialist{
		{ID: 1, Name: "Dr. John", Email: "john@example.com", Position: "Dentist"},
		{ID: 2, Name: "Dr. Jane", Email: "jane@example.com", Position: "Therapist"},
	}

	mock := &mockHandler{specialists: mockSpecialists}
	router.GET("/specialists", mock.GetSpecialists)

	req, _ := http.NewRequest(http.MethodGet, "/specialists", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got []models.Specialist
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(got))
	assert.Equal(t, "Dr. John", got[0].Name)
}

func TestGetSpecialistByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockSpecialists := []models.Specialist{
		{ID: 1, Name: "Dr. John", Email: "john@example.com", Position: "Dentist"},
	}

	mock := &mockHandler{specialists: mockSpecialists}
	router.GET("/specialists/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, s := range mock.specialists {
			if strconv.Itoa(int(s.ID)) == id {
				c.JSON(http.StatusOK, s)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Specialist not found"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/specialists/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Dr. John")
}

func TestGetSpecialistByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockSpecialists := []models.Specialist{
		{ID: 1, Name: "Dr. John", Email: "john@example.com", Position: "Dentist"},
	}

	mock := &mockHandler{specialists: mockSpecialists}
	router.GET("/specialists/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, s := range mock.specialists {
			if strconv.Itoa(int(s.ID)) == id {
				c.JSON(http.StatusOK, s)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Specialist not found"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/specialists/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateSpecialist_InvalidData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mock := &mockHandler{specialists: []models.Specialist{}}
	router.POST("/specialists", func(c *gin.Context) {
		var newSpecialist models.Specialist
		if err := c.ShouldBindJSON(&newSpecialist); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		if newSpecialist.Name == "" || newSpecialist.Email == "" || newSpecialist.Position == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}
		mock.specialists = append(mock.specialists, newSpecialist)
		c.JSON(http.StatusCreated, newSpecialist)
	})

	req, _ := http.NewRequest(http.MethodPost, "/specialists", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Missing required fields")
}

func TestUpdateSpecialist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockSpecialists := []models.Specialist{
		{ID: 1, Name: "Dr. John", Email: "john@example.com", Position: "Dentist"},
	}

	mock := &mockHandler{specialists: mockSpecialists}
	router.PUT("/specialists/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedSpecialist models.Specialist
		if err := c.ShouldBindJSON(&updatedSpecialist); err == nil {
			for i, s := range mock.specialists {
				if strconv.Itoa(int(s.ID)) == id {
					mock.specialists[i] = updatedSpecialist
					c.JSON(http.StatusOK, updatedSpecialist)
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Specialist not found"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		}
	})

	updatedSpecialist := models.Specialist{Name: "Dr. John Updated", Email: "john_updated@example.com", Position: "Updated Dentist"}
	jsonData, _ := json.Marshal(updatedSpecialist)
	req, _ := http.NewRequest(http.MethodPut, "/specialists/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Dr. John Updated")
}

func TestDeleteSpecialist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockSpecialists := []models.Specialist{
		{ID: 1, Name: "Dr. John", Email: "john@example.com", Position: "Dentist"},
	}

	mock := &mockHandler{specialists: mockSpecialists}
	router.DELETE("/specialists/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, s := range mock.specialists {
			if strconv.Itoa(int(s.ID)) == id {
				mock.specialists = append(mock.specialists[:i], mock.specialists[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Specialist deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Specialist not found"})
	})

	req, _ := http.NewRequest(http.MethodDelete, "/specialists/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Specialist deleted")
}

func TestGetSpecialists_Empty(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mock := &mockHandler{specialists: []models.Specialist{}}
	router.GET("/specialists", mock.GetSpecialists)

	req, _ := http.NewRequest(http.MethodGet, "/specialists", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got []models.Specialist
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(got))
}

func TestInvalidJSONFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mock := &mockHandler{specialists: []models.Specialist{}}
	router.POST("/specialists", func(c *gin.Context) {
		var newSpecialist models.Specialist
		if err := c.ShouldBindJSON(&newSpecialist); err == nil {
			mock.specialists = append(mock.specialists, newSpecialist)
			c.JSON(http.StatusCreated, newSpecialist)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		}
	})

	req, _ := http.NewRequest(http.MethodPost, "/specialists", bytes.NewBuffer([]byte(`{"Name": "Dr. Alex", "Email": "alex@example.com"`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid data")
}
func TestCreateSpecialist_MissingRequiredFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mock := &mockHandler{specialists: []models.Specialist{}}
	router.POST("/specialists", func(c *gin.Context) {
		var newSpecialist models.Specialist
		if err := c.ShouldBindJSON(&newSpecialist); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		if newSpecialist.Name == "" || newSpecialist.Email == "" || newSpecialist.Position == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}
		mock.specialists = append(mock.specialists, newSpecialist)
		c.JSON(http.StatusCreated, newSpecialist)
	})

	req, _ := http.NewRequest(http.MethodPost, "/specialists", bytes.NewBuffer([]byte(`{"Name": "Dr. Alex"}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Missing required fields")
}

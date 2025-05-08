package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Handler struct {
	RestyClient *resty.Client
}

func NewHandler() *Handler {
	return &Handler{
		RestyClient: resty.New(),
	}
}

func (h *Handler) GetUserFromUserService(c *gin.Context) {
	userID := c.Param("id")
	resp, err := h.RestyClient.R().
		SetResult(map[string]interface{}{}).
		Get("http://localhost:8081/users/" + userID)

	if err != nil || resp.IsError() {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, resp.Result())
}

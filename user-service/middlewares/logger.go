package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		c.Next()

		duration := time.Since(start)
		log.Printf("[RequestID: %s] %s %s | Status: %d | Duration: %v | Time: %s",
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
			start.Format(time.RFC3339),
		)
	}
}

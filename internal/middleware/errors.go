// internal/middleware/errors.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
	}
}

func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
	}
}

// ErrorFormatter turns accumulated errors into a single 500 if not already written.
func ErrorFormatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 && !c.Writer.Written() {
			// In production, avoid leaking internals. Log c.Errors for observability.
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
	}
}

// internal/server/server.go
package server

import (
	"northwind-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	r := gin.New()
	// Logging + Recovery
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Request ID, CORS (optional)
	r.Use(middleware.RequestID()) // if you add it
	// r.Use(cors.Default())        // gin-contrib/cors if needed

	// Global 404/405 and error formatter
	// r.NoRoute(middleware.NotFoundHandler())
	r.NoMethod(middleware.MethodNotAllowedHandler())
	r.Use(middleware.ErrorFormatter())

	return r
}

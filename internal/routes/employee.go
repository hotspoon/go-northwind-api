// internal/routes/employees.go
package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(rg *gin.RouterGroup, h *handlers.EmployeeHandler) {
	employees := rg.Group("/employees")
	{
		employees.GET("", h.GetAll)
		employees.GET("/:id", h.GetOne)
		employees.POST("", h.Create)
		employees.PUT("/:id", h.Update)
		employees.DELETE("/:id", h.Delete)
	}
}

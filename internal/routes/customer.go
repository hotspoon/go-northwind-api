// internal/routes/customers.go
package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(rg *gin.RouterGroup, h *handlers.CustomerHandler) {
	customers := rg.Group("/customers")
	{
		customers.GET("", h.GetAll)
		customers.GET("/:id", h.GetOne)
		customers.POST("", h.Create)
		customers.PUT("/:id", h.Update)
		customers.DELETE("/:id", h.Delete)
	}
}

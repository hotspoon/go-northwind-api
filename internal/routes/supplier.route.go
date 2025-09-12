package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSupplierRoutes(rg *gin.RouterGroup, h *handlers.SupplierHandler) {
	categories := rg.Group("/suppliers")
	{
		categories.GET("", h.GetAll)
		categories.GET("/:id", h.GetOne)
		categories.POST("", h.Create)
		categories.PUT("/:id", h.Update)
		categories.DELETE("/:id", h.Delete)
	}
}

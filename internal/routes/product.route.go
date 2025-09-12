package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, h *handlers.ProductHandler) {
	products := rg.Group("/products")
	{
		products.GET("", h.GetAll)
		products.GET("/:id", h.GetOne)
		products.POST("", h.Create)
		products.PUT("/:id", h.Update)
		products.DELETE("/:id", h.Delete)
		products.GET("/:id/supplier", h.GetSupplier)
		products.GET("/:id/category", h.GetCategory)
	}
}

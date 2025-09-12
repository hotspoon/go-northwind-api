package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, h *handlers.CategoryHandler) {
	categories := rg.Group("/categories")
	{
		categories.GET("", h.GetAll)
		categories.GET("/:id", h.GetOne)
		categories.POST("", h.Create)
		categories.PUT("/:id", h.Update)
		categories.DELETE("/:id", h.Delete)
	}
}

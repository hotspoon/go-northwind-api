package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(rg *gin.RouterGroup, h *handlers.OrderHandler) {
	orders := rg.Group("/orders")
	{
		orders.GET("", h.GetAll)
		orders.GET("/paginated", h.GetPaginated) // <-- Add this line
		orders.GET("/:id", h.GetOne)
		orders.POST("", h.Create)
		orders.PUT("/:id", h.Update)
		orders.DELETE("/:id", h.Delete)
		orders.GET("/:id/details", h.GetOrderDetails)
	}
}

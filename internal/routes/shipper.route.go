package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterShipperRoutes(rg *gin.RouterGroup, h *handlers.ShipperHandler) {
	shippers := rg.Group("/shippers")
	{
		shippers.GET("", h.GetAll)
		shippers.GET("/:id", h.GetOne)
		shippers.POST("", h.Create)
		shippers.PUT("/:id", h.Update)
		shippers.DELETE("/:id", h.Delete)
	}
}

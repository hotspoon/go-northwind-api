package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRegionRoutes(rg *gin.RouterGroup, h *handlers.RegionHandler) {
	regions := rg.Group("/regions")
	{
		regions.GET("", h.GetAll)
		regions.GET("/:id", h.GetOne)
	}
}

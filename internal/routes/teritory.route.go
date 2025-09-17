package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTeritoryRoutes(rg *gin.RouterGroup, h *handlers.RegionHandler) {
	teritories := rg.Group("/territories")
	{
		teritories.GET("/:id/employees", h.GetEmployeesByTerritoryID)
	}
}

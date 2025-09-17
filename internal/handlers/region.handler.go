package handlers

import (
	"net/http"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type RegionHandler struct {
	Repo *repositories.RegionRepository
}

// @Summary Get all regions
// @Description Returns a list of all regions
// @Tags Regions
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Region
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/regions [get]
func (h *RegionHandler) GetAll(c *gin.Context) {
	regions, err := h.Repo.GetAllRegions(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, regions)
}

// @Summary Get region by ID
// @Description Returns a single region by ID
// @Tags Regions
// @Produce json
// @Param id path int true "Region ID"
// @Security BearerAuth
// @Success 200 {object} models.Region
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/regions/{id} [get]
func (h *RegionHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	region, err := h.Repo.GetRegionsByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, region)
}

// GET /territories/{id}/employees → karyawan di territory tertentu
// @Summary Get employees by territory ID
// @Description Returns a list of employees associated with a specific territory ID
// @Tags Regions
// @Produce json
// @Param id path string true "Territory ID"
// @Security BearerAuth
// @Success 200 {array} models.Employee
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/territories/{id}/employees [get]
func (h *RegionHandler) GetEmployeesByTerritoryID(c *gin.Context) {
	id := c.Param("id")
	employees, err := h.Repo.GetEmployeesByTerritoryID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

package handlers

import (
	"net/http"
	"northwind-api/internal/models"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type ShipperHandler struct {
	Repo *repositories.ShipperRepository
}

// @Summary Get all shippers
// @Description Returns a list of all shippers
// @Tags Shippers
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Shipper
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/shippers [get]
func (h *ShipperHandler) GetAll(c *gin.Context) {
	shippers, err := h.Repo.GetAllShippers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shippers)
}

// @Summary Get shipper by ID
// @Description Returns a single shipper by ID
// @Tags Shippers
// @Produce json
// @Param id path int true "Shipper ID"
// @Security BearerAuth
// @Success 200 {object} models.Shipper
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/shippers/{id} [get]
func (h *ShipperHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	shipper, err := h.Repo.GetShipperByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shipper)
}

// @Summary Create a new shipper
// @Description Creates a new shipper
// @Tags Shippers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param shipper body models.Shipper true "Shipper to create"
// @Success 201 {object} models.Shipper
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/shippers [post]
func (h *ShipperHandler) Create(c *gin.Context) {
	var shipper models.Shipper
	if err := c.ShouldBindJSON(&shipper); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Repo.CreateShipper(c.Request.Context(), shipper)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Shipper created successfully"})
}

// @Summary Update an existing shipper
// @Description Updates an existing shipper by ID
// @Tags Shippers
// @Accept json
// @Produce json
// @Param id path int true "Shipper ID"
// @Param shipper body models.Shipper true "Shipper data to update"
// @Security BearerAuth
// @Success 200 {object} models.Shipper
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/shippers/{id} [put]
func (h *ShipperHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var shipper models.Shipper
	if err := c.ShouldBindJSON(&shipper); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shipper.ShipperID = utils.ParseInt(id)
	if err := h.Repo.UpdateShipper(c.Request.Context(), &shipper); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shipper updated successfully"})
}

// @Summary Delete a shipper
// @Description Deletes a shipper by ID
// @Tags Shippers
// @Param id path int true "Shipper ID"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/shippers/{id} [delete]
func (h *ShipperHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Repo.DeleteShipper(c.Request.Context(), utils.ParseInt(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shipper deleted successfully"})
}

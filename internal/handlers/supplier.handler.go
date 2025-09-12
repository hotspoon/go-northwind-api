package handlers

import (
	"net/http"
	"northwind-api/internal/models"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	Repo *repositories.SupplierRepository
}

// @Summary Get all suppliers
// @Description Returns a list of all suppliers
// @Tags Suppliers
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Supplier
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/suppliers [get]
func (h *SupplierHandler) GetAll(c *gin.Context) {
	suppliers, err := h.Repo.GetAllSuppliers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}

// @Summary Get supplier by ID
// @Description Returns a single supplier by ID
// @Tags Suppliers
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.Supplier
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/suppliers/{id} [get]
func (h *SupplierHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.Repo.GetSupplierByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

// @Summary Create a new supplier
// @Description Creates a new supplier
// @Tags Suppliers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param supplier body models.Supplier true "Supplier to create"
// @Success 201 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/suppliers [post]
func (h *SupplierHandler) Create(c *gin.Context) {
	var supplier models.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.CreateSupplier(c.Request.Context(), &supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Supplier created successfully"})
}

// @Summary Update a supplier
// @Description Updates an existing supplier
// @Tags Suppliers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supplier ID"
// @Param supplier body models.Supplier true "Supplier to update"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/suppliers/{id} [put]
func (h *SupplierHandler) Update(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
	var supplier models.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	supplier.SupplierID = int64(id)
	if err := h.Repo.UpdateSupplier(c.Request.Context(), &supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Supplier updated successfully"})
}

// @Summary Delete a supplier
// @Description Deletes a supplier by ID
// @Tags Suppliers
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/suppliers/{id} [delete]
func (h *SupplierHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.DeleteSupplier(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully"})
}

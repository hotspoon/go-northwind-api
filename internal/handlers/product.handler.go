package handlers

import (
	"net/http"
	"northwind-api/internal/models"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Repo *repositories.ProductRepository
}

// @Summary Get all products
// @Description Returns a list of all products
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Product
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/products [get]
func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.Repo.GetAllProducts(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// @Summary Get product by ID
// @Description Returns a single product by ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/products/{id} [get]
func (h *ProductHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	product, err := h.Repo.GetProductByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary Create a new product
// @Description Creates a new product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body models.Product true "Product to create"
// @Success 201 {object} models.Product
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/products [post]
func (h *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Repo.CreateProduct(c.Request.Context(), product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// product.ID = int(id)

	c.JSON(http.StatusCreated, gin.H{"message": "product created successfully"})
}

// @Summary Update a product
// @Description Updates an existing product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product to update"
// @Success 200 {object} models.Product
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/products/{id} [put]
func (h *ProductHandler) Update(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	product.ProductID = id
	if err := h.Repo.UpdateProduct(c.Request.Context(), &product); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

// @Summary Delete a product
// @Description Deletes a product by ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/products/{id} [delete]
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.DeleteProduct(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

// GET /products/{id}/supplier
// @Summary Get supplier for a product
// @Description Returns the supplier for a given product ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} models.ProductSupplier
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/products/{id}/supplier [get]
func (h *ProductHandler) GetSupplier(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.Repo.GetSupplierByProductID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

// GET /products/{id}/category
// @Summary Get category for a product
// @Description Returns the category for a given product ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} models.ProductCategory
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/products/{id}/category [get]
func (h *ProductHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := h.Repo.GetCategoryByProductID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

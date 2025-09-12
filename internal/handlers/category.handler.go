package handlers

import (
	"net/http"
	"northwind-api/internal/models"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	Repo *repositories.CategoryRepository
}

// @Summary Get all categories
// @Description Returns a list of all categories
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Category
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/categories [get]
func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.Repo.GetAllCategories(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// @Summary Get category by ID
// @Description Returns a single category by ID
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/categories/{id} [get]
func (h *CategoryHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	category, err := h.Repo.GetCategoryByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// @Summary Create a new category
// @Description Creates a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category body models.Category true "Category to create"
// @Success 201 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/categories [post]
func (h *CategoryHandler) Create(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Repo.CreateCategory(c.Request.Context(), &category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	category.CategoryID = id
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

// @Summary Update a category
// @Description Updates an existing category
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category to update"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/categories/{id} [put]
func (h *CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	category.CategoryID = int64(utils.ParseInt(id))
	if err := h.Repo.UpdateCategory(c.Request.Context(), &category); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// @Summary Delete a category
// @Description Deletes a category by ID
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/categories/{id} [delete]
func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.DeleteCategory(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

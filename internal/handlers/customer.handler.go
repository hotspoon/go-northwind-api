package handlers

import (
	"net/http"
	"northwind-api/internal/models"
	"northwind-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Repo *repositories.CustomerRepository
}

// @Summary Get all customers
// @Description Returns a list of all customers
// @Tags Customers
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Customer
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/customers [get]
func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.Repo.GetAllCustomers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// @Summary Get customer by ID
// @Description Returns a single customer by ID
// @Tags Customers
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Customer
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/customers/{id} [get]
func (h *CustomerHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.Repo.GetCustomerByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// @Summary Create a new customer
// @Description Creates a new customer
// @Tags Customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param customer body models.Customer true "Customer to create"
// @Success 201 {object} models.Customer
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/customers [post]
func (h *CustomerHandler) Create(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	id, err := h.Repo.CreateCustomer(c.Request.Context(), &customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer.CustomerID = id

	c.JSON(http.StatusCreated, customer)
}

// @Summary Update an existing customer
// @Description Updates an existing customer by ID
// @Tags Customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Param customer body models.Customer true "Customer to update"
// @Success 200 {object} models.Customer
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/customers/{id} [put]
func (h *CustomerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	customer.CustomerID = id
	if err := h.Repo.UpdateCustomer(c.Request.Context(), &customer); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer updated successfully"})
}

// @Summary Delete a customer
// @Description Deletes a customer by ID
// @Tags Customers
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Customer
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/customers/{id} [delete]
func (h *CustomerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.DeleteCustomer(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
}

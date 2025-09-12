package handlers

import (
	"net/http"
	"northwind-api/internal/models"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Repo *repositories.OrderRepository
}

// @Summary Get all orders
// @Description Returns a list of all orders
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/orders [get]
func (h *OrderHandler) GetAll(c *gin.Context) {
	orders, err := h.Repo.GetAllOrders(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// @Summary Get paginated orders
// @Description Returns a paginated list of orders
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.Paginated[models.Order]
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/orders/paginated [get]
func (h *OrderHandler) GetPaginated(c *gin.Context) {
	// Parse query parameters with defaults
	page := 1
	pageSize := 10

	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if ps := c.Query("page_size"); ps != "" {
		if v, err := strconv.Atoi(ps); err == nil && v > 0 {
			pageSize = v
		}
	}

	// Fetch paginated orders
	paginatedOrders, err := h.Repo.GetOrdersPage(c.Request.Context(), page, pageSize)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, paginatedOrders)
}

// @Summary Get order by ID
// @Description Returns a single order by ID
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/orders/{id} [get]
func (h *OrderHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	order, err := h.Repo.GetOrderByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

// @Summary Create a new order
// @Description Creates a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body models.Order true "Order to create"
// @Success 201 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/orders [post]
func (h *OrderHandler) Create(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Repo.CreateOrder(c.Request.Context(), &order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order.OrderID = id

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order_id": order.OrderID})
}

// @Summary Update an order
// @Description Updates an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Param order body models.Order true "Order to update"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/orders/{id} [put]
func (h *OrderHandler) Update(c *gin.Context) {
	id := utils.ParseInt(c.Param("id"))
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	order.OrderID = int64(id)
	if err := h.Repo.UpdateOrder(c.Request.Context(), &order); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// @Summary Delete an order
// @Description Deletes an existing order
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/orders/{id} [delete]
func (h *OrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.DeleteOrder(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// GET /orders/{id}/details → detail item yang dipesan
// @Summary Get order details by Order ID
// @Description Returns order details for a specific order by Order ID
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {array} models.OrderDetail
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/orders/{id}/details [get]
func (h *OrderHandler) GetOrderDetails(c *gin.Context) {
	id := c.Param("id")
	details, err := h.Repo.GetOrderDetailsByOrderID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

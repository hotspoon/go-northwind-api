package handlers

import (
	"net/http"
	"northwind-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	Repo *repositories.ReportRepository
}

// @Summary Top customers by total purchases
// @Description Returns customers with the highest total purchases
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.TopCustomer
// @Router /api/v1/reports/top-customers [get]
func (h *ReportHandler) GetTopCustomers(c *gin.Context) {
	result, err := h.Repo.GetTopCustomers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Top selling products
// @Description Returns products with the highest sales
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.TopProduct
// @Router /api/v1/reports/top-products [get]
func (h *ReportHandler) GetTopProducts(c *gin.Context) {
	result, err := h.Repo.GetTopProducts(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Sales by category
// @Description Returns sales report grouped by category
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.SalesByCategory
// @Router /api/v1/reports/sales-by-category [get]
func (h *ReportHandler) GetSalesByCategory(c *gin.Context) {
	result, err := h.Repo.GetSalesByCategory(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Sales by employee
// @Description Returns sales report grouped by employee
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.SalesByEmployee
// @Router /api/v1/reports/sales-by-employee [get]
func (h *ReportHandler) GetSalesByEmployee(c *gin.Context) {
	result, err := h.Repo.GetSalesByEmployee(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Sales summary
// @Description Returns an overall sales summary
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.SalesSummary
// @Router /api/v1/reports/sales-summary [get]
func (h *ReportHandler) GetSalesSummary(c *gin.Context) {
	result, err := h.Repo.GetSalesSummary(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Monthly sales
// @Description Returns sales totals grouped by month
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MonthlySales
// @Router /api/v1/reports/monthly-sales [get]
func (h *ReportHandler) GetMonthlySales(c *gin.Context) {
	result, err := h.Repo.GetMonthlySales(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Inventory status
// @Description Returns current inventory levels and status
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.InventoryStatus
// @Router /api/v1/reports/inventory-status [get]
func (h *ReportHandler) GetInventoryStatus(c *gin.Context) {
	result, err := h.Repo.GetInventoryStatus(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Top suppliers
// @Description Returns suppliers ranked by performance or volume
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.TopSupplier
// @Router /api/v1/reports/top-suppliers [get]
func (h *ReportHandler) GetTopSuppliers(c *gin.Context) {
	result, err := h.Repo.GetTopSuppliers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Customer growth
// @Description Returns customer growth over time
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.CustomerGrowth
// @Router /api/v1/reports/customer-growth [get]
func (h *ReportHandler) GetCustomerGrowth(c *gin.Context) {
	result, err := h.Repo.GetCustomerGrowth(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Order status summary
// @Description Returns a breakdown of orders by status
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.OrderStatusSummary
// @Router /api/v1/reports/order-status-summary [get]
func (h *ReportHandler) GetOrderStatusSummary(c *gin.Context) {
	result, err := h.Repo.GetOrderStatusSummary(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Region sales
// @Description Returns sales grouped by region
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.RegionSales
// @Router /api/v1/reports/region-sales [get]
func (h *ReportHandler) GetRegionSales(c *gin.Context) {
	result, err := h.Repo.GetRegionSales(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Employee performance
// @Description Returns performance metrics for employees
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.EmployeePerformance
// @Router /api/v1/reports/employee-performance [get]
func (h *ReportHandler) GetEmployeePerformance(c *gin.Context) {
	result, err := h.Repo.GetEmployeePerformance(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Product profitability
// @Description Returns profitability metrics for products
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.ProductProfitability
// @Router /api/v1/reports/product-profitability [get]
func (h *ReportHandler) GetProductProfitability(c *gin.Context) {
	result, err := h.Repo.GetProductProfitability(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Average order value
// @Description Returns the average value of orders
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.AverageOrderValue
// @Router /api/v1/reports/average-order-value [get]
func (h *ReportHandler) GetAverageOrderValue(c *gin.Context) {
	result, err := h.Repo.GetAverageOrderValue(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

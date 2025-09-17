package routes

import (
	"northwind-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup, h *handlers.ReportHandler) {
	reports := rg.Group("/reports")
	{
		reports.GET("/top-customers", h.GetTopCustomers)
		reports.GET("/top-products", h.GetTopProducts)
		reports.GET("/sales-by-category", h.GetSalesByCategory)
		reports.GET("/sales-by-employee", h.GetSalesByEmployee)
		reports.GET("/sales-summary", h.GetSalesSummary)
		reports.GET("/monthly-sales", h.GetMonthlySales)
		reports.GET("/inventory-status", h.GetInventoryStatus)
		reports.GET("/top-suppliers", h.GetTopSuppliers)
		reports.GET("/customer-growth", h.GetCustomerGrowth)
		reports.GET("/order-status-summary", h.GetOrderStatusSummary)
		reports.GET("/region-sales", h.GetRegionSales)
		reports.GET("/employee-performance", h.GetEmployeePerformance)
		reports.GET("/product-profitability", h.GetProductProfitability)
		reports.GET("/average-order-value", h.GetAverageOrderValue)
	}
}

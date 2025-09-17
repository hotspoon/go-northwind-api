package models

type TopCustomer struct {
	CustomerID    string  `json:"customer_id"`
	CompanyName   string  `json:"company_name"`
	TotalPurchase float64 `json:"total_purchase"`
}

type TopProduct struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	TotalSold   int    `json:"total_sold"`
}

type SalesByCategory struct {
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
	TotalSales   float64 `json:"total_sales"`
}

type SalesByEmployee struct {
	EmployeeID   int     `json:"employee_id"`
	EmployeeName string  `json:"employee_name"`
	TotalSales   float64 `json:"total_sales"`
}

type SalesSummary struct {
	TotalRevenue      float64 `json:"total_revenue"`
	TotalOrders       int64   `json:"total_orders"`
	TotalCustomers    int64   `json:"total_customers"`
	AverageOrderValue float64 `json:"average_order_value"`
	FirstOrderDate    string  `json:"first_order_date"`
	LastOrderDate     string  `json:"last_order_date"`
}

type MonthlySales struct {
	YearMonth  string  `json:"year_month"` // e.g. 1997-01
	TotalSales float64 `json:"total_sales"`
	Orders     int64   `json:"orders"`
}

type InventoryStatus struct {
	ProductID    int64  `json:"product_id"`
	ProductName  string `json:"product_name"`
	UnitsInStock int64  `json:"units_in_stock"`
	UnitsOnOrder int64  `json:"units_on_order"`
	ReorderLevel int64  `json:"reorder_level"`
	Status       string `json:"status"` // "OK", "LOW", "OUT"
}

type TopSupplier struct {
	SupplierID  int64   `json:"supplier_id"`
	CompanyName string  `json:"company_name"`
	TotalSales  float64 `json:"total_sales"`
	TotalQty    int64   `json:"total_qty"`
}

type CustomerGrowth struct {
	YearMonth        string `json:"year_month"`
	NewCustomers     int64  `json:"new_customers"`     // first time placing an order that month
	CumulativeUnique int64  `json:"cumulative_unique"` // running total
}

type OrderStatusSummary struct {
	Status string `json:"status"` // "Pending", "Shipped", "Late"
	Count  int64  `json:"count"`
}

type RegionSales struct {
	Region     string  `json:"region"` // using ShipRegion; fallback to ShipCountry if NULL
	TotalSales float64 `json:"total_sales"`
	Orders     int64   `json:"orders"`
}

type EmployeePerformance struct {
	EmployeeID      int64   `json:"employee_id"`
	EmployeeName    string  `json:"employee_name"`
	TotalSales      float64 `json:"total_sales"`
	OrdersHandled   int64   `json:"orders_handled"`
	AvgOrderValue   float64 `json:"avg_order_value"`
	UniqueCustomers int64   `json:"unique_customers"`
}

type ProductProfitability struct {
	ProductID      int64   `json:"product_id"`
	ProductName    string  `json:"product_name"`
	Revenue        float64 `json:"revenue"`
	COGS           float64 `json:"cogs"` // approximated using Products.UnitPrice as cost (see note)
	GrossProfit    float64 `json:"gross_profit"`
	GrossMarginPct float64 `json:"gross_margin_pct"`
}

type AverageOrderValue struct {
	Average float64 `json:"average"`
}

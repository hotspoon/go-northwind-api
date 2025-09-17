package repositories

import (
	"context"
	"database/sql"
	"northwind-api/internal/models"
)

type ReportRepository struct {
	DB *sql.DB
}

func (r *ReportRepository) GetTopCustomers(ctx context.Context) ([]models.TopCustomer, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT c.CustomerID, c.CompanyName, SUM(od.UnitPrice * od.Quantity) AS TotalPurchase
        FROM Customers c
        JOIN Orders o ON c.CustomerID = o.CustomerID
        JOIN OrderDetails od ON o.OrderID = od.OrderID
        GROUP BY c.CustomerID, c.CompanyName
        ORDER BY TotalPurchase DESC
        LIMIT 10
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.TopCustomer
	for rows.Next() {
		var tc models.TopCustomer
		if err := rows.Scan(&tc.CustomerID, &tc.CompanyName, &tc.TotalPurchase); err != nil {
			return nil, err
		}
		result = append(result, tc)
	}
	return result, nil
}

func (r *ReportRepository) GetTopProducts(ctx context.Context) ([]models.TopProduct, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT p.ProductID, p.ProductName, SUM(od.Quantity) AS TotalSold
        FROM Products p
        JOIN OrderDetails od ON p.ProductID = od.ProductID
        GROUP BY p.ProductID, p.ProductName
        ORDER BY TotalSold DESC
        LIMIT 10
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.TopProduct
	for rows.Next() {
		var tp models.TopProduct
		if err := rows.Scan(&tp.ProductID, &tp.ProductName, &tp.TotalSold); err != nil {
			return nil, err
		}
		result = append(result, tp)
	}
	return result, nil
}

func (r *ReportRepository) GetSalesByCategory(ctx context.Context) ([]models.SalesByCategory, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT c.CategoryID, c.CategoryName, SUM(od.UnitPrice * od.Quantity) AS TotalSales
        FROM Categories c
        JOIN Products p ON c.CategoryID = p.CategoryID
        JOIN OrderDetails od ON p.ProductID = od.ProductID
        GROUP BY c.CategoryID, c.CategoryName
        ORDER BY TotalSales DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.SalesByCategory
	for rows.Next() {
		var sc models.SalesByCategory
		if err := rows.Scan(&sc.CategoryID, &sc.CategoryName, &sc.TotalSales); err != nil {
			return nil, err
		}
		result = append(result, sc)
	}
	return result, nil
}

func (r *ReportRepository) GetSalesByEmployee(ctx context.Context) ([]models.SalesByEmployee, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT e.EmployeeID, e.FirstName || ' ' || e.LastName AS EmployeeName, SUM(od.UnitPrice * od.Quantity) AS TotalSales
        FROM Employees e
        JOIN Orders o ON e.EmployeeID = o.EmployeeID
        JOIN OrderDetails od ON o.OrderID = od.OrderID
        GROUP BY e.EmployeeID, EmployeeName
        ORDER BY TotalSales DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.SalesByEmployee
	for rows.Next() {
		var se models.SalesByEmployee
		if err := rows.Scan(&se.EmployeeID, &se.EmployeeName, &se.TotalSales); err != nil {
			return nil, err
		}
		result = append(result, se)
	}
	return result, nil
}

// Sales summary (total revenue, orders, customers, AOV, first/last order date)
func (r *ReportRepository) GetSalesSummary(ctx context.Context) (models.SalesSummary, error) {
	var s models.SalesSummary
	// Revenue menghitung diskon: UnitPrice * Quantity * (1 - Discount)
	q := `
		WITH ord AS (
			SELECT o.OrderID, o.OrderDate,
				   SUM(od.UnitPrice * od.Quantity * (1.0 - od.Discount)) AS order_total
			FROM Orders o
			JOIN OrderDetails od ON o.OrderID = od.OrderID
			GROUP BY o.OrderID, o.OrderDate
		),
		first_last AS (
			SELECT MIN(OrderDate) AS first_dt, MAX(OrderDate) AS last_dt FROM Orders
		),
		total_customers AS (
			SELECT COUNT(DISTINCT CustomerID) AS cnt FROM Orders
		)
		SELECT
			COALESCE(SUM(order_total),0) AS total_revenue,
			COALESCE(COUNT(*),0)        AS total_orders,
			(SELECT cnt FROM total_customers) AS total_customers,
			CASE WHEN COUNT(*)=0 THEN 0 ELSE SUM(order_total)/COUNT(*) END AS avg_order_value,
			(SELECT first_dt FROM first_last),
			(SELECT last_dt FROM first_last)
		FROM ord;
	`
	row := r.DB.QueryRowContext(ctx, q)
	if err := row.Scan(
		&s.TotalRevenue,
		&s.TotalOrders,
		&s.TotalCustomers,
		&s.AverageOrderValue,
		&s.FirstOrderDate,
		&s.LastOrderDate,
	); err != nil {
		return s, err
	}
	return s, nil
}

// Monthly sales (group by YYYY-MM)
func (r *ReportRepository) GetMonthlySales(ctx context.Context) ([]models.MonthlySales, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT strftime('%Y-%m', o.OrderDate) AS ym,
		       SUM(od.UnitPrice * od.Quantity * (1.0 - od.Discount)) AS total_sales,
		       COUNT(DISTINCT o.OrderID) AS orders
		FROM Orders o
		JOIN OrderDetails od ON o.OrderID = od.OrderID
		GROUP BY ym
		ORDER BY ym;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.MonthlySales
	for rows.Next() {
		var m models.MonthlySales
		if err := rows.Scan(&m.YearMonth, &m.TotalSales, &m.Orders); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// Inventory status (simple rule: OUT=0, LOW<=ReorderLevel, else OK)
func (r *ReportRepository) GetInventoryStatus(ctx context.Context) ([]models.InventoryStatus, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT p.ProductID, p.ProductName,
		       COALESCE(p.UnitsInStock,0),
		       COALESCE(p.UnitsOnOrder,0),
		       COALESCE(p.ReorderLevel,0),
		       CASE
		         WHEN COALESCE(p.UnitsInStock,0) = 0 THEN 'OUT'
		         WHEN COALESCE(p.UnitsInStock,0) <= COALESCE(p.ReorderLevel,0) THEN 'LOW'
		         ELSE 'OK'
		       END AS status
		FROM Products p
		ORDER BY p.ProductName;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.InventoryStatus
	for rows.Next() {
		var i models.InventoryStatus
		if err := rows.Scan(&i.ProductID, &i.ProductName, &i.UnitsInStock, &i.UnitsOnOrder, &i.ReorderLevel, &i.Status); err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, rows.Err()
}

// Top suppliers by sales (sum revenue & qty via their products)
func (r *ReportRepository) GetTopSuppliers(ctx context.Context) ([]models.TopSupplier, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT s.SupplierID, s.CompanyName,
		       SUM(od.UnitPrice * od.Quantity * (1.0 - od.Discount)) AS total_sales,
		       SUM(od.Quantity) AS total_qty
		FROM Suppliers s
		JOIN Products p ON s.SupplierID = p.SupplierID
		JOIN OrderDetails od ON p.ProductID = od.ProductID
		GROUP BY s.SupplierID, s.CompanyName
		ORDER BY total_sales DESC
		LIMIT 10;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.TopSupplier
	for rows.Next() {
		var t models.TopSupplier
		if err := rows.Scan(&t.SupplierID, &t.CompanyName, &t.TotalSales, &t.TotalQty); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// Customer growth (first order month per customer; with running total)
func (r *ReportRepository) GetCustomerGrowth(ctx context.Context) ([]models.CustomerGrowth, error) {
	// SQLite prior to 3.25.0 has limited window functions; assuming modern SQLite with window support.
	rows, err := r.DB.QueryContext(ctx, `
		WITH first_orders AS (
			SELECT o.CustomerID, MIN(date(o.OrderDate)) AS first_date
			FROM Orders o
			GROUP BY o.CustomerID
		),
		first_month AS (
			SELECT strftime('%Y-%m', first_date) AS ym, COUNT(*) AS new_customers
			FROM first_orders
			GROUP BY ym
		),
		series AS (
			SELECT ym
			FROM (
				SELECT DISTINCT strftime('%Y-%m', OrderDate) AS ym FROM Orders
			)
		),
		filled AS (
			SELECT s.ym,
			       COALESCE(fm.new_customers, 0) AS new_customers
			FROM series s
			LEFT JOIN first_month fm ON fm.ym = s.ym
		)
		SELECT ym,
		       new_customers,
		       SUM(new_customers) OVER (ORDER BY ym ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS cumulative_unique
		FROM filled
		ORDER BY ym;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.CustomerGrowth
	for rows.Next() {
		var cg models.CustomerGrowth
		if err := rows.Scan(&cg.YearMonth, &cg.NewCustomers, &cg.CumulativeUnique); err != nil {
			return nil, err
		}
		out = append(out, cg)
	}
	return out, rows.Err()
}

// Order status summary (derive simple statuses)
func (r *ReportRepository) GetOrderStatusSummary(ctx context.Context) ([]models.OrderStatusSummary, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT status, COUNT(*) AS cnt
		FROM (
			SELECT CASE
				WHEN o.ShippedDate IS NULL THEN 'Pending'
				WHEN o.RequiredDate IS NOT NULL AND date(o.ShippedDate) > date(o.RequiredDate) THEN 'Late'
				ELSE 'Shipped'
			END AS status
			FROM Orders o
		) s
		GROUP BY status
		ORDER BY cnt DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.OrderStatusSummary
	for rows.Next() {
		var s models.OrderStatusSummary
		if err := rows.Scan(&s.Status, &s.Count); err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, rows.Err()
}

// Region sales (pakai ShipRegion; fallback ke ShipCountry jika ShipRegion NULL)
func (r *ReportRepository) GetRegionSales(ctx context.Context) ([]models.RegionSales, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT COALESCE(NULLIF(TRIM(o.ShipRegion),''), o.ShipCountry) AS region,
		       SUM(od.UnitPrice * od.Quantity * (1.0 - od.Discount)) AS total_sales,
		       COUNT(DISTINCT o.OrderID) AS orders
		FROM Orders o
		JOIN OrderDetails od ON o.OrderID = od.OrderID
		GROUP BY region
		ORDER BY total_sales DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.RegionSales
	for rows.Next() {
		var rs models.RegionSales
		if err := rows.Scan(&rs.Region, &rs.TotalSales, &rs.Orders); err != nil {
			return nil, err
		}
		out = append(out, rs)
	}
	return out, rows.Err()
}

// Employee performance (total sales, orders handled, AOV, unique customers)
func (r *ReportRepository) GetEmployeePerformance(ctx context.Context) ([]models.EmployeePerformance, error) {
	rows, err := r.DB.QueryContext(ctx, `
		WITH emp_orders AS (
			SELECT o.EmployeeID, o.OrderID, o.CustomerID,
			       SUM(od.UnitPrice * od.Quantity * (1.0 - od.Discount)) AS order_total
			FROM Orders o
			JOIN OrderDetails od ON o.OrderID = od.OrderID
			GROUP BY o.EmployeeID, o.OrderID, o.CustomerID
		)
		SELECT e.EmployeeID,
		       (e.FirstName || ' ' || e.LastName) AS EmployeeName,
		       COALESCE(SUM(eo.order_total),0) AS total_sales,
		       COUNT(DISTINCT eo.OrderID) AS orders_handled,
		       CASE WHEN COUNT(DISTINCT eo.OrderID)=0 THEN 0
		            ELSE SUM(eo.order_total) / COUNT(DISTINCT eo.OrderID) END AS avg_order_value,
		       COUNT(DISTINCT eo.CustomerID) AS unique_customers
		FROM Employees e
		LEFT JOIN emp_orders eo ON e.EmployeeID = eo.EmployeeID
		GROUP BY e.EmployeeID, EmployeeName
		ORDER BY total_sales DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.EmployeePerformance
	for rows.Next() {
		var ep models.EmployeePerformance
		if err := rows.Scan(
			&ep.EmployeeID, &ep.EmployeeName, &ep.TotalSales, &ep.OrdersHandled,
			&ep.AvgOrderValue, &ep.UniqueCustomers,
		); err != nil {
			return nil, err
		}
		out = append(out, ep)
	}
	return out, rows.Err()
}

// Product profitability (NOTE: pakai Products.UnitPrice sbg pendekatan COGS → kasar)
func (r *ReportRepository) GetProductProfitability(ctx context.Context) ([]models.ProductProfitability, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT p.ProductID, p.ProductName,
		       -- revenue pakai harga jual actual di OrderDetails dengan diskon
		       SUM(od.UnitPrice * od.Quantity * (1.0 - od.Discount)) AS revenue,
		       -- COGS pakai pendekatan: Products.UnitPrice sbg cost (kasar; Northwind tdk punya cost)
		       SUM(p.UnitPrice * od.Quantity) AS cogs
		FROM Products p
		JOIN OrderDetails od ON p.ProductID = od.ProductID
		GROUP BY p.ProductID, p.ProductName
		ORDER BY revenue DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.ProductProfitability
	for rows.Next() {
		var pp models.ProductProfitability
		if err := rows.Scan(&pp.ProductID, &pp.ProductName, &pp.Revenue, &pp.COGS); err != nil {
			return nil, err
		}
		pp.GrossProfit = pp.Revenue - pp.COGS
		if pp.Revenue == 0 {
			pp.GrossMarginPct = 0
		} else {
			pp.GrossMarginPct = (pp.GrossProfit / pp.Revenue) * 100.0
		}
		out = append(out, pp)
	}
	return out, rows.Err()
}

// Average order value (overall)
func (r *ReportRepository) GetAverageOrderValue(ctx context.Context) (models.AverageOrderValue, error) {
	var aov models.AverageOrderValue
	row := r.DB.QueryRowContext(ctx, `
		WITH ord AS (
			SELECT o.OrderID,
			       SUM(od.UnitPrice * od.Quantity * (1.0 - od.Discount)) AS order_total
			FROM Orders o
			JOIN OrderDetails od ON o.OrderID = od.OrderID
			GROUP BY o.OrderID
		)
		SELECT CASE WHEN COUNT(*)=0 THEN 0 ELSE SUM(order_total)/COUNT(*) END
		FROM ord;
	`)
	if err := row.Scan(&aov.Average); err != nil {
		return aov, err
	}
	return aov, nil
}

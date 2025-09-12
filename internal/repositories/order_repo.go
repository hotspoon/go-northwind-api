package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"northwind-api/internal/models"

	"github.com/rs/zerolog/log"
)

type OrderRepository struct {
	DB *sql.DB
}

func (r *OrderRepository) CreateOrder(ctx context.Context, o *models.Order) (int64, error) {
	result, err := r.DB.ExecContext(ctx, `
		INSERT INTO Orders (CustomerID, EmployeeID, OrderDate, RequiredDate, ShippedDate,
			ShipVia, Freight, ShipName, ShipAddress, ShipCity, ShipRegion, ShipPostalCode, ShipCountry)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, o.CustomerID, o.EmployeeID, o.OrderDate, o.RequiredDate, o.ShippedDate,
		o.ShipVia, o.Freight, o.ShipName, o.ShipAddress, o.ShipCity,
		o.ShipRegion, o.ShipPostalCode, o.ShipCountry)
	if err != nil {
		log.Error().Err(err).Msg("error creating order")
		return 0, fmt.Errorf("error creating order: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("error getting last insert id for order")
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}
	return id, nil
}

func (r *OrderRepository) GetOrdersPage(ctx context.Context, page, pageSize int) (*models.Paginated[models.Order], error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// 1) Ambil total rows
	var total int
	if err := r.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM Orders`).Scan(&total); err != nil {
		log.Error().Err(err).Msg("error counting orders")
		return nil, fmt.Errorf("error counting orders: %w", err)
	}

	// 2) Ambil page data (urutkan agar deterministik)
	rows, err := r.DB.QueryContext(ctx, `
		SELECT OrderID, CustomerID, EmployeeID, OrderDate, RequiredDate, ShippedDate,
		       ShipVia, Freight, ShipName, ShipAddress, ShipCity, ShipRegion, ShipPostalCode, ShipCountry
		FROM Orders
		ORDER BY OrderID ASC
		LIMIT ? OFFSET ?`,
		pageSize, offset,
	)
	if err != nil {
		log.Error().Err(err).Msg("error fetching paged orders")
		return nil, fmt.Errorf("error fetching paged orders: %w", err)
	}
	defer rows.Close()

	items := make([]models.Order, 0, pageSize)
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(
			&o.OrderID, &o.CustomerID, &o.EmployeeID,
			&o.OrderDate, &o.RequiredDate, &o.ShippedDate,
			&o.ShipVia, &o.Freight, &o.ShipName, &o.ShipAddress,
			&o.ShipCity, &o.ShipRegion, &o.ShipPostalCode, &o.ShipCountry,
		); err != nil {
			log.Error().Err(err).Msg("error scanning order row")
			return nil, fmt.Errorf("error scanning order row: %w", err)
		}
		items = append(items, o)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating order rows")
		return nil, fmt.Errorf("error iterating order rows: %w", err)
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	res := &models.Paginated[models.Order]{
		Items:      items,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: total,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
	return res, nil
}

func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT OrderID, CustomerID, EmployeeID, OrderDate, RequiredDate, ShippedDate,
			ShipVia, Freight, ShipName, ShipAddress, ShipCity, ShipRegion, ShipPostalCode, ShipCountry
		FROM Orders
	`)
	if err != nil {
		log.Error().Err(err).Msg("error fetching orders")
		return nil, fmt.Errorf("error fetching orders: %w", err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID, &order.CustomerID, &order.EmployeeID,
			&order.OrderDate, &order.RequiredDate, &order.ShippedDate,
			&order.ShipVia, &order.Freight, &order.ShipName,
			&order.ShipAddress, &order.ShipCity, &order.ShipRegion,
			&order.ShipPostalCode, &order.ShipCountry); err != nil {
			log.Error().Err(err).Msg("error scanning order row")
			return nil, fmt.Errorf("error scanning order row: %w", err)
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over order rows")
		return nil, fmt.Errorf("error iterating over order rows: %w", err)
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, id int) (models.Order, error) {
	var order models.Order
	err := r.DB.QueryRowContext(ctx, `
		SELECT OrderID, CustomerID, EmployeeID, OrderDate, RequiredDate, ShippedDate,
			ShipVia, Freight, ShipName, ShipAddress, ShipCity, ShipRegion, ShipPostalCode, ShipCountry
		FROM Orders
		WHERE OrderID = ?
	`, id).Scan(&order.OrderID, &order.CustomerID, &order.EmployeeID,
		&order.OrderDate, &order.RequiredDate, &order.ShippedDate,
		&order.ShipVia, &order.Freight, &order.ShipName,
		&order.ShipAddress, &order.ShipCity, &order.ShipRegion,
		&order.ShipPostalCode, &order.ShipCountry)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Int("id", id).Msg("Order not found")
			return order, fmt.Errorf("order not found")
		}
		log.Error().Err(err).Int("id", id).Msg("error fetching order by ID")
		return order, fmt.Errorf("error fetching order by ID: %w", err)
	}
	return order, nil
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, o *models.Order) error {
	_, err := r.DB.ExecContext(ctx, `
		UPDATE Orders
		SET CustomerID = ?, EmployeeID = ?, OrderDate = ?, RequiredDate = ?, ShippedDate = ?,
			ShipVia = ?, Freight = ?, ShipName = ?, ShipAddress = ?, ShipCity = ?,
			ShipRegion = ?, ShipPostalCode = ?, ShipCountry = ?
		WHERE OrderID = ?
	`, o.CustomerID, o.EmployeeID, o.OrderDate, o.RequiredDate, o.ShippedDate,
		o.ShipVia, o.Freight, o.ShipName, o.ShipAddress, o.ShipCity,
		o.ShipRegion, o.ShipPostalCode, o.ShipCountry, o.OrderID)
	if err != nil {
		log.Error().Err(err).Int64("id", o.OrderID).Msg("error updating order")
		return fmt.Errorf("error updating order: %w", err)
	}
	return nil
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, `
		DELETE FROM Orders
		WHERE OrderID = ?
	`, id)
	if err != nil {
		log.Error().Err(err).Int("id", id).Msg("error deleting order")
		return fmt.Errorf("error deleting order: %w", err)
	}
	return nil
}

// GET /orders/{id}/details → detail item yang dipesan
func (r *OrderRepository) GetOrderDetailsByOrderID(ctx context.Context, orderID int) ([]models.OrderDetail, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT OrderID, ProductID, UnitPrice, Quantity, Discount
		FROM OrderDetails
		WHERE OrderID = ?
	`, orderID)
	if err != nil {
		log.Error().Err(err).Int("orderID", orderID).Msg("error fetching order details")
		return nil, fmt.Errorf("error fetching order details: %w", err)
	}
	defer rows.Close()

	var details []models.OrderDetail
	for rows.Next() {
		var detail models.OrderDetail
		if err := rows.Scan(&detail.OrderID, &detail.ProductID,
			&detail.UnitPrice, &detail.Quantity, &detail.Discount); err != nil {
			log.Error().Err(err).Int("orderID", orderID).Msg("error scanning order detail row")
			return nil, fmt.Errorf("error scanning order detail row: %w", err)
		}
		details = append(details, detail)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Int("orderID", orderID).Msg("error iterating over order detail rows")
		return nil, fmt.Errorf("error iterating over order detail rows: %w", err)
	}
	return details, nil
}

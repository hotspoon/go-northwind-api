package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"northwind-api/internal/models"

	"github.com/rs/zerolog/log"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p models.Product) (int64, error) {
	result, err := r.DB.ExecContext(ctx, `
		INSERT INTO Products
			(ProductName, SupplierID, CategoryID, QuantityPerUnit, UnitPrice,
			 UnitsInStock, UnitsOnOrder, ReorderLevel, Discontinued)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		p.ProductName,
		p.SupplierID,      // *int -> akan menjadi NULL jika nil
		p.CategoryID,      // *int
		p.QuantityPerUnit, // *string
		p.UnitPrice,
		p.UnitsInStock,
		p.UnitsOnOrder,
		p.ReorderLevel,
		p.Discontinued, // di Northwind klasik biasanya bool/int; di modelmu string
	)
	if err != nil {
		log.Error().Err(err).Msg("error creating product")
		return 0, fmt.Errorf("error creating product: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("error getting last insert id for product")
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}
	return id, nil
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			ProductID,
			ProductName,
			SupplierID,
			CategoryID,
			QuantityPerUnit,
			UnitPrice,
			UnitsInStock,
			UnitsOnOrder,
			ReorderLevel,
			Discontinued
		FROM Products
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query products")
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	defer rows.Close()

	var list []models.Product
	for rows.Next() {
		var (
			p                 models.Product
			supplierID        sql.NullInt64
			categoryID        sql.NullInt64
			quantityPerUnitNS sql.NullString
		)
		if err := rows.Scan(
			&p.ProductID,
			&p.ProductName,
			&supplierID,
			&categoryID,
			&quantityPerUnitNS,
			&p.UnitPrice,
			&p.UnitsInStock,
			&p.UnitsOnOrder,
			&p.ReorderLevel,
			&p.Discontinued,
		); err != nil {
			log.Error().Err(err).Msg("failed to scan product")
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		p.SupplierID = ptrInt64OrNil(supplierID)
		p.CategoryID = ptrInt64OrNil(categoryID)
		p.QuantityPerUnit = ptrStringOrNil(quantityPerUnitNS)

		list = append(list, p)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("rows error on products")
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return list, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (models.Product, error) {
	var (
		p                 models.Product
		supplierID        sql.NullInt64
		categoryID        sql.NullInt64
		quantityPerUnitNS sql.NullString
	)
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			ProductID,
			ProductName,
			SupplierID,
			CategoryID,
			QuantityPerUnit,
			UnitPrice,
			UnitsInStock,
			UnitsOnOrder,
			ReorderLevel,
			Discontinued
		FROM Products
		WHERE ProductID = ?
	`, id).Scan(
		&p.ProductID,
		&p.ProductName,
		&supplierID,
		&categoryID,
		&quantityPerUnitNS,
		&p.UnitPrice,
		&p.UnitsInStock,
		&p.UnitsOnOrder,
		&p.ReorderLevel,
		&p.Discontinued,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, fmt.Errorf("product with ID %d not found", id)
		}
		log.Error().Err(err).Int("product_id", id).Msg("failed to query product by id")
		return models.Product{}, fmt.Errorf("error fetching product by id: %w", err)
	}

	p.SupplierID = ptrInt64OrNil(supplierID)
	p.CategoryID = ptrInt64OrNil(categoryID)
	p.QuantityPerUnit = ptrStringOrNil(quantityPerUnitNS)
	return p, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, p *models.Product) error {
	_, err := r.DB.ExecContext(ctx, `
		UPDATE Products SET
			ProductName = ?,
			SupplierID = ?,
			CategoryID = ?,
			QuantityPerUnit = ?,
			UnitPrice = ?,
			UnitsInStock = ?,
			UnitsOnOrder = ?,
			ReorderLevel = ?,
			Discontinued = ?
		WHERE ProductID = ?
	`,
		p.ProductName,
		p.SupplierID,      // *int, boleh nil
		p.CategoryID,      // *int
		p.QuantityPerUnit, // *string
		p.UnitPrice,
		p.UnitsInStock,
		p.UnitsOnOrder,
		p.ReorderLevel,
		p.Discontinued,
		p.ProductID,
	)
	if err != nil {
		log.Error().Err(err).Int("product_id", p.ProductID).Msg("error updating product")
		return fmt.Errorf("error updating product: %w", err)
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	result, err := r.DB.ExecContext(ctx, `DELETE FROM Products WHERE ProductID = ?`, id)
	if err != nil {
		log.Error().Err(err).Int("product_id", id).Msg("error deleting product")
		return fmt.Errorf("error deleting product: %w", err)
	}
	aff, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("product_id", id).Msg("error fetching rows affected for product delete")
		return fmt.Errorf("error fetching rows affected: %w", err)
	}
	if aff == 0 {
		log.Warn().Int("product_id", id).Msg("no product found to delete")
		return fmt.Errorf("no product found with ID %d", id)
	}
	log.Info().Int("product_id", id).Msg("product deleted")
	return nil
}

// --- Lookup relasi: supplier & category ---

func (r *ProductRepository) GetSupplierByProductID(ctx context.Context, productID int) (models.ProductSupplier, error) {
	var out models.ProductSupplier
	err := r.DB.QueryRowContext(ctx, `
		SELECT s.SupplierID, s.CompanyName
		FROM Products p
		JOIN Suppliers s ON s.SupplierID = p.SupplierID
		WHERE p.ProductID = ?
	`, productID).Scan(&out.SupplierID, &out.CompanyName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ProductSupplier{}, fmt.Errorf("supplier for product %d not found", productID)
		}
		log.Error().Err(err).Int("product_id", productID).Msg("failed to query supplier by product id")
		return models.ProductSupplier{}, fmt.Errorf("error fetching supplier for product: %w", err)
	}
	return out, nil
}

func (r *ProductRepository) GetCategoryByProductID(ctx context.Context, productID int) (models.ProductCategory, error) {
	var out models.ProductCategory
	err := r.DB.QueryRowContext(ctx, `
		SELECT c.CategoryID, c.CategoryName
		FROM Products p
		JOIN Categories c ON c.CategoryID = p.CategoryID
		WHERE p.ProductID = ?
	`, productID).Scan(&out.CategoryID, &out.CategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ProductCategory{}, fmt.Errorf("category for product %d not found", productID)
		}
		log.Error().Err(err).Int("product_id", productID).Msg("failed to query category by product id")
		return models.ProductCategory{}, fmt.Errorf("error fetching category for product: %w", err)
	}
	return out, nil
}

// --- Helpers untuk konversi sql.Null* ke pointer ---

func ptrInt64OrNil(n sql.NullInt64) *int {
	if !n.Valid {
		return nil
	}
	v := int(n.Int64)
	return &v
}

func ptrStringOrNil(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	v := s.String
	return &v
}

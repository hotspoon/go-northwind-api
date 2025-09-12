package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"northwind-api/internal/models"

	"github.com/rs/zerolog/log"
)

type SupplierRepository struct {
	DB *sql.DB
}

func (r *SupplierRepository) GetAllSuppliers(ctx context.Context) ([]models.Supplier, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			SupplierID,
			CompanyName,
			ContactName,
			ContactTitle,
			Address,
			City,
			Region,
			PostalCode,
			Country,
			Phone,
			Fax,
			HomePage
		FROM Suppliers
	`)
	if err != nil {
		log.Error().Err(err).Msg("error fetching suppliers")
		return nil, fmt.Errorf("error fetching suppliers: %w", err)
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		if err := rows.Scan(
			&supplier.SupplierID,
			&supplier.CompanyName,
			&supplier.ContactName,
			&supplier.ContactTitle,
			&supplier.Address,
			&supplier.City,
			&supplier.Region,
			&supplier.PostalCode,
			&supplier.Country,
			&supplier.Phone,
			&supplier.Fax,
			&supplier.HomePage,
		); err != nil {
			log.Error().Err(err).Msg("error scanning supplier")
			return nil, fmt.Errorf("error scanning supplier: %w", err)
		}
		suppliers = append(suppliers, supplier)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("rows error on suppliers")
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return suppliers, nil
}

func (r *SupplierRepository) GetSupplierByID(ctx context.Context, id int) (models.Supplier, error) {
	var supplier models.Supplier
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			SupplierID,
			CompanyName,
			ContactName,
			ContactTitle,
			Address,
			City,
			Region,
			PostalCode,
			Country,
			Phone,
			Fax,
			HomePage
		FROM Suppliers
		WHERE SupplierID = ?
	`, id).Scan(
		&supplier.SupplierID,
		&supplier.CompanyName,
		&supplier.ContactName,
		&supplier.ContactTitle,
		&supplier.Address,
		&supplier.City,
		&supplier.Region,
		&supplier.PostalCode,
		&supplier.Country,
		&supplier.Phone,
		&supplier.Fax,
		&supplier.HomePage,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return supplier, fmt.Errorf("supplier with ID %d not found",
				id)
		}
		log.Error().Err(err).Msg("error fetching supplier by ID")
		return supplier, fmt.Errorf("error fetching supplier by ID: %w", err)
	}
	return supplier, nil
}

func (r *SupplierRepository) GetSuppliersByProductID(ctx context.Context, productID int) ([]models.Supplier, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			s.SupplierID,
			s.CompanyName,
			s.ContactName,
			s.ContactTitle,
			s.Address,
			s.City,
			s.Region,
			s.PostalCode,
			s.Country,
			s.Phone,
			s.Fax,
			s.HomePage
		FROM Suppliers s
		JOIN Products p ON s.SupplierID = p.SupplierID
		WHERE p.ProductID = ?
	`, productID)
	if err != nil {
		log.Error().Err(err).Msg("error fetching suppliers by product ID")
		return nil, fmt.Errorf("error fetching suppliers by product ID: %w", err)
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		if err := rows.Scan(
			&supplier.SupplierID,
			&supplier.CompanyName,
			&supplier.ContactName,
			&supplier.ContactTitle,
			&supplier.Address,
			&supplier.City,
			&supplier.Region,
			&supplier.PostalCode,
			&supplier.Country,
			&supplier.Phone,
			&supplier.Fax,
			&supplier.HomePage,
		); err != nil {
			log.Error().Err(err).Msg("error scanning supplier by product ID")
			return nil, fmt.Errorf("error scanning supplier: %w", err)
		}
		suppliers = append(suppliers, supplier)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("rows error on suppliers by product ID")
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return suppliers, nil
}

func (r *ProductRepository) CreateSupplier(ctx context.Context, s *models.Supplier) error {
	result, err := r.DB.ExecContext(ctx, `
		INSERT INTO Suppliers (
			CompanyName,
			ContactName,
			ContactTitle,
			Address,
			City,
			Region,
			PostalCode,
			Country,
			Phone,
			Fax,
			HomePage
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		s.CompanyName,
		s.ContactName,
		s.ContactTitle,
		s.Address,
		s.City,
		s.Region,
		s.PostalCode,
		s.Country,
		s.Phone,
		s.Fax,
		s.HomePage,
	)
	if err != nil {
		log.Error().Err(err).Msg("error creating supplier")
		return fmt.Errorf("error creating supplier: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("error fetching last insert ID for supplier")
		return fmt.Errorf("error fetching last insert ID: %w", err)
	}
	s.SupplierID = id
	return nil
}

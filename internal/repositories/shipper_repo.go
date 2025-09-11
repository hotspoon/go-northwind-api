package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"northwind-api/internal/models"

	"github.com/rs/zerolog/log"
)

type ShipperRepository struct {
	DB *sql.DB
}

func (r *ShipperRepository) GetAllShippers(ctx context.Context) ([]models.Shipper, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			ShipperID, CompanyName, Phone
		FROM Shippers
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query shippers")
		return nil, fmt.Errorf("error fetching shippers: %w", err)
	}
	defer rows.Close()

	var shippers []models.Shipper
	for rows.Next() {
		var shipper models.Shipper
		if err := rows.Scan(&shipper.ShipperID, &shipper.CompanyName, &shipper.Phone); err != nil {
			log.Error().Err(err).Msg("failed to scan shipper")
			return nil, fmt.Errorf("error scanning shipper: %w", err)
		}
		shippers = append(shippers, shipper)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over shippers")
		return nil, fmt.Errorf("error iterating over shippers: %w", err)
	}
	return shippers, nil
}

func (r *ShipperRepository) GetShipperByID(ctx context.Context, id int) (models.Shipper, error) {
	var shipper models.Shipper
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			ShipperID, CompanyName, Phone
		FROM Shippers
		WHERE ShipperID = ?
	`, id).Scan(&shipper.ShipperID, &shipper.CompanyName, &shipper.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return shipper, fmt.Errorf("shipper with ID %d not found", id)
		}
		log.Error().Err(err).Msg("failed to query shipper by ID")
		return shipper, fmt.Errorf("error fetching shipper by ID: %w", err)
	}
	return shipper, nil
}

func (r *ShipperRepository) CreateShipper(ctx context.Context, shipper models.Shipper) (int64, error) {
	result, err := r.DB.ExecContext(
		ctx,
		"INSERT INTO Shippers (CompanyName, Phone) VALUES (?, ?)",
		shipper.CompanyName, shipper.Phone,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error creating shipper")
		return 0, fmt.Errorf("error creating shipper: %w", err)
	}
	return result.LastInsertId()
}

func (r *ShipperRepository) UpdateShipper(ctx context.Context, shipper *models.Shipper) error {
	_, err := r.DB.ExecContext(
		ctx,
		"UPDATE Shippers SET CompanyName = ?, Phone = ? WHERE ShipperID = ?",
		shipper.CompanyName, shipper.Phone, shipper.ShipperID,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error updating shipper")
		return fmt.Errorf("error updating shipper: %w", err)
	}
	return nil
}

func (r *ShipperRepository) DeleteShipper(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM Shippers WHERE ShipperID = ?", id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting shipper")
		return fmt.Errorf("error deleting shipper: %w", err)
	}
	return nil
}

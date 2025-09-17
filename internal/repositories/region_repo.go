package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"northwind-api/internal/models"

	"github.com/rs/zerolog/log"
)

type RegionRepository struct {
	DB *sql.DB
}

func (r *RegionRepository) GetAllRegions(ctx context.Context) ([]models.Region, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT RegionID, RegionDescription
		FROM Regions
		ORDER BY RegionID ASC`)
	if err != nil {
		log.Error().Err(err).Msg("error querying all regions")
		return nil, fmt.Errorf("error querying all regions: %w", err)
	}
	defer rows.Close()

	var regions []models.Region
	for rows.Next() {
		var region models.Region
		if err := rows.Scan(&region.RegionID, &region.RegionDescription); err != nil {
			log.Error().Err(err).Msg("error scanning region row")
			return nil, fmt.Errorf("error scanning region row: %w", err)
		}
		regions = append(regions, region)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating region rows")
		return nil, fmt.Errorf("error iterating region rows: %w", err)
	}
	return regions, nil
}

func (r *RegionRepository) GetRegionsByID(ctx context.Context, id int) (*models.Region, error) {
	var region models.Region
	err := r.DB.QueryRowContext(ctx, `
		SELECT RegionID, RegionDescription
		FROM Regions
		WHERE RegionID = ?`, id).Scan(&region.RegionID, &region.RegionDescription)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("region with ID %d not found", id)
		}
		log.Error().Err(err).Msg("error querying region by ID")
		return nil, fmt.Errorf("error querying region by ID: %w", err)
	}
	return &region, nil
}

// GET /territories/{id}/employees → karyawan di territory tertentu
func (r *RegionRepository) GetEmployeesByTerritoryID(ctx context.Context, id string) ([]models.Employee, error) {
	// how to log
	log.Info().Msgf("Fetching employees for territory ID: %s", id)
	rows, err := r.DB.QueryContext(ctx, `
		SELECT e.EmployeeID, e.LastName, e.FirstName, e.Title, e.TitleOfCourtesy, e.BirthDate, e.HireDate,
		       e.Address, e.City, e.Region, e.PostalCode, e.Country, e.HomePhone, e.Extension, e.Photo,
		       e.Notes, e.ReportsTo, e.PhotoPath
		FROM Employees e
		JOIN EmployeeTerritories et ON e.EmployeeID = et.EmployeeID
		WHERE et.TerritoryID = ?`, id)
	if err != nil {
		log.Error().Err(err).Msg("error querying employees by territory ID")
		return nil, fmt.Errorf("error querying employees by territory ID: %w", err)
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.EmployeeID, &emp.LastName, &emp.FirstName, &emp.Title, &emp.TitleOfCourtesy,
			&emp.BirthDate, &emp.HireDate, &emp.Address, &emp.City, &emp.Region, &emp.PostalCode,
			&emp.Country, &emp.HomePhone, &emp.Extension, &emp.Photo, &emp.Notes, &emp.ReportsTo,
			&emp.PhotoPath); err != nil {
			log.Error().Err(err).Msg("error scanning employee row")
			return nil, fmt.Errorf("error scanning employee row: %w", err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating employee rows")
		return nil, fmt.Errorf("error iterating employee rows: %w", err)
	}
	return employees, nil
}

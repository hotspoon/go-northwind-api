package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"northwind-api/internal/models"

	// "northwind-api/internal/utils"

	"github.com/rs/zerolog/log"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, emp models.Employee) (int64, error) {
	result, err := r.DB.ExecContext(
		ctx,
		`INSERT INTO Employee
		(LastName, FirstName, Title, TitleOfCourtesy, BirthDate, HireDate, Address, City, Region, PostalCode, Country, HomePhone, Extension, Photo, Notes, ReportsTo, PhotoPath)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		emp.LastName, emp.FirstName, emp.Title, emp.TitleOfCourtesy, emp.BirthDate, emp.HireDate,
		emp.Address, emp.City, emp.Region, emp.PostalCode, emp.Country, emp.HomePhone,
		emp.Extension, emp.Photo, emp.Notes, emp.ReportsTo, emp.PhotoPath,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error creating employee")
		return 0, fmt.Errorf("error creating employee: %w", err)
	}
	return result.LastInsertId()
}

func (r *EmployeeRepository) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			EmployeeID, LastName, FirstName, Title, TitleOfCourtesy, BirthDate, HireDate,
			Address, City, Region, PostalCode, Country, HomePhone, Extension, Photo, Notes, ReportsTo, PhotoPath
		FROM Employee
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query employees")
		return nil, fmt.Errorf("error fetching employees: %w", err)
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(
			&employee.EmployeeID,
			&employee.LastName,
			&employee.FirstName,
			&employee.Title,
			&employee.TitleOfCourtesy,
			&employee.BirthDate,
			&employee.HireDate,
			&employee.Address,
			&employee.City,
			&employee.Region,
			&employee.PostalCode,
			&employee.Country,
			&employee.HomePhone,
			&employee.Extension,
			&employee.Photo,
			&employee.Notes,
			&employee.ReportsTo,
			&employee.PhotoPath,
		); err != nil {
			log.Error().Err(err).Msg("failed to scan employee")
			return nil, fmt.Errorf("error scanning employee: %w", err)
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("rows error")
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return employees, nil
}

func (r *EmployeeRepository) GetEmployeeByID(ctx context.Context, id int) (models.Employee, error) {
	var employee models.Employee
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			EmployeeID, LastName, FirstName, Title, TitleOfCourtesy, BirthDate, HireDate,
			Address, City, Region, PostalCode, Country, HomePhone, Extension, Photo, Notes, ReportsTo, PhotoPath
		FROM Employee
		WHERE EmployeeID = ?
	`, id).Scan(
		&employee.EmployeeID,
		&employee.LastName,
		&employee.FirstName,
		&employee.Title,
		&employee.TitleOfCourtesy,
		&employee.BirthDate,
		&employee.HireDate,
		&employee.Address,
		&employee.City,
		&employee.Region,
		&employee.PostalCode,
		&employee.Country,
		&employee.HomePhone,
		&employee.Extension,
		&employee.Photo,
		&employee.Notes,
		&employee.ReportsTo,
		&employee.PhotoPath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Employee{}, fmt.Errorf("employee with ID %d not found", id)
		}
		log.Error().Err(err).Msg("failed to query employee by ID")
		return models.Employee{}, fmt.Errorf("error fetching employee by ID: %w", err)
	}
	return employee, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, emp *models.Employee) error {
	_, err := r.DB.ExecContext(
		ctx,
		`UPDATE Employee SET
			LastName = ?, FirstName = ?, Title = ?, TitleOfCourtesy = ?, BirthDate = ?, HireDate = ?,
			Address = ?, City = ?, Region = ?, PostalCode = ?, Country = ?, HomePhone = ?, Extension = ?,
			Photo = ?, Notes = ?, ReportsTo = ?, PhotoPath = ?
		WHERE EmployeeID = ?`,
		emp.LastName, emp.FirstName, emp.Title, emp.TitleOfCourtesy, emp.BirthDate, emp.HireDate,
		emp.Address, emp.City, emp.Region, emp.PostalCode, emp.Country, emp.HomePhone,
		emp.Extension, emp.Photo, emp.Notes, emp.ReportsTo, emp.PhotoPath, emp.EmployeeID,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error updating employee")
		return fmt.Errorf("error updating employee: %w", err)
	}

	return nil
}

func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	result, err := r.DB.ExecContext(ctx, "DELETE FROM Employee WHERE EmployeeID = ?", id)
	if err != nil {
		log.Error().Err(err).Int("employee_id", id).Msg("Error deleting employee")
		return fmt.Errorf("error deleting employee: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("employee_id", id).Msg("Error fetching rows affected for employee delete")
		return fmt.Errorf("error fetching rows affected: %w", err)
	}
	if rowsAffected == 0 {
		log.Warn().Int("employee_id", id).Msg("No employee found to delete")
		return fmt.Errorf("no employee found with ID %d", id)
	}

	log.Info().Int("employee_id", id).Msg("Employee deleted")
	return nil
}

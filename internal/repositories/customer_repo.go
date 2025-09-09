package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"northwind-api/internal/models"
	"northwind-api/internal/utils"

	"github.com/rs/zerolog/log"
)

type CustomerRepository struct {
	DB *sql.DB
}

func (r *CustomerRepository) GetAllCustomers(ctx context.Context) ([]models.Customer, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT
			COALESCE(CustomerID, '') AS CustomerID,
			COALESCE(CompanyName, '') AS CompanyName,
			COALESCE(ContactName, '') AS ContactName,
			COALESCE(ContactTitle, '') AS ContactTitle,
			COALESCE(Address, '') AS Address,
			COALESCE(City, '') AS City,
			COALESCE(Region, '') AS Region,
			COALESCE(PostalCode, '') AS PostalCode,
			COALESCE(Country, '') AS Country,
			COALESCE(Phone, '') AS Phone,
			COALESCE(Fax, '') AS Fax
		FROM Customers
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to query customers")
		return nil, fmt.Errorf("error fetching customers: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(
			&customer.CustomerID,
			&customer.CompanyName,
			&customer.ContactName,
			&customer.ContactTitle,
			&customer.Address,
			&customer.City,
			&customer.Region,
			&customer.PostalCode,
			&customer.Country,
			&customer.Phone,
			&customer.Fax,
		); err != nil {
			log.Error().Err(err).Msg("failed to scan customer")
			return nil, fmt.Errorf("error scanning customer: %w", err)
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over customers")
		return nil, fmt.Errorf("error iterating over customers: %w", err)
	}
	return customers, nil
}

func (r *CustomerRepository) GetCustomerByID(ctx context.Context, id string) (models.Customer, error) {
	var customer models.Customer
	err := r.DB.QueryRowContext(ctx, `
		SELECT
			COALESCE(CustomerID, '') AS CustomerID,
			COALESCE(CompanyName, '') AS CompanyName,
			COALESCE(ContactName, '') AS ContactName,
			COALESCE(ContactTitle, '') AS ContactTitle,
			COALESCE(Address, '') AS Address,
			COALESCE(City, '') AS City,
			COALESCE(Region, '') AS Region,
			COALESCE(PostalCode, '') AS PostalCode,
			COALESCE(Country, '') AS Country,
			COALESCE(Phone, '') AS Phone,
			COALESCE(Fax, '') AS Fax
		FROM Customers
		WHERE CustomerID = ?
	`, id).Scan(
		&customer.CustomerID,
		&customer.CompanyName,
		&customer.ContactName,
		&customer.ContactTitle,
		&customer.Address,
		&customer.City,
		&customer.Region,
		&customer.PostalCode,
		&customer.Country,
		&customer.Phone,
		&customer.Fax,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Customer{}, fmt.Errorf("customer with ID %s not found", id)
		}
		log.Error().Err(err).Msg("failed to query customer by ID")
		return models.Customer{}, fmt.Errorf("error fetching customer by ID: %w", err)
	}
	return customer, nil
}

func (h *CustomerRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (string, error) {
	// generate ID
	customer.CustomerID = utils.GenerateCustomerID()
	log.Debug().
		Str("customer_id", customer.CustomerID).
		Str("company_name", customer.CompanyName).
		Str("contact_name", customer.ContactName).
		Msg("Creating customer")

	_, err := h.DB.ExecContext(
		ctx,
		`INSERT INTO Customers
            (CustomerID, CompanyName, ContactName, ContactTitle, Address, City, Region, PostalCode, Country, Phone, Fax)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		customer.CustomerID,
		customer.CompanyName,
		customer.ContactName,
		customer.ContactTitle,
		customer.Address,
		customer.City,
		customer.Region,
		customer.PostalCode,
		customer.Country,
		customer.Phone,
		customer.Fax,
	)
	if err != nil {
		log.Error().Err(err).
			Str("customer_id", customer.CustomerID).
			Msg("Error creating customer")
		return "", fmt.Errorf("error creating customer: %w", err)
	}

	log.Info().
		Str("customer_id", customer.CustomerID).
		Msg("Customer created")
	return customer.CustomerID, nil
}

func (r *CustomerRepository) UpdateCustomer(ctx context.Context, customer *models.Customer) error {
	log.Debug().
		Str("customer_id", customer.CustomerID).
		Str("company_name", customer.CompanyName).
		Str("contact_name", customer.ContactName).
		Msg("Updating customer")

	result, err := r.DB.ExecContext(
		ctx,
		`UPDATE Customers SET
			CompanyName = ?, ContactName = ?, ContactTitle = ?, Address = ?, City = ?,
			Region = ?, PostalCode = ?, Country = ?, Phone = ?, Fax = ?
		 WHERE CustomerID = ?`,
		customer.CompanyName,
		customer.ContactName,
		customer.ContactTitle,
		customer.Address,
		customer.City,
		customer.Region,
		customer.PostalCode,
		customer.Country,
		customer.Phone,
		customer.Fax,
		customer.CustomerID,
	)
	if err != nil {
		log.Error().Err(err).
			Str("customer_id", customer.CustomerID).
			Msg("Error updating customer")
		return fmt.Errorf("error updating customer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).
			Str("customer_id", customer.CustomerID).
			Msg("Error fetching rows affected for customer update")
		return fmt.Errorf("error fetching rows affected: %w", err)
	}
	if rowsAffected == 0 {
		log.Warn().
			Str("customer_id", customer.CustomerID).
			Msg("No customer found to update")
		return fmt.Errorf("no customer found with ID %s", customer.CustomerID)
	}

	log.Info().
		Str("customer_id", customer.CustomerID).
		Msg("Customer updated")
	return nil
}

func (r *CustomerRepository) DeleteCustomer(ctx context.Context, id string) error {
	result, err := r.DB.ExecContext(ctx, "DELETE FROM Customers WHERE CustomerID = ?", id)
	if err != nil {
		log.Error().Err(err).Str("customer_id", id).Msg("Error deleting customer")
		return fmt.Errorf("error deleting customer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Str("customer_id", id).Msg("Error fetching rows affected for customer delete")
		return fmt.Errorf("error fetching rows affected: %w", err)
	}
	if rowsAffected == 0 {
		log.Warn().Str("customer_id", id).Msg("No customer found to delete")
		return fmt.Errorf("no customer found with ID %s", id)
	}

	log.Info().Str("customer_id", id).Msg("Customer deleted")
	return nil
}

package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"northwind-api/internal/models"

	"github.com/rs/zerolog/log"
)

type CategoryRepository struct {
	DB *sql.DB
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, c *models.Category) (int64, error) {
	result, err := r.DB.ExecContext(ctx, `
		INSERT INTO Categories (CategoryName, Description, Picture)
		VALUES (?, ?, ?)
	`, c.CategoryName, c.Description, c.Picture)
	if err != nil {
		log.Error().Err(err).Msg("error creating category")
		return 0, fmt.Errorf("error creating category: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("error getting last insert id for category")
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}
	return id, nil
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT CategoryID, CategoryName, Description
		FROM Categories
	`)
	if err != nil {
		log.Error().Err(err).Msg("error fetching categories")
		return nil, fmt.Errorf("error fetching categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName, &category.Description); err != nil {
			log.Error().Err(err).Msg("error scanning category row")
			return nil, fmt.Errorf("error scanning category row: %w", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("error iterating over category rows")
		return nil, fmt.Errorf("error iterating over category rows: %w", err)
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id int) (models.Category, error) {
	var category models.Category
	err := r.DB.QueryRowContext(ctx, `
		SELECT CategoryID, CategoryName, Description
		FROM Categories
		WHERE CategoryID = ?
	`, id).Scan(&category.CategoryID, &category.CategoryName, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Int("id", id).Msg("category not found")
			return category, fmt.Errorf("category not found")
		}
		log.Error().Err(err).Int("id", id).Msg("error fetching category by ID")
		return category, fmt.Errorf("error fetching category by ID: %w", err)
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, c *models.Category) error {
	result, err := r.DB.ExecContext(ctx, `
		UPDATE Categories
		SET CategoryName = ?, Description = ?, Picture = ?
		WHERE CategoryID = ?
	`, c.CategoryName, c.Description, c.Picture, c.CategoryID)
	if err != nil {
		log.Error().Err(err).Int64("id", c.CategoryID).Msg("error updating category")
		return fmt.Errorf("error updating category: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int64("id", c.CategoryID).Msg("error getting rows affected for category update")
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		log.Warn().Int64("id", c.CategoryID).Msg("no category found to update")
		return fmt.Errorf("category not found")
	}
	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	result, err := r.DB.ExecContext(ctx, `DELETE FROM Categories WHERE CategoryID = ?`, id)
	if err != nil {
		log.Error().Err(err).Int("id", id).Msg("error deleting category")
		return fmt.Errorf("error deleting category: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("id", id).Msg("error getting rows affected for category delete")
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		log.Warn().Int("id", id).Msg("no category found to delete")
		return fmt.Errorf("category not found")
	}
	return nil
}

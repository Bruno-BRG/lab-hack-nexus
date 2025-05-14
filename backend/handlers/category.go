package handlers

import (
	"database/sql"
	"lab-hack-nexus/backend/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Icon        string `json:"icon"`
	IsActive    bool   `json:"is_active"`
}

// CreateCategory handles the creation of a new category
func CreateCategory(c *fiber.Ctx) error {
	var category Category
	
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Generate slug from name
	category.Slug = strings.ToLower(strings.ReplaceAll(category.Name, " ", "-"))

	result, err := config.DB.Exec(`
		INSERT INTO categories (name, description, slug, icon, is_active)
		VALUES (?, ?, ?, ?, ?)
	`, category.Name, category.Description, category.Slug, category.Icon, true)
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	id, _ := result.LastInsertId()
	category.ID = id

	return c.Status(201).JSON(category)
}

// GetCategories returns all active categories
func GetCategories(c *fiber.Ctx) error {
	rows, err := config.DB.Query(`
		SELECT id, name, description, slug, icon, is_active 
		FROM categories
		WHERE is_active = true
		ORDER BY name ASC
	`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch categories",
		})
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.Slug,
			&category.Icon,
			&category.IsActive,
		); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to scan category",
			})
		}
		categories = append(categories, category)
	}

	return c.JSON(categories)
}

// GetCategoryBySlug returns a specific category by its slug
func GetCategoryBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var category Category

	err := config.DB.QueryRow(`
		SELECT id, name, description, slug, icon, is_active
		FROM categories
		WHERE slug = ? AND is_active = true
	`, slug).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.Slug,
		&category.Icon,
		&category.IsActive,
	)

	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{
			"error": "Category not found",
		})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch category",
		})
	}

	return c.JSON(category)
}

// UpdateCategory updates an existing category
func UpdateCategory(c *fiber.Ctx) error {
	var category Category
	
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	category.Slug = strings.ToLower(strings.ReplaceAll(category.Name, " ", "-"))

	_, err := config.DB.Exec(`
		UPDATE categories
		SET name = ?, description = ?, slug = ?, icon = ?
		WHERE id = ?
	`, category.Name, category.Description, category.Slug, category.Icon, category.ID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}

	return c.JSON(category)
}

// DeleteCategory soft deletes a category by setting is_active to false
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.DB.Exec(`
		UPDATE categories
		SET is_active = false
		WHERE id = ?
	`, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}

	return c.SendStatus(204)
}

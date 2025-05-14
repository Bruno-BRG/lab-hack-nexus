package handlers

import (
	"database/sql"
	"lab-hack-nexus/backend/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	AuthorID     string `json:"author_id"`
	CategoryID   int64  `json:"category_id"`
	ThumbnailURL string `json:"thumbnail_url"`
	IsPublished  bool   `json:"is_published"`
}

// CreatePost creates a new post
func CreatePost(c *fiber.Ctx) error {
	var post Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Generate slug from title
	post.Slug = strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))

	result, err := config.DB.Exec(`
		INSERT INTO posts (title, content, slug, summary, author_id, category_id, thumbnail_url, is_published)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, post.Title, post.Content, post.Slug, post.Summary, post.AuthorID, post.CategoryID, post.ThumbnailURL, post.IsPublished)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}

	id, _ := result.LastInsertId()
	post.ID = id

	return c.Status(201).JSON(post)
}

// GetPosts returns all published posts, optionally filtered by category
func GetPosts(c *fiber.Ctx) error {
	categorySlug := c.Query("category")
	var rows *sql.Rows
	var err error

	if categorySlug != "" {
		rows, err = config.DB.Query(`
			SELECT p.id, p.title, p.content, p.slug, p.summary, p.author_id, p.category_id, p.thumbnail_url, p.is_published
			FROM posts p
			JOIN categories c ON p.category_id = c.id
			WHERE p.is_published = true AND c.slug = ?
			ORDER BY p.created_at DESC
		`, categorySlug)
	} else {
		rows, err = config.DB.Query(`
			SELECT id, title, content, slug, summary, author_id, category_id, thumbnail_url, is_published
			FROM posts
			WHERE is_published = true
			ORDER BY created_at DESC
		`)
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch posts",
		})
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Slug,
			&post.Summary,
			&post.AuthorID,
			&post.CategoryID,
			&post.ThumbnailURL,
			&post.IsPublished,
		); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to scan post",
			})
		}
		posts = append(posts, post)
	}

	return c.JSON(fiber.Map{
		"data":  posts,
		"count": len(posts),
	})
}

// GetPostBySlug returns a specific post by its slug
func GetPostBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var post Post

	err := config.DB.QueryRow(`
		SELECT id, title, content, slug, summary, author_id, category_id, thumbnail_url, is_published
		FROM posts
		WHERE slug = ? AND is_published = true
	`, slug).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Slug,
		&post.Summary,
		&post.AuthorID,
		&post.CategoryID,
		&post.ThumbnailURL,
		&post.IsPublished,
	)

	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch post",
		})
	}

	return c.JSON(post)
}

// UpdatePost updates an existing post
func UpdatePost(c *fiber.Ctx) error {
	var post Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	post.Slug = strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))

	_, err := config.DB.Exec(`
		UPDATE posts
		SET title = ?, content = ?, slug = ?, summary = ?, category_id = ?, thumbnail_url = ?, is_published = ?
		WHERE id = ?
	`, post.Title, post.Content, post.Slug, post.Summary, post.CategoryID, post.ThumbnailURL, post.IsPublished, post.ID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}

	return c.JSON(post)
}

// DeletePost deletes a post
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.DB.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}

	return c.SendStatus(204)
}

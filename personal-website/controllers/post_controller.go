package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"personal-website/config"
	"personal-website/models"

	"github.com/gin-gonic/gin"
)

// GetPosts returns a list of posts
func GetPosts(c *gin.Context) {
	// Get page number from query
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	perPage := 10

	// Get posts
	pagination, err := models.GetPaginatedPosts(c, config.DB, page, perPage, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, pagination)
}

// GetPost returns a single post
func GetPost(c *gin.Context) {
	// Get post ID from URL
	id := c.Param("id")

	// Get post
	var post models.Post
	if err := config.DB.Preload("Author").Preload("Tags").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if post is published
	if !post.Published {
		// Check if user is logged in and is the author or admin
		userID, exists := c.Get("user_id")
		if !exists || (post.AuthorID != userID.(uint) && userID != uint(1)) { // Assuming admin has ID 1
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
	}

	// Increment view count
	config.DB.Model(&post).Update("view_count", post.ViewCount+1)

	c.JSON(http.StatusOK, post)
}

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
	// Get user ID from context
	userID, _ := c.Get("user_id")

	// Bind JSON
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set author ID
	post.AuthorID = userID.(uint)

	// Generate slug from title
	post.Slug = generateSlug(post.Title)

	// Set published time if post is published
	if post.Published {
		now := time.Now()
		post.PublishedAt = &now
	}

	// Save post
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// UpdatePost updates a post
func UpdatePost(c *gin.Context) {
	// Get post ID from URL
	id := c.Param("id")

	// Get user ID from context
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	// Get post
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if user is the author or admin
	if post.AuthorID != userID.(uint) && role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this post"})
		return
	}

	// Bind JSON
	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if post is being published
	if !post.Published && updatedPost.Published {
		now := time.Now()
		updatedPost.PublishedAt = &now
	}

	// Update post
	if err := config.DB.Model(&post).Updates(updatedPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post
func DeletePost(c *gin.Context) {
	// Get post ID from URL
	id := c.Param("id")

	// Get user ID from context
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	// Get post
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if user is the author or admin
	if post.AuthorID != userID.(uint) && role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this post"})
		return
	}

	// Delete post
	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// Helper function to generate a slug from a title
func generateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)

	return slug
}

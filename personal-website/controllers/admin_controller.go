package controllers

import (
	"net/http"
	"personal-website/config"
	"personal-website/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminDashboard renders the admin dashboard
func AdminDashboard(c *gin.Context) {
	// Get user from context
	user, _ := c.Get("user")

	// Get post count
	var postCount int64
	config.DB.Model(&models.Post{}).Count(&postCount)

	// Get user count
	var userCount int64
	config.DB.Model(&models.User{}).Count(&userCount)

	// Get comment count
	var commentCount int64
	config.DB.Model(&models.Comment{}).Count(&commentCount)

	// Get recent posts
	var recentPosts []models.Post
	config.DB.Order("created_at DESC").Limit(5).Find(&recentPosts)

	c.HTML(http.StatusOK, "admin/dashboard.html", gin.H{
		"title":        "Admin Dashboard",
		"user":         user,
		"postCount":    postCount,
		"userCount":    userCount,
		"commentCount": commentCount,
		"recentPosts":  recentPosts,
	})
}

// AdminPosts renders the admin posts page
func AdminPosts(c *gin.Context) {
	// Get user from context
	user, _ := c.Get("user")

	// Get page number from query
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	perPage := 10

	// Get posts
	pagination, err := models.GetPaginatedPosts(c, config.DB, page, perPage, false)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"title": "Error",
			"error": "Failed to get posts",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/posts.html", gin.H{
		"title":       "Manage Posts",
		"user":        user,
		"posts":       pagination.Posts,
		"pagination":  pagination,
		"currentPage": page,
	})
}

// AdminNewPost renders the new post page
func AdminNewPost(c *gin.Context) {
	// Get user from context
	user, _ := c.Get("user")

	// Get all tags
	var tags []models.Tag
	config.DB.Find(&tags)

	c.HTML(http.StatusOK, "admin/post_form.html", gin.H{
		"title": "New Post",
		"user":  user,
		"tags":  tags,
	})
}

// AdminEditPost renders the edit post page
func AdminEditPost(c *gin.Context) {
	// Get user from context
	user, _ := c.Get("user")

	// Get post ID from URL
	id := c.Param("id")

	// Get post
	var post models.Post
	if err := config.DB.Preload("Tags").First(&post, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"title": "Not Found",
			"error": "Post not found",
		})
		return
	}

	// Get all tags
	var tags []models.Tag
	config.DB.Find(&tags)

	c.HTML(http.StatusOK, "admin/post_form.html", gin.H{
		"title": "Edit Post",
		"user":  user,
		"post":  post,
		"tags":  tags,
	})
}

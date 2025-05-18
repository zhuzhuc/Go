package controllers

import (
	"net/http"
	"strconv"
	"time"

	"personal-website/config"
	"personal-website/models"
	"personal-website/utils"

	"github.com/gin-gonic/gin"
)

// HomePage renders the home page
func HomePage(c *gin.Context) {
	// Get latest posts
	var posts []models.Post
	config.DB.Where("published = ?", true).Order("created_at DESC").Limit(5).Find(&posts)

	// Check if user is logged in
	userID, isLoggedIn := utils.GetUserIDFromSession(c)
	var user models.User
	if isLoggedIn {
		config.DB.First(&user, userID)
	}

	// Add current year for footer
	currentYear := time.Now().Year()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Home",
		"posts":       posts,
		"isLoggedIn":  isLoggedIn,
		"user":        user,
		"currentYear": currentYear,
	})
}

// AboutPage renders the about page
func AboutPage(c *gin.Context) {
	// Check if user is logged in
	userID, isLoggedIn := utils.GetUserIDFromSession(c)
	var user models.User
	if isLoggedIn {
		config.DB.First(&user, userID)
	}

	// Add current year for footer
	currentYear := time.Now().Year()

	c.HTML(http.StatusOK, "about.html", gin.H{
		"title":       "About",
		"isLoggedIn":  isLoggedIn,
		"user":        user,
		"currentYear": currentYear,
	})
}

// ProjectsPage renders the projects page
func ProjectsPage(c *gin.Context) {
	// Check if user is logged in
	userID, isLoggedIn := utils.GetUserIDFromSession(c)
	var user models.User
	if isLoggedIn {
		config.DB.First(&user, userID)
	}

	// Add current year for footer
	currentYear := time.Now().Year()

	c.HTML(http.StatusOK, "projects.html", gin.H{
		"title":       "Projects",
		"isLoggedIn":  isLoggedIn,
		"user":        user,
		"currentYear": currentYear,
	})
}

// BlogListPage renders the blog list page
func BlogListPage(c *gin.Context) {
	// Get page number from query
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	perPage := 10

	// Get posts
	pagination, err := models.GetPaginatedPosts(c, config.DB, page, perPage, true)
	if err != nil {
		// Add current year for footer
		currentYear := time.Now().Year()

		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"title":       "Error",
			"error":       "Failed to get posts",
			"currentYear": currentYear,
		})
		return
	}

	// Check if user is logged in
	userID, isLoggedIn := utils.GetUserIDFromSession(c)
	var user models.User
	if isLoggedIn {
		config.DB.First(&user, userID)
	}

	// Add current year for footer
	currentYear := time.Now().Year()

	c.HTML(http.StatusOK, "blog.html", gin.H{
		"title":       "Blog",
		"posts":       pagination.Posts,
		"pagination":  pagination,
		"isLoggedIn":  isLoggedIn,
		"user":        user,
		"currentPage": page,
		"currentYear": currentYear,
	})
}

// BlogDetailPage renders a single blog post
func BlogDetailPage(c *gin.Context) {
	// Get post ID from URL
	id := c.Param("id")

	// Get post
	var post models.Post
	if err := config.DB.Preload("Author").Preload("Tags").First(&post, id).Error; err != nil {
		// Add current year for footer
		currentYear := time.Now().Year()

		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"title":       "Not Found",
			"error":       "Post not found",
			"currentYear": currentYear,
		})
		return
	}

	// Check if post is published
	if !post.Published {
		// Check if user is logged in and is the author or admin
		userID, isLoggedIn := utils.GetUserIDFromSession(c)
		if !isLoggedIn || (post.AuthorID != userID && userID != 1) { // Assuming admin has ID 1
			// Add current year for footer
			currentYear := time.Now().Year()

			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"title":       "Not Found",
				"error":       "Post not found",
				"currentYear": currentYear,
			})
			return
		}
	}

	// Increment view count
	config.DB.Model(&post).Update("view_count", post.ViewCount+1)

	// Get related posts
	var relatedPosts []models.Post
	config.DB.Where("published = ? AND id != ?", true, post.ID).Order("created_at DESC").Limit(3).Find(&relatedPosts)

	// Check if user is logged in
	userID, isLoggedIn := utils.GetUserIDFromSession(c)
	var user models.User
	if isLoggedIn {
		config.DB.First(&user, userID)
	}

	// Add current year for footer
	currentYear := time.Now().Year()

	c.HTML(http.StatusOK, "blog_detail.html", gin.H{
		"title":        post.Title,
		"post":         post,
		"relatedPosts": relatedPosts,
		"isLoggedIn":   isLoggedIn,
		"user":         user,
		"currentYear":  currentYear,
	})
}

// ContactPage renders the contact page
func ContactPage(c *gin.Context) {
	// Check if user is logged in
	userID, isLoggedIn := utils.GetUserIDFromSession(c)
	var user models.User
	if isLoggedIn {
		config.DB.First(&user, userID)
	}

	// Add current year for footer
	currentYear := time.Now().Year()

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title":       "Contact",
		"isLoggedIn":  isLoggedIn,
		"user":        user,
		"currentYear": currentYear,
	})
}

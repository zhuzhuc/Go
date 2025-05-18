package controllers

import (
	"net/http"

	"personal-website/config"
	"personal-website/models"
	"personal-website/utils"

	"github.com/gin-gonic/gin"
)

// LoginPage renders the login page
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

// RegisterPage renders the registration page
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

// Login handles user login
func Login(c *gin.Context) {
	// Get form data
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate input
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"title": "Login",
			"error": "Username and password are required",
		})
		return
	}

	// Find user
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"title": "Login",
			"error": "Invalid username or password",
		})
		return
	}

	// Check password
	if !user.CheckPassword(password) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"title": "Login",
			"error": "Invalid username or password",
		})
		return
	}

	// Save user session
	utils.SaveUserSession(c, user.ID)

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"title": "Login",
			"error": "Failed to generate authentication token",
		})
		return
	}

	// Set token cookie
	c.SetCookie("jwt", token, 3600*24, "/", "", false, true)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/")
}

// Register handles user registration
func Register(c *gin.Context) {
	// Get form data
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	// Validate input
	if username == "" || email == "" || password == "" {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"title": "Register",
			"error": "All fields are required",
		})
		return
	}

	// Check if passwords match
	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"title": "Register",
			"error": "Passwords do not match",
		})
		return
	}

	// Create user
	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     "user", // Default role
	}

	// Save to database
	if err := config.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"title": "Register",
			"error": "Failed to create user",
		})
		return
	}

	// Redirect to login page
	c.Redirect(http.StatusSeeOther, "/login")
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// Clear user session
	utils.ClearUserSession(c)

	// Clear JWT cookie
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/")
}

// GetUser returns the current user
func GetUser(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	// Get user from database
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Return safe user
	c.JSON(http.StatusOK, user.ToSafeUser())
}

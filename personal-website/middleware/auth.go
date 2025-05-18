package middleware

import (
	"net/http"
	"personal-website/config"
	"personal-website/models"
	"personal-website/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth middleware for JWT authentication
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Check if the header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// Validate token
		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user ID in context
		userID := uint(claims["id"].(float64))
		c.Set("user_id", userID)
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

// AuthRequired middleware for session-based authentication
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from session
		userID, exists := utils.GetUserIDFromSession(c)
		if !exists {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// Get user from database
		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			utils.ClearUserSession(c)
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Set("role", user.Role)

		c.Next()
	}
}

// AdminRequired middleware for admin access
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// First check if user is authenticated
		AuthRequired()(c)
		if c.IsAborted() {
			return
		}

		// Check if user is admin
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"title": "Access Denied",
				"error": "You do not have permission to access this page",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhuzhuc/blog/database"
	"github.com/zhuzhuc/blog/model"
)

// JWT secret key - should match the one in auth controller
const jwtSecret = "your-secret-key" // In production, use environment variable

// AuthRequired is a middleware that checks if the user is authenticated
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusText": "Error",
				"message":    "Authorization header is required",
			})
		}

		// Check if the header starts with "Bearer "
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusText": "Error",
				"message":    "Invalid authorization format",
			})
		}

		// Extract the token
		tokenString := authHeader[7:]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusText": "Error",
				"message":    "Invalid or expired token",
			})
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusText": "Error",
				"message":    "Invalid token claims",
			})
		}

		// Get user ID from claims
		userID := uint(claims["id"].(float64))

		// Verify user exists in database
		var user model.User
		result := database.DBConnection.First(&user, userID)
		if result.Error != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusText": "Error",
				"message":    "User not found",
			})
		}

		// Set user in context
		c.Locals("user", user)
		c.Locals("user_id", user.ID)
		c.Locals("username", user.Username)
		c.Locals("role", user.Role)

		// Continue to the next middleware/handler
		return c.Next()
	}
}

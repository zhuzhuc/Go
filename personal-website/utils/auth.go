package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares a password with a hash
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID uint, username, role string) (string, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret" // Default secret for development
	}

	// Set expiration time
	expirationHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		expirationHours = 24 // Default to 24 hours
	}

	// Create claims
	claims := jwt.MapClaims{
		"id":       userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * time.Duration(expirationHours)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates a JWT token
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret" // Default secret for development
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Get claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

// SaveUserSession saves user information to session
func SaveUserSession(c *gin.Context, userID uint) {
	session := sessions.Default(c)
	session.Set("user_id", userID)
	session.Save()
}

// GetUserIDFromSession gets user ID from session
func GetUserIDFromSession(c *gin.Context) (uint, bool) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		return 0, false
	}
	return userID.(uint), true
}

// ClearUserSession clears user session
func ClearUserSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

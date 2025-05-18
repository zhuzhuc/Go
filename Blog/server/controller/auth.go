package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhuzhuc/blog/database"
	"github.com/zhuzhuc/blog/model"
)

// JWT secret key
const jwtSecret = "your-secret-key" // In production, use environment variable

// Login handles user login
func Login(c *fiber.Ctx) error {
	// Create response context
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Login",
	}

	// Parse request body
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginRequest); err != nil {
		log.Println("Error parsing login request:", err)
		context["statusText"] = "Error"
		context["message"] = "Invalid request format"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Validate input
	if loginRequest.Username == "" || loginRequest.Password == "" {
		context["statusText"] = "Error"
		context["message"] = "Username and password are required"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Find user
	var user model.User
	result := database.DBConnection.Where("username = ?", loginRequest.Username).First(&user)
	if result.Error != nil {
		log.Println("User not found:", loginRequest.Username)
		context["statusText"] = "Error"
		context["message"] = "Invalid username or password"
		return c.Status(fiber.StatusUnauthorized).JSON(context)
	}

	// Check password
	if !user.CheckPassword(loginRequest.Password) {
		log.Println("Invalid password for user:", loginRequest.Username)
		context["statusText"] = "Error"
		context["message"] = "Invalid username or password"
		return c.Status(fiber.StatusUnauthorized).JSON(context)
	}

	// Generate JWT token
	token, err := generateToken(user)
	if err != nil {
		log.Println("Error generating token:", err)
		context["statusText"] = "Error"
		context["message"] = "Authentication error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// Return success response with token and user info
	safeUser := user.ToSafeUser()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "Login successful",
		"token":      token,
		"user":       safeUser,
	})
}

// Register handles user registration
func Register(c *fiber.Ctx) error {
	// Create response context
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Register",
	}

	// Parse request body
	var registerRequest struct {
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	if err := c.BodyParser(&registerRequest); err != nil {
		log.Println("Error parsing register request:", err)
		context["statusText"] = "Error"
		context["message"] = "Invalid request format"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Validate input
	if registerRequest.Username == "" || registerRequest.Email == "" || registerRequest.Password == "" {
		context["statusText"] = "Error"
		context["message"] = "All fields are required"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Check if passwords match
	if registerRequest.Password != registerRequest.ConfirmPassword {
		context["statusText"] = "Error"
		context["message"] = "Passwords do not match"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Check if username already exists
	var existingUser model.User
	result := database.DBConnection.Where("username = ?", registerRequest.Username).First(&existingUser)
	if result.Error == nil {
		context["statusText"] = "Error"
		context["message"] = "Username already exists"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Check if email already exists
	result = database.DBConnection.Where("email = ?", registerRequest.Email).First(&existingUser)
	if result.Error == nil {
		context["statusText"] = "Error"
		context["message"] = "Email already exists"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Create user
	user := model.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		Role:     "user", // Default role
	}

	// Save to database
	if err := database.DBConnection.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		context["statusText"] = "Error"
		context["message"] = "Failed to create user"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// Generate JWT token
	token, err := generateToken(user)
	if err != nil {
		log.Println("Error generating token:", err)
		context["statusText"] = "Error"
		context["message"] = "Authentication error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// Return success response with token and user info
	safeUser := user.ToSafeUser()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "Registration successful",
		"token":      token,
		"user":       safeUser,
	})
}

// generateToken creates a new JWT token for a user
func generateToken(user model.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

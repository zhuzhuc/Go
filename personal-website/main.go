package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"personal-website/config"
	"personal-website/controllers"
	"personal-website/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Initialize database
	config.InitDB()

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize router
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Session middleware
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	router.Use(sessions.Sessions("personal-website", store))

	// Custom template functions
	router.SetFuncMap(template.FuncMap{
		"title": strings.Title,
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"seq": func(start, end int) []int {
			var result []int
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
			return result
		},
	})

	// Load templates and static files
	router.LoadHTMLFiles(
		"templates/index.html",
		"templates/about.html",
		"templates/blog.html",
		"templates/blog_detail.html",
		"templates/contact.html",
		"templates/error.html",
		"templates/login.html",
		"templates/projects.html",
		"templates/register.html",
		"templates/test.html",
		"templates/admin/dashboard.html",
		"templates/admin/posts.html",
		"templates/admin/post_form.html",
	)
	router.Static("/static", "./static")

	// Public routes
	router.GET("/", controllers.HomePage)
	router.GET("/about", controllers.AboutPage)
	router.GET("/projects", controllers.ProjectsPage)
	router.GET("/blog", controllers.BlogListPage)
	router.GET("/blog/:id", controllers.BlogDetailPage)
	router.GET("/contact", controllers.ContactPage)
	router.GET("/login", controllers.LoginPage)
	router.GET("/register", controllers.RegisterPage)
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
	router.GET("/logout", controllers.Logout)

	// Test route
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", gin.H{
			"currentYear": time.Now().Year(),
		})
	})

	// API routes
	api := router.Group("/api")
	{
		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)

		// Protected routes
		authorized := api.Group("/")
		authorized.Use(middleware.JWTAuth())
		{
			authorized.POST("/posts", controllers.CreatePost)
			authorized.PUT("/posts/:id", controllers.UpdatePost)
			authorized.DELETE("/posts/:id", controllers.DeletePost)
			authorized.GET("/user", controllers.GetUser)
		}
	}

	// Admin routes
	admin := router.Group("/admin")
	admin.Use(middleware.AuthRequired())
	{
		admin.GET("/", controllers.AdminDashboard)
		admin.GET("/posts", controllers.AdminPosts)
		admin.GET("/posts/new", controllers.AdminNewPost)
		admin.GET("/posts/edit/:id", controllers.AdminEditPost)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

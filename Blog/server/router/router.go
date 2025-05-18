package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuzhuc/blog/controller"
	"github.com/zhuzhuc/blog/middleware"
)

// SetupRoutes sets up the application routes
func SetupRoutes(app *fiber.App) {
	// Authentication routes
	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)

	// Public blog routes
	app.Get("/", controller.BlogList)
	app.Get("/:id", controller.BlogDetail)

	// Protected routes - require authentication
	protected := app.Group("/")
	protected.Use(middleware.AuthRequired())

	// Protected blog routes
	protected.Post("/", controller.BlogCreate)
	protected.Put("/:id", controller.BlogUpdate)
	protected.Delete("/:id", controller.BlogDelete)

	// 图片上传路由 - 确保路径与前端请求匹配
	app.Post("/upload-image", middleware.AuthRequired(), controller.UploadImage)
}

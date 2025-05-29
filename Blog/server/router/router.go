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

	// 用户资料相关路由 - 使用 /api/user 前缀避免与博客路由冲突
	api := app.Group("/api")
	apiProtected := api.Group("/")
	apiProtected.Use(middleware.AuthRequired())

	// 用户API - 使用修复后的控制器
	user := apiProtected.Group("/user")
	user.Get("/profile", controller.GetUserProfile)
	user.Put("/profile", controller.UpdateProfile)
	user.Put("/password", controller.UpdatePassword)
	user.Get("/posts", controller.GetUserPosts)
	user.Post("/avatar", controller.UploadAvatar)

	// 文件上传路由
	protected.Post("/upload-image", controller.UploadImage)

	// 静态文件服务
	app.Static("/uploads", "./uploads")
	app.Static("/static", "./static")

	// 调试和管理员功能
	debug := app.Group("/debug")
	debug.Post("/reset-password", controller.ResetPassword)
	debug.Post("/user-info", controller.DebugUserInfo)
	debug.Post("/update-avatar", controller.UpdateAvatar)
	debug.Post("/blog-create", controller.DebugBlogCreate)
	debug.Get("/blog-list", controller.DebugBlogList)
}

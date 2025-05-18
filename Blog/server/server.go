package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zhuzhuc/blog/database"
	"github.com/zhuzhuc/blog/middleware"
	"github.com/zhuzhuc/blog/router"
)

func init() {
	// 尝试从多个位置加载 .env 文件
	err := godotenv.Load(".env")
	if err != nil {
		// 尝试从 /etc/.env 加载
		err = godotenv.Load("/etc/.env")
		if err != nil {
			// 尝试从当前目录的父目录加载
			err = godotenv.Load("../.env")
			if err != nil {
				// 如果所有尝试都失败，则输出错误但不中断程序
				log.Println("Warning: Error loading .env file, using default or environment variables")
			}
		}
	}

	// 打印当前工作目录，帮助调试
	dir, _ := os.Getwd()
	log.Printf("Current working directory: %s", dir)

	// 打印环境变量
	log.Println("Database environment variables:")
	log.Printf("db_host: %s", os.Getenv("db_host"))
	log.Printf("db_port: %s", os.Getenv("db_port"))
	log.Printf("db_user: %s", os.Getenv("db_user"))
	log.Printf("db_name: %s", os.Getenv("db_name"))

	// 等待 MySQL 启动
	log.Println("Waiting for MySQL to be ready...")
	time.Sleep(10 * time.Second)

	// 尝试连接数据库
	database.ConnectDB()
}

func main() {
	database.ConnectDB()

	sqlDb, err := database.DBConnection.DB()
	if err != nil {
		panic("failed to get database connection: " + err.Error())
	}
	defer sqlDb.Close()
	app := fiber.New()

	// // Create an instance of logger.Config
	// config := logger.Config{
	// 	// You can customize the configuration here
	// 	Format:     "${time} ${method} ${path} - ${status}\n",
	// 	TimeFormat: "15:04:05",
	// 	TimeZone:   "Local",
	// }

	// Provide a logger.Writer and the config instance
	app.Static("/static", "./static")

	// 添加调试路由，用于检查静态文件
	app.Get("/debug/static", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":    "Static files debug info",
			"static_dir": "./static",
			"static_url": "/static",
		})
	})

	// 添加路由，用于列出静态目录中的文件
	app.Get("/debug/list-files", func(c *fiber.Ctx) error {
		fileList := make([]string, 0)

		err := filepath.Walk("./static", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fileList = append(fileList, path)
			}
			return nil
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"files": fileList,
		})
	})

	// 添加测试路由，用于测试图片上传
	app.Post("/debug/upload", func(c *fiber.Ctx) error {
		// 获取上传的文件
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "No file uploaded: " + err.Error(),
			})
		}

		// 确保目录存在
		os.MkdirAll("./static/uploads", 0o755)

		// 保存文件
		filename := "./static/uploads/test-" + file.Filename
		if err := c.SaveFile(file, filename); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to save file: " + err.Error(),
			})
		}

		// 返回文件路径
		return c.JSON(fiber.Map{
			"message": "File uploaded successfully",
			"path":    "/static/uploads/test-" + file.Filename,
			"url":     "http://localhost:3000/static/uploads/test-" + file.Filename,
		})
	})

	// 使用中间件包中的 SetupMiddleware 函数
	middleware.SetupMiddleware(app)

	// Routes
	router.SetupRoutes(app)
	app.Listen(":9000")
}

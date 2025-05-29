package controller

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/zhuzhuc/blog/database"
	"github.com/zhuzhuc/blog/model"
)

// DebugBlogCreate 用于调试博客创建问题
func DebugBlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Debug Blog Create",
	}

	// 解析请求体
	var debugRequest struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Username string `json:"username"`
	}

	if err := c.BodyParser(&debugRequest); err != nil {
		log.Println("Error parsing debug request:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "无效的请求格式",
		})
	}

	// 验证输入
	if debugRequest.Title == "" || debugRequest.Content == "" || debugRequest.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "标题、内容和用户名不能为空",
		})
	}

	// 查找用户
	var user model.User
	result := database.DBConnection.Where("username = ?", debugRequest.Username).First(&user)
	if result.Error != nil {
		log.Println("用户不存在:", debugRequest.Username)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "用户不存在",
		})
	}

	// 创建博客记录
	blog := model.Blog{
		Title:    debugRequest.Title,
		Post:     debugRequest.Content,
		AuthorID: user.ID,
		Author:   user.Username,
	}

	// 检查数据库连接
	sqlDB, err := database.DBConnection.DB()
	if err != nil {
		log.Println("获取数据库连接失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "数据库连接错误: " + err.Error(),
		})
	}

	// 检查数据库连接是否有效
	if err := sqlDB.Ping(); err != nil {
		log.Println("数据库连接无效:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "数据库连接无效: " + err.Error(),
		})
	}

	// 保存到数据库
	res := database.DBConnection.Create(&blog)
	if res.Error != nil {
		log.Println("保存博客失败:", res.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "保存博客失败: " + res.Error.Error(),
		})
	}

	// 检查上传目录
	uploadDir := "./static/uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		log.Println("上传目录不存在，尝试创建:", uploadDir)
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			log.Println("创建上传目录失败:", err)
			context["uploadDirStatus"] = fmt.Sprintf("创建上传目录失败: %s", err.Error())
		} else {
			context["uploadDirStatus"] = "上传目录已创建"
		}
	} else {
		// 检查目录权限
		testFile := filepath.Join(uploadDir, "test.txt")
		if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
			log.Println("写入测试文件失败:", err)
			context["uploadDirStatus"] = fmt.Sprintf("上传目录权限不足: %s", err.Error())
		} else {
			os.Remove(testFile) // 删除测试文件
			context["uploadDirStatus"] = "上传目录权限正常"
		}
	}

	// 返回成功响应
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "调试博客创建成功",
		"blog":       blog,
		"dbStatus":   "数据库连接正常",
		"uploadDir":  context["uploadDirStatus"],
	})
}

// DebugBlogList 获取所有博客（调试用）
func DebugBlogList(c *fiber.Ctx) error {
	var blogs []model.Blog
	result := database.DBConnection.Find(&blogs)
	if result.Error != nil {
		log.Println("获取博客列表失败:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "获取博客列表失败: " + result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "获取博客列表成功",
		"blogs":      blogs,
		"count":      len(blogs),
	})
}

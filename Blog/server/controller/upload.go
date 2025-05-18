package controller

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UploadImage 处理图片上传
func UploadImage(c *fiber.Ctx) error {
	// 打印请求头，帮助调试
	log.Println("Upload Image Request Headers:")
	log.Println("Authorization:", c.Get("Authorization"))
	log.Println("Content-Type:", c.Get("Content-Type"))

	// 获取用户信息
	user := c.Locals("user")
	log.Println("User from context:", user)

	// 创建上下文
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Image Upload",
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("No file uploaded or error parsing file: ", err)
		context["statusText"] = "Error"
		context["message"] = "No file uploaded or error parsing file: " + err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// 确保文件是图片
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".webp" {
		context["statusText"] = "Error"
		context["message"] = "Only image files are allowed (.jpg, .jpeg, .png, .gif, .webp)"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// 确保 uploads 目录存在
	uploadDir := "./static/uploads"
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		log.Println("Error creating upload directory: ", err)
		context["statusText"] = "Error"
		context["message"] = "Failed to create upload directory: " + err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// 生成安全的文件名
	fileExt := filepath.Ext(file.Filename)
	safeFilename := generateSafeFilename() + fileExt

	// 保存文件
	filename := filepath.Join(uploadDir, safeFilename)
	if err := c.SaveFile(file, filename); err != nil {
		log.Println("Error saving file: ", err)
		context["statusText"] = "Error"
		context["message"] = "Failed to save file: " + err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// 返回文件路径
	imagePath := "/static/uploads/" + safeFilename
	context["message"] = "Image uploaded successfully"
	context["path"] = imagePath
	context["url"] = imagePath

	return c.Status(fiber.StatusOK).JSON(context)
}

// generateSafeFilename 生成一个安全的文件名
func generateSafeFilename() string {
	// 使用时间戳和随机数生成唯一文件名
	timestamp := time.Now().UnixNano()
	r := rand.New(rand.NewSource(timestamp))
	randomNum := r.Intn(10000)
	return fmt.Sprintf("image_%d_%d", timestamp, randomNum)
}

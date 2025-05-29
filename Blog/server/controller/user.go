package controller

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/zhuzhuc/blog/database"
	"github.com/zhuzhuc/blog/model"
	"golang.org/x/crypto/bcrypt"
)

// GetUserProfile 获取用户个人资料
func GetUserProfile(c *fiber.Ctx) error {
	// 从上下文中获取用户信息
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "未授权访问",
		})
	}

	// 返回用户信息（不包含密码）
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "获取用户资料成功",
		"user":       user.ToSafeUser(),
	})
}

// UpdateProfile 更新用户个人资料
func UpdateProfile(c *fiber.Ctx) error {
	// 从上下文中获取用户信息
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "未授权访问",
		})
	}

	// 解析请求体
	var updateRequest struct {
		Bio string `json:"bio"`
	}

	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "无效的请求格式",
		})
	}

	// 从数据库获取完整的用户信息
	var dbUser model.User
	if err := database.DBConnection.First(&dbUser, user.ID).Error; err != nil {
		log.Println("获取用户信息失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "获取用户信息失败",
		})
	}

	// 更新用户信息
	dbUser.Bio = updateRequest.Bio

	// 保存到数据库
	if err := database.DBConnection.Save(&dbUser).Error; err != nil {
		log.Println("更新用户资料失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "更新用户资料失败",
		})
	}

	// 返回更新后的用户信息
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "更新用户资料成功",
		"user":       dbUser.ToSafeUser(),
	})
}

// UpdatePassword 更新用户密码
func UpdatePassword(c *fiber.Ctx) error {
	// 从上下文中获取用户信息
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "未授权访问",
		})
	}

	// 解析请求体
	var passwordRequest struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
	}

	if err := c.BodyParser(&passwordRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "无效的请求格式",
		})
	}

	// 从数据库获取完整的用户信息
	var dbUser model.User
	if err := database.DBConnection.First(&dbUser, user.ID).Error; err != nil {
		log.Println("获取用户信息失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "获取用户信息失败",
		})
	}

	// 验证当前密码
	if !dbUser.CheckPassword(passwordRequest.CurrentPassword) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "当前密码不正确",
		})
	}

	// 更新密码 - 直接使用bcrypt手动哈希密码，避免依赖BeforeSave钩子
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordRequest.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("密码哈希失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "密码更新失败: 无法哈希密码",
		})
	}

	// 设置哈希后的密码
	dbUser.Password = string(hashedPassword)

	// 打印日志
	log.Printf("用户 %d 正在更新密码，新密码哈希: %s", dbUser.ID, string(hashedPassword))

	// 保存到数据库
	if err := database.DBConnection.Save(&dbUser).Error; err != nil {
		log.Println("更新密码失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "更新密码失败",
		})
	}

	// 从auth.go导入generateToken函数
	// 生成新的JWT令牌 - 直接使用auth.go中的逻辑
	// 创建token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = dbUser.ID
	claims["username"] = dbUser.Username
	claims["email"] = dbUser.Email
	claims["role"] = dbUser.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token有效期24小时

	// 生成签名token
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		log.Println("生成令牌失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "生成令牌失败",
		})
	}

	// 返回成功响应，包含新的令牌和用户信息
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "密码更新成功",
		"token":      tokenString,
		"user":       dbUser.ToSafeUser(),
	})
}

// UploadAvatar 上传用户头像
func UploadAvatar(c *fiber.Ctx) error {
	// 从上下文中获取用户信息
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "未授权访问",
		})
	}

	// 获取上传的文件
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "未找到上传的文件",
		})
	}

	// 验证文件类型
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "不支持的文件类型，请上传 jpg、jpeg、png 或 gif 格式的图片",
		})
	}

	// 创建上传目录
	uploadDir := "./uploads/avatars"
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		log.Println("创建上传目录失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "服务器错误",
		})
	}

	// 生成唯一文件名
	filename := fmt.Sprintf("%d_%s%s", user.ID, uuid.New().String(), ext)
	filepath := fmt.Sprintf("%s/%s", uploadDir, filename)

	// 保存文件
	if err := c.SaveFile(file, filepath); err != nil {
		log.Println("保存文件失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "保存文件失败",
		})
	}

	// 从数据库获取完整的用户信息
	var dbUser model.User
	if err := database.DBConnection.First(&dbUser, user.ID).Error; err != nil {
		log.Println("获取用户信息失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "获取用户信息失败",
		})
	}

	// 更新用户头像路径 - 确保路径以斜杠开头
	avatarURL := fmt.Sprintf("/uploads/avatars/%s", filename)
	dbUser.Avatar = avatarURL

	// 打印日志
	log.Printf("用户 %d 头像已更新: %s", dbUser.ID, avatarURL)

	// 保存到数据库
	if err := database.DBConnection.Save(&dbUser).Error; err != nil {
		log.Println("更新用户头像失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "更新用户头像失败",
		})
	}

	// 生成新的JWT令牌
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = dbUser.ID
	claims["username"] = dbUser.Username
	claims["email"] = dbUser.Email
	claims["role"] = dbUser.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token有效期24小时

	// 生成签名token
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		log.Println("生成令牌失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "生成令牌失败",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "头像上传成功",
		"avatarUrl":  avatarURL,
		"token":      tokenString,
		"user":       dbUser.ToSafeUser(),
	})
}

// GetUserPosts 获取用户发布的文章
func GetUserPosts(c *fiber.Ctx) error {
	// 从上下文中获取用户信息
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "未授权访问",
		})
	}

	// 查询用户发布的文章
	var posts []model.Blog
	if err := database.DBConnection.Where("author_id = ?", user.ID).Find(&posts).Error; err != nil {
		log.Println("获取用户文章失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "获取用户文章失败",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "获取用户文章成功",
		"posts":      posts,
	})
}

package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/zhuzhuc/blog/database"
	"github.com/zhuzhuc/blog/model"
	"golang.org/x/crypto/bcrypt"
)

// ResetPassword 重置用户密码（管理员功能）
func ResetPassword(c *fiber.Ctx) error {
	// 解析请求体
	var resetRequest struct {
		Username    string `json:"username"`
		NewPassword string `json:"newPassword"`
	}

	if err := c.BodyParser(&resetRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "无效的请求格式",
		})
	}

	// 验证输入
	if resetRequest.Username == "" || resetRequest.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "用户名和新密码不能为空",
		})
	}

	// 查找用户
	var user model.User
	result := database.DBConnection.Where("username = ?", resetRequest.Username).First(&user)
	if result.Error != nil {
		log.Println("用户不存在:", resetRequest.Username)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "用户不存在",
		})
	}

	// 手动哈希新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetRequest.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("密码哈希失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "密码重置失败: 无法哈希密码",
		})
	}

	// 直接更新数据库中的密码字段
	result = database.DBConnection.Model(&user).Update("password", string(hashedPassword))
	if result.Error != nil {
		log.Println("更新密码失败:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "密码重置失败: 数据库更新错误",
		})
	}

	// 打印日志
	log.Printf("用户 %s (ID: %d) 的密码已重置，新密码哈希: %s", user.Username, user.ID, string(hashedPassword))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "密码重置成功",
	})
}

// DebugUserInfo 获取用户信息（调试用）
func DebugUserInfo(c *fiber.Ctx) error {
	// 解析请求体
	var debugRequest struct {
		Username string `json:"username"`
	}

	if err := c.BodyParser(&debugRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "无效的请求格式",
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

	// 检查头像URL
	avatarStatus := "无头像"
	if user.Avatar != "" {
		avatarStatus = fmt.Sprintf("头像URL: %s", user.Avatar)

		// 检查文件是否存在
		filePath := "." + user.Avatar // 转换为相对文件路径
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			avatarStatus += " (文件不存在)"
		} else {
			avatarStatus += " (文件存在)"
		}
	}

	// 返回用户信息（不包含密码）
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "获取用户信息成功",
		"user": fiber.Map{
			"id":           user.ID,
			"username":     user.Username,
			"email":        user.Email,
			"avatar":       user.Avatar,
			"avatarStatus": avatarStatus,
			"bio":          user.Bio,
			"role":         user.Role,
			"createdAt":    user.CreatedAt,
			"updatedAt":    user.UpdatedAt,
		},
	})
}

// UpdateAvatar 更新用户头像（调试用）
func UpdateAvatar(c *fiber.Ctx) error {
	// 解析请求体
	var updateRequest struct {
		Username  string `json:"username"`
		AvatarURL string `json:"avatarUrl"`
	}

	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "无效的请求格式",
		})
	}

	// 验证输入
	if updateRequest.Username == "" || updateRequest.AvatarURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "用户名和头像URL不能为空",
		})
	}

	// 查找用户
	var user model.User
	result := database.DBConnection.Where("username = ?", updateRequest.Username).First(&user)
	if result.Error != nil {
		log.Println("用户不存在:", updateRequest.Username)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "用户不存在",
		})
	}

	// 更新头像URL
	user.Avatar = updateRequest.AvatarURL

	// 保存到数据库
	if err := database.DBConnection.Save(&user).Error; err != nil {
		log.Println("更新头像失败:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"message":    "更新头像失败: 数据库更新错误",
		})
	}

	// 打印日志
	log.Printf("用户 %s (ID: %d) 的头像已更新为: %s", user.Username, user.ID, updateRequest.AvatarURL)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "OK",
		"message":    "头像更新成功",
		"user":       user.ToSafeUser(),
	})
}

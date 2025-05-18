package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zhuzhuc/blog/database"
	"github.com/zhuzhuc/blog/model"
)

// Blog List
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog List",
	}

	// time.Sleep(time.Millisecond * 100)
	db := database.DBConnection

	var records []model.Blog
	db.Find(&records)
	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)
}

// Add a Blog into database
// blog detail page
func BlogDetail(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"message":    "",
	}

	id := c.Params("id")
	var record model.Blog
	database.DBConnection.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not found")
		context["message"] = "Record not found"
		c.Status(404)
		return c.JSON(context)
	}
	context["record"] = record

	context["statusText"] = "OK"
	context["message"] = "Blog Detail"
	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog Create",
	}

	record := new(model.Blog)
	if err := c.BodyParser(record); err != nil {
		log.Println("Error parsing request: ", err)
		context["statusText"] = "Bad Request"
		context["message"] = "Invalid blog data"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}

	// 获取当前登录用户信息
	user := c.Locals("user")
	if user != nil {
		// 如果用户已登录，设置作者信息
		userModel, ok := user.(model.User)
		if ok {
			record.AuthorID = userModel.ID
			record.Author = userModel.Username
		}
	} else {
		// 如果用户未登录，设置默认作者
		record.Author = "Anonymous"
	}

	// File upload
	file, err := c.FormFile("file")
	if err != nil {
		// 如果没有文件，这不是错误，只是记录一下
		log.Println("No file uploaded or error parsing file: ", err)
	} else if file != nil && file.Size > 0 {
		// 确保 uploads 目录存在
		filename := "./static/uploads/" + file.Filename
		if err := c.SaveFile(file, filename); err != nil {
			log.Println("Error saving file: ", err)
			context["statusText"] = "Error"
			context["message"] = "Failed to save file: " + err.Error()
		} else {
			log.Println("File saved successfully: ", filename)
			record.Image = "/static/uploads/" + file.Filename
		}
	}

	res := database.DBConnection.Create(&record)
	if res.Error != nil {
		log.Println("Error saving record: ", res.Error)
		context["statusText"] = "Internal Server Error"
		context["message"] = "Failed to save blog"
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(context)
	}

	context["message"] = "Blog Created Successfully"
	context["data"] = record
	c.Status(fiber.StatusCreated)
	return c.JSON(context)
}

// Update a Blog
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog Update",
	}

	id := c.Params("id")

	var record model.Blog
	database.DBConnection.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not found")
		context["statusText"] = ""
		context["message"] = "Record not found"
		c.Status(400)
		return c.JSON(context)
	}

	// 获取当前登录用户
	user, ok := c.Locals("user").(model.User)
	if !ok {
		log.Println("User not found in context")
		context["statusText"] = "Error"
		context["message"] = "Authentication error"
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(context)
	}

	// 检查用户是否是博客作者
	if record.AuthorID != 0 && record.AuthorID != user.ID {
		log.Printf("Permission denied: User %d attempted to edit blog %d owned by user %d",
			user.ID, record.ID, record.AuthorID)
		context["statusText"] = "Error"
		context["message"] = "You do not have permission to edit this blog"
		c.Status(fiber.StatusForbidden)
		return c.JSON(context)
	}

	// 解析表单数据
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error parsing request: ", err)
	}

	// 检查是否要删除图片
	removeImage := c.FormValue("removeImage")
	if removeImage == "true" {
		log.Println("Removing image for blog ID:", id)
		record.Image = ""
	} else {
		// 处理文件上传
		file, err := c.FormFile("file")
		if err == nil && file != nil && file.Size > 0 {
			filename := "./static/uploads/" + file.Filename
			if err := c.SaveFile(file, filename); err != nil {
				log.Println("Error saving file: ", err)
				context["statusText"] = "Error"
				context["message"] = "Failed to save file: " + err.Error()
			} else {
				log.Println("File saved successfully: ", filename)
				record.Image = "/static/uploads/" + file.Filename
			}
		}
	}

	res := database.DBConnection.Save(&record)
	if res.Error != nil {
		log.Println("Error saving record: ", res.Error)
		context["statusText"] = "Error"
		context["message"] = "Failed to save record: " + res.Error.Error()
		c.Status(500)
		return c.JSON(context)
	}

	context["message"] = "Blog Updated Successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)
}

// Delete a Blog
func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog Delete",
	}

	id := c.Params("id")
	var record model.Blog
	database.DBConnection.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not found")
		context["statusText"] = ""
		context["message"] = "Record not found"
		c.Status(400)
		return c.JSON(context)
	}

	// 获取当前登录用户
	user, ok := c.Locals("user").(model.User)
	if !ok {
		log.Println("User not found in context")
		context["statusText"] = "Error"
		context["message"] = "Authentication error"
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(context)
	}

	// 检查用户是否是博客作者
	if record.AuthorID != 0 && record.AuthorID != user.ID {
		log.Printf("Permission denied: User %d attempted to delete blog %d owned by user %d",
			user.ID, record.ID, record.AuthorID)
		context["statusText"] = "Error"
		context["message"] = "You do not have permission to delete this blog"
		c.Status(fiber.StatusForbidden)
		return c.JSON(context)
	}
	res := database.DBConnection.Delete(&record)
	if res.Error != nil {
		log.Println("Error deleting record: ", res.Error)
	}

	context["message"] = "Blog Deleted Successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)
}

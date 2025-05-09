package controllers

import (
	"ChronoTrack/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	task.UserID = user.ID

	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task",
		})
	}

	return c.JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	var tasks []models.Task
	database.DB.Where("user_id = ?", user.ID).Find(&tasks)

	return c.JSON(tasks)
}

func UpdateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	id := c.Params("id")

	var task models.Task
	result := database.DB.First(&task, "id = ? AND user_id = ?", id, user.ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to find task"})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	database.DB.Save(&task)
	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	id := c.Params("id")

	result := database.DB.Delete(&models.Task{}, "id = ? AND user_id = ?", id, user.ID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete task"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

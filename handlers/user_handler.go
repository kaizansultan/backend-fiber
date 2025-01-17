package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/kaizansultan/backend-fiber/database"
	"github.com/kaizansultan/backend-fiber/models"
)

// mendapatkan semua User
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

// mendapatkan satu ID user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// membuat User baru
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Warn(err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	database.DB.Create(&user)
	return c.Status(201).JSON(user)
}

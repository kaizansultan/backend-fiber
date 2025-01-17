package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaizansultan/backend-fiber/handlers"
)

func RegisterUserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Get("/", handlers.GetAllUsers)
	userGroup.Get("/:id", handlers.GetUser)
	userGroup.Post("/", handlers.CreateUser)
}

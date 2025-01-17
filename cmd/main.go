package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kaizansultan/backend-fiber/database"
	"github.com/kaizansultan/backend-fiber/routes"
)

func main() {
	log.Println("Server is setting up")
	database.ConnectDB()

	app := fiber.New()

	routes.RegisterUserRoutes(app)

	log.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

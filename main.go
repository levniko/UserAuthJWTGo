package main

import (
	//"./database"
	"test3/database"
	"test3/routes"

	"github.com/gofiber/fiber/v2"
	//"github.com/rs/cors"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}

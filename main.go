package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

}

func main() {
	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

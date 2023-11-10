package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/root27/URL-Shortener/routes"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/shorten", routes.Shorten)

	app.Get("/:id", routes.Redirect)

}

func main() {
	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

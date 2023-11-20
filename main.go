package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/root27/URL-Shortener/routes"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/click-counter", func(c *fiber.Ctx) error {
		return c.Render("click-counter", fiber.Map{})
	})

	app.Post("/shorten", routes.Shorten)

	app.Get("/:id", routes.Redirect)

	app.Get("/metrics/:id", routes.GetMetrics)

}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

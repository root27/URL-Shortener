package routes

import "github.com/gofiber/fiber/v2"

func GetMetrics(c *fiber.Ctx) error {

	id := c.Params("id")

	// Get click counts from redis

	return c.Render("metrics", fiber.Map{})
}

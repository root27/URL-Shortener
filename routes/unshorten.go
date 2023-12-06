package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/root27/URL-Shortener/redis"
)

type UnshortenRequest struct {
	Url string `json:"url"`
}

func Unshorten(c *fiber.Ctx) error {

	var body UnshortenRequest

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if body.Url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL cannot be empty",
		})
	}

	id := strings.Split(body.Url, "/")[3]

	client := redis.NewClient()

	url, err := client.Get(ctx, id).Result()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No valid url found",
		})

	}

	return c.JSON(fiber.Map{
		"originalUrl": url,
	})

}

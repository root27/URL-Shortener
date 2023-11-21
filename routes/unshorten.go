package routes

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
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

	id := strings.Split(body.Url, "/")

	fmt.Println(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Unshortened",
	})

}

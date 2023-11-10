package routes

import "github.com/gofiber/fiber/v2"

func Redirect(c *fiber.Ctx) error {

	id := c.Params("id")

	//Get the URL from the database

	//Redirect the user to it

	if urls[id] == "" {
		return c.SendString("Invalid URL")
	}

	originalURL, found := urls[id]

	if !found {
		return c.SendString("Invalid URL")

	}

	return c.Redirect(originalURL)

}

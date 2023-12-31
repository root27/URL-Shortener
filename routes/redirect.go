package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/root27/URL-Shortener/redis"
)

func Redirect(c *fiber.Ctx) error {

	client := redis.NewClient()

	id := c.Params("id")

	//Get the URL from the database

	//Redirect the user to it

	originalURL, err := client.Get(ctx, id).Result()

	if err != nil {
		log.Println(err)
		return c.Render("error", fiber.Map{})
	}

	return c.Redirect(originalURL)

}

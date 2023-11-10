package routes

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var urls = make(map[string]string)

//Shorten the URL

func Shorten(c *fiber.Ctx) error {

	if c.Query("url") == "" {
		return c.JSON(fiber.Map{
			"error": "No URL provided",
		})
	}

	url := c.Query("url")

	random_string := GenerateRandomString(3)

	shortURL := c.BaseURL() + "/" + random_string

	//Store the URL in the database

	urls[random_string] = url

	//Insert the URL in the database

	//Return the JSON response

	return c.JSON(fiber.Map{
		"success": true,
		"url":     shortURL,
	})
}

//Generate random string

func GenerateRandomString(n int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	random := make([]rune, n)

	for i := range random {
		random[i] = letters[rand.Intn(len(letters))]
	}

	return string(random)

}

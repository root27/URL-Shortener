package routes

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var urls = make(map[string]string)

type Request struct {
	Url string `json:"url"`
}

//Shorten the URL

func Shorten(c *fiber.Ctx) error {

	request := new(Request)

	err := c.BodyParser(request)

	if err != nil {
		log.Println(err)
		return c.SendString("Invalid URL")
	}

	random_string := GenerateRandomString(3)

	shortURL := c.BaseURL() + "/" + random_string

	//Store the URL in the database

	urls[random_string] = request.Url

	//Insert the URL in the database

	//Return the JSON response

	return c.JSON(fiber.Map{

		"url": shortURL,
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

package routes

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/root27/URL-Shortener/redis"
)

var urls = make(map[string]string)

type Request struct {
	Url string `json:"url"`
}

var ctx = context.Background()

//Shorten the URL

func Shorten(c *fiber.Ctx) error {

	client := redis.NewClient()

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

	err = client.Set(ctx, random_string, request.Url, time.Second*10).Err()

	if err != nil {
		log.Println(err)

	}

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

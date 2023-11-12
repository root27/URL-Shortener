package routes

import (
	"context"
	"log"
	"math/rand"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/root27/URL-Shortener/redis"
)

type Request struct {
	Url     string `json:"url"`
	Minutes int    `json:"minutes"`
}

var ctx = context.Background()

//Shorten the URL

func Shorten(c *fiber.Ctx) error {

	var redisExpiration time.Duration

	client := redis.NewClient()

	request := new(Request)

	err := c.BodyParser(request)

	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  "Cannot parse JSON",
		})
	}

	if request.Url == "" {
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  "URL cannot be empty",
		})
	}

	if !Checkurl(request.Url) {
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid URL",
		})
	}

	random_string := GenerateRandomString(3)

	shortURL := c.BaseURL() + "/" + random_string

	//Store the URL in the database

	if request.Minutes == 0 {
		redisExpiration = 0
	} else {
		redisExpiration = time.Duration(request.Minutes) * time.Minute
	}

	err = client.Set(ctx, random_string, request.Url, redisExpiration).Err()

	if err != nil {
		log.Println(err)

	}

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

func Checkurl(url string) bool {

	regexp := regexp.MustCompile(`^http(s)?://[a-z0-9]+(.[a-z0-9]+)*(:[0-9]+)?(/.*)?$`)

	return regexp.MatchString(url)
}

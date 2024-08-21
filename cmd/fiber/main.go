package main

import (
	"framework/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {

	router := fiber.New()

	router.Get("/hello", func(c *fiber.Ctx) error {

		name := c.Query("name")
		auth := c.Get("Authorization", "")

		msg := ""

		if auth == "" {
			msg = "no auth"
		} else {
			msg = msg + auth
		}


		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": name,
			"auth" : msg,
		})
	})

	type Request struct {
		User string `json:"user" xml:"user" form:"user"`
	}

	router.Post("/", func(c *fiber.Ctx) error {

	

		var req = Request{}

		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"payload": req,
		})
	})

	router.Listen(config.APP_PORT)
}

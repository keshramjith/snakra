package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type response struct {
	Msg string `json:"msg"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		resp := response{Msg: ""}
		return c.JSON(resp)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		blob := c.Body()
		fmt.Println(blob)
		return c.JSON(response{Msg: "recording received!"})
	})

	app.Listen(":3001")
}

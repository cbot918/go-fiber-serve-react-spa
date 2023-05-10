package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Static("/", "./ui/dist")

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	})

	err := app.Listen(":3001")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{

	})
	app.Use(logger.New())

	app.Get("/",index)

	app.ListenTLS(":443", "./cert.pem", "./key.pem")
}

func index(c *fiber.Ctx) error{
	return c.SendString("Hello FIBER!")
}


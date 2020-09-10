package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/:value", func(c *fiber.Ctx) {
		c.Send("Get request with value: " + c.Params("value"))
	})
	app.Listen(8080)
}

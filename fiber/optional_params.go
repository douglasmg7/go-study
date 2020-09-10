package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/:name?", func(c *fiber.Ctx) {
		if c.Params("name") != "" {
			c.Send("Hello " + c.Params("name"))
		} else {
			c.Send("No name")
		}
	})
	app.Listen(8080)
}

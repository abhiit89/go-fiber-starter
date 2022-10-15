package main

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/fiber/v2/middleware/monitor"
func main() {
	app := fiber.New()
        app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

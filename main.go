package main

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/fiber/v2/middleware/monitor"
import "github.com/gofiber/fiber/v2/middleware/logger"
import "github.com/gofiber/fiber/v2/middleware/requestid"

func main() {
	app := fiber.New()
	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

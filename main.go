package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Fiber HTML</title>
			</head>
			<body>
				<h1>Hello from Fiber!</h1>
				<p>This is raw HTML rendered directly from Go</p>
			</body>
			</html>
		`
		c.Type("html") // Sets Content-Type: text/html
		return c.SendString(html)
	})

	app.Listen(":8080")
}

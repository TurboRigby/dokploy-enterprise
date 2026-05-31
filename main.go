package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Fuck you, Dokploy! Maybe learn how to make enterprise licenses?")
	})

	app.Post("/licenses/validate", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"valid": true,
		})
	})

	app.Post("/licenses/activate", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{})
	})

	app.Post("/licenses/deactivate", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{})
	})

	app.Listen(":3000")
}

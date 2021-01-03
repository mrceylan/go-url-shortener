package routes

import "github.com/gofiber/fiber/v2"

func SetupHealthRoutes(app *fiber.App) {
	api := app.Group("/health")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Ok")
	})
}

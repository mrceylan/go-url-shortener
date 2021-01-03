package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrceylan/go-url-shortener/pkg/logging"
)

func SetupLogRoutes(app *fiber.App, srv logging.LogService) {
	api := app.Group("/log")
	api.Get("/list", getLogs(srv))
}

func getLogs(srv logging.LogService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		result, err := srv.GetLogs()
		if err != nil {
			return err
		}
		return c.JSON(result)
	}
}

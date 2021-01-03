package routes

import (
	"fmt"
	"time"

	"github.com/mrceylan/go-url-shortener/pkg/logging"
	"github.com/mrceylan/go-url-shortener/pkg/url"

	"github.com/gofiber/fiber/v2"
)

type UrlInput struct {
	Url string
}

func SetupURLRoutes(app *fiber.App, srv url.URLService, logging logging.LogService) {
	api := app.Group("/url")
	api.Post("/", createURL(srv))
	api.Get("/:hash", urlRedirectLog(logging), getURL(srv))
}

func createURL(srv url.URLService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		p := UrlInput{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		result, err := srv.CreateNewURL(p.Url)
		if err != nil {
			return err
		}

		return c.Status(200).JSON(&fiber.Map{
			"hash": fmt.Sprintf("%s://%s/url/%s", c.Protocol(), c.Hostname(), result),
		})
	}
}

func getURL(srv url.URLService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		hash := c.Params("hash")
		result, err := srv.GetURL(hash)
		if err != nil {
			return err
		}
		return c.Redirect(result, 301)
	}
}

func urlRedirectLog(srv logging.LogService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		_log := logging.RedirectLog{
			Hash:        c.Params("hash"),
			Requestid:   fmt.Sprintf("%v", c.Locals("requestid")),
			Requestdate: time.Now().UTC(),
		}
		go srv.CreateRedirectLog(_log)
		return c.Next()
	}
}

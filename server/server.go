package server

import (
	"context"
	"fmt"
	"log"

	"github.com/mrceylan/go-url-shortener/pkg/http/routes"
	"github.com/mrceylan/go-url-shortener/pkg/logging"
	"github.com/mrceylan/go-url-shortener/pkg/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/mrceylan/go-url-shortener/pkg/configuration"
	"github.com/mrceylan/go-url-shortener/pkg/storage/mongo"
	"github.com/mrceylan/go-url-shortener/pkg/storage/redis"
)

var (
	redisConnection *redis.Connection
	mongoConnection *mongo.Connection
)

func InitServer() {
	configuration.LoadConfigurations()
	redisConnection = redis.RedisInit()
	mongoConnection = mongo.MongoInit()
	defer redisConnection.Client.Close()
	defer mongoConnection.Client.Disconnect(context.Background())

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	initServices(app)

	fmt.Println(fmt.Sprintf("Server started to listening on port : %v", configuration.Config.Port))
	app.Listen(fmt.Sprintf(":%v", configuration.Config.Port))

}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	log.Println(err)
	err = ctx.Status(code).SendString("An error occured")
	if err != nil {
		return ctx.Status(500).SendString("Internal Server Error")
	}
	return nil
}

func initServices(app *fiber.App) {
	loggingService := logging.InitService(mongoConnection)
	routes.SetupLogRoutes(app, loggingService)

	urlService := url.InitService(redisConnection)
	routes.SetupURLRoutes(app, urlService, loggingService)

	routes.SetupHealthRoutes(app)
}

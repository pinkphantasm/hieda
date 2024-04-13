package server

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/shelepuginivan/microservices-template/src/static_service/internal/pkg/health"
)

func registerApi(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(health.NewResponse("All systems operational"))
	})
}

func registerStatic(app *fiber.App) {
	app.Static("/", "./public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})
}

func registerSwagger(app *fiber.App) {
	app.Use(swagger.New(swagger.Config{
		Title:    "Static Service API",
		FilePath: "./api/swagger.json",
	}))

	app.Use(swagger.New(swagger.Config{
		Title:    "Static Service API (API Gateway)",
		FilePath: "./api/swagger.gateway.json",
	}))
}

func New() *fiber.App {
	app := fiber.New()

	registerApi(app)
	registerStatic(app)
	registerSwagger(app)

	return app
}

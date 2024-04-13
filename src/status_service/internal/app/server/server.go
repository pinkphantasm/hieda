// Package server provides methods to initialize app server.
package server

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/pinkphantasm/hieda/src/status_service/internal/app/controllers"
	"github.com/pinkphantasm/hieda/src/status_service/internal/pkg/health"
)

func registerHandlers(app *fiber.App) {
	healthAdapter := health.NewAdapter()
	healthController := controllers.NewHealth(healthAdapter)

	app.Get("/", healthController.StatusPage)

	api := app.Group("/api")
	api.Get("/health", healthController.SelfHealth)
	api.Get("/status", healthController.CheckStatuses)
}

func registerSwagger(app *fiber.App) {
	app.Use(swagger.New(swagger.Config{
		Title:    "Status Service API",
		FilePath: "./api/swagger.json",
	}))

	app.Use(swagger.New(swagger.Config{
		Title:    "Status Service API (API Gateway)",
		FilePath: "./api/swagger.gateway.json",
	}))
}

func New() *fiber.App {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	registerHandlers(app)
	registerSwagger(app)

	return app
}

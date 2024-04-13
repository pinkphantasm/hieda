package server

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/pinkphantasm/hieda/src/index_service/internal/app/controllers"
	"github.com/pinkphantasm/hieda/src/index_service/internal/pkg/health"
)

func registerApi(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(health.NewResponse("All systems operational"))
	})
}

func registerSwagger(app *fiber.App) {
	app.Use(swagger.New(swagger.Config{
		Title:    "Index Service API",
		FilePath: "./api/swagger.json",
	}))

	app.Use(swagger.New(swagger.Config{
		Title:    "Index Service API (API Gateway)",
		FilePath: "./api/swagger.gateway.json",
	}))
}

func registerViews(app *fiber.App, co *controllers.ViewsController) {
	app.Get("/", co.IndexPage)
}

func New() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	viewsController := controllers.NewViews()

	registerApi(app)
	registerSwagger(app)
	registerViews(app, viewsController)

	return app
}

package server

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/pinkphantasm/hieda/src/cert_service/internal/app/controllers"
	"github.com/pinkphantasm/hieda/src/cert_service/internal/pkg/crypto"
	"github.com/pinkphantasm/hieda/src/cert_service/internal/pkg/hash"
	"github.com/pinkphantasm/hieda/src/cert_service/internal/pkg/health"
)

func registerApi(app *fiber.App, co *controllers.ApiController) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(health.NewResponse("All systems operational"))
	})

	api.Post("/sign", co.Sign)
	api.Post("/verify", co.Verify)
}

func registerSwagger(app *fiber.App) {
	app.Use(swagger.New(swagger.Config{
		Title:    "Sign Service API",
		FilePath: "./api/swagger.json",
	}))

	app.Use(swagger.New(swagger.Config{
		Title:    "Sign Service API (API Gateway)",
		FilePath: "./api/swagger.gateway.json",
	}))
}

func registerViews(app *fiber.App, co *controllers.ViewsController) {
	app.Get("/sign", co.SignPage)
	app.Get("/verify", co.VerifyPage)
}

func New() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	ca := crypto.NewAdapter()
	ha := hash.NewAdapter()
	apiController := controllers.NewApi(ca, ha)
	viewsController := controllers.NewViews()

	registerApi(app, apiController)
	registerSwagger(app)
	registerViews(app, viewsController)

	return app
}

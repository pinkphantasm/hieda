package server

import (
	"io"
	"net/http"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/pinkphantasm/hieda/src/sign_service/internal/pkg/crypto"
	"github.com/pinkphantasm/hieda/src/sign_service/internal/pkg/hash"
	"github.com/pinkphantasm/hieda/src/sign_service/internal/pkg/health"
)

func registerHandlers(
	app *fiber.App,
	ca *crypto.Adapter,
	ha *hash.Adapter,
) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(health.NewResponse("All systems operational"))
	})

	api.Post("/sign", func(c *fiber.Ctx) error {
		uploadedFile, err := c.FormFile("file")
		if err != nil {
			return err
		}

		file, err := uploadedFile.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		sha256sum := ha.Hash(data)
		signature, err := ca.Sign(sha256sum)

		if err != nil {
			return err
		}

		c.Response().Header.Set(fiber.HeaderContentType, "application/octet-stream")
		c.Response().Header.Set(fiber.HeaderContentDisposition, "attachment; filename=signature.hieda")
		return c.Send(signature)
	})

	api.Post("/verify", func(c *fiber.Ctx) error {
		uploadedFile, err := c.FormFile("file")
		if err != nil {
			return err
		}

		file, err := uploadedFile.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		signatureUploadedFile, err := c.FormFile("signature")
		if err != nil {
			return err
		}

		signatureFile, err := signatureUploadedFile.Open()
		if err != nil {
			return err
		}
		defer signatureFile.Close()

		signature, err := io.ReadAll(signatureFile)
		if err != nil {
			return err
		}

		if !ca.Compare(ha.Hash(data), signature) {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"details": "err invalid signature",
			})
		}

		return c.JSON(fiber.Map{
			"details": "ok",
		})
	})
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

func New() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	ca := crypto.NewAdapter()
	ha := hash.NewAdapter()

	registerHandlers(app, ca, ha)
	registerSwagger(app)

	return app
}

package controllers

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pinkphantasm/hieda/src/cert_service/internal/pkg/crypto"
	"github.com/pinkphantasm/hieda/src/cert_service/internal/pkg/hash"
)

type ApiController struct {
	ca *crypto.Adapter
	ha *hash.Adapter
}

func NewApi(ca *crypto.Adapter, ha *hash.Adapter) *ApiController {
	return &ApiController{ca: ca, ha: ha}
}

func (co ApiController) Sign(c *fiber.Ctx) error {
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

	sha256sum := co.ha.Hash(data)
	signature, err := co.ca.Sign(sha256sum)

	if err != nil {
		return err
	}

	c.Response().Header.Set(fiber.HeaderContentType, "application/octet-stream")
	c.Response().Header.Set(fiber.HeaderContentDisposition, "attachment; filename=signature.hieda")
	return c.Send(signature)
}

func (co ApiController) Verify(c *fiber.Ctx) error {
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

	if !co.ca.Compare(co.ha.Hash(data), signature) {
		return co.VerificationFailure(c, c.FormValue("format", "json"))
	}

	return co.VerificationSuccess(c, c.FormValue("format", "json"))
}

func (co ApiController) VerificationSuccess(c *fiber.Ctx, format string) error {
	switch format {
	case "html":
		c.Response().Header.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.SendString(`<div>
	<h3>Verification Succeeded!</h3>
	<p>Powered by Hieda</p>
</div>`)

	default:
		return c.JSON(fiber.Map{
			"details": "ok",
		})
	}
}

func (co ApiController) VerificationFailure(c *fiber.Ctx, format string) error {
	c.Status(http.StatusUnauthorized)

	switch format {
	case "html":
		c.Response().Header.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.SendString(`<div">
	<h3>Verification Failed!</h3>
	<p>Powered by Hieda</p>
</div>`)

	default:
		return c.JSON(fiber.Map{
			"details": "err invalid signature",
		})
	}
}

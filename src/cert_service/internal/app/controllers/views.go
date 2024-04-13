package controllers

import "github.com/gofiber/fiber/v2"

type ViewsController struct{}

func NewViews() *ViewsController {
	return &ViewsController{}
}

func (co ViewsController) SignPage(c *fiber.Ctx) error {
	return c.Render("sign", fiber.Map{})
}

func (co ViewsController) VerifyPage(c *fiber.Ctx) error {
	return c.Render("verify", fiber.Map{})
}

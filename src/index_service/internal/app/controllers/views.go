package controllers

import "github.com/gofiber/fiber/v2"

type ViewsController struct{}

func NewViews() *ViewsController {
	return &ViewsController{}
}

func (co ViewsController) IndexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

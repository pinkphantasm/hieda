// Package controllers provides HTTP controllers for the service API.
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pinkphantasm/hieda/src/status_service/internal/pkg/config"
	"github.com/pinkphantasm/hieda/src/status_service/internal/pkg/health"
	"github.com/pinkphantasm/hieda/src/status_service/internal/pkg/service"
)

// HealthController is a Fiber service health controller.
type HealthController struct {
	a *health.Adapter    // Underlying health adapter.
	s []*service.Service // Services whose status is requested.
}

// NewHealth returns a new instance of HealthController.
func NewHealth(a *health.Adapter) *HealthController {
	return &HealthController{
		a: a,
		s: config.New().Services,
	}
}

// SelfHealth responds with status of this service.
func (co HealthController) SelfHealth(c *fiber.Ctx) error {
	return c.JSON(health.NewResponse(health.DetailsOperational))
}

// CheckStatuses responds with statuses of the services.
func (co HealthController) CheckStatuses(c *fiber.Ctx) error {
	return c.JSON(co.a.CheckHealth(co.s))
}

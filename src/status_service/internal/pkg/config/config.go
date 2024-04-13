// Package config provides env configuration for service.
package config

import (
	"os"
	"strings"

	"github.com/shelepuginivan/microservices-template/src/status_service/internal/pkg/service"
)

type Config struct {
	Addr     string
	Name     string
	Services []*service.Service
}

func New() *Config {
	var services []*service.Service

	for _, entry := range strings.Fields(os.Getenv(ENV_SERVICES)) {
		entryItems := strings.Split(entry, "=")

		if len(entryItems) != 2 {
			continue
		}

		services = append(services, &service.Service{
			Addr: entryItems[1],
			Name: entryItems[0],
		})
	}

	return &Config{
		Addr:     os.Getenv(ENV_ADDR),
		Name:     ServiceName,
		Services: services,
	}
}

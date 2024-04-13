// Package config provides env configuration for service.
package config

import (
	"os"
)

type Config struct {
	Addr string
	Name string
}

func New() *Config {
	return &Config{
		Addr: os.Getenv(ENV_ADDR),
		Name: ServiceName,
	}
}

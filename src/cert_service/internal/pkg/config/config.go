// Package config provides env configuration for service.
package config

import (
	"os"
)

type Config struct {
	Addr       string
	PrivateKey string
	Name       string
}

func New() *Config {
	return &Config{
		Addr:       os.Getenv(ENV_ADDR),
		PrivateKey: os.Getenv(ENV_PRIVATE_KEY),
		Name:       ServiceName,
	}
}

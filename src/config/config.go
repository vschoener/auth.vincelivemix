package config

import (
	"github.com/caarlos0/env/v6"
)

// Config holds the app configuration
type Config struct {
	WebConfig      WebConfig
	DatabaseConfig DatabaseConfig
}

// ProvideConfig build the configuration application
func ProvideConfig() (Config, error) {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		return config, err
	}

	return config, nil
}

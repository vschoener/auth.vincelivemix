package config

import "github.com/caarlos0/env/v6"

type Config struct {
	WebConfig WebConfig
}

func New() (Config, error) {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		return config, err
	}

	return config, nil
}

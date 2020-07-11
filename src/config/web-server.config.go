package config

// WebConfig structure
type WebConfig struct {
	Port string `env:"PORT" envDefault:"8000"`
}

func ProvideWebConfig(config Config) WebConfig {
	return config.WebConfig
}

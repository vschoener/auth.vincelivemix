package config

// WebConfig structure
type WebConfig struct {
	Port string `env:"PORT" envDefault:"8000"`
}

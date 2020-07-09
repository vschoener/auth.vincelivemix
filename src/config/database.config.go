package config

// DatabaseConfig structure
type DatabaseConfig struct {
	URL string `env:"DATABASE_URL" envDefault:"postgresql://user:secret@localhost:5432/authentication"`
}

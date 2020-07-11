package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vschoener/auth.vincelivemix/src/config"
)

// Database struct hold the connection
type Database struct {
	config     config.DatabaseConfig
	Connection pgxpool.Pool
}

// NewDatabase creates a new Database service
func ProvideDatabase(databaseConfig config.DatabaseConfig) Database {
	return Database{
		config: databaseConfig,
	}
}

// Connect build the database service
func (d *Database) Connect() error {
	fmt.Print("Connecting to database...\n")
	dbPool, err := pgxpool.Connect(context.Background(), d.config.URL)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %w\n", err)
	}

	fmt.Print("Connected to database\n")

	d.Connection = *dbPool

	return nil
}

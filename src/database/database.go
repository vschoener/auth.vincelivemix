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
	Connection *pgxpool.Pool
}

// ProvideDatabaseConnection creates a new Database connection
func ProvideDatabaseConnection(databaseConfig config.DatabaseConfig) *Database {
	database := &Database{
		config: databaseConfig,
	}

	database.connect()

	return database
}

// Connect build the database service
func (d *Database) connect() error {
	fmt.Print("Connecting to database...\n")
	dbPool, err := pgxpool.Connect(context.Background(), d.config.URL)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %w\n", err)
	}

	fmt.Print("Connected to database\n")

	d.Connection = dbPool

	return nil
}

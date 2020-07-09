package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vschoener/auth.vincelivemix/src/config"
)

// Database struct hold the connection
type Database struct {
	Connection pgxpool.Pool
}

// Connect build the database service
func Connect(databaseConfig config.DatabaseConfig) (Database, error) {
	database := Database{}

	fmt.Print("Connecting to database...\n")
	dbPool, err := pgxpool.Connect(context.Background(), databaseConfig.URL)
	if err != nil {
		return database, fmt.Errorf("Unable to connect to database: %w\n", err)
	}

	fmt.Print("Connected to database\n")

	database.Connection = *dbPool

	return database, nil
}

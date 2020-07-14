package app

import (
	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/database"
	router "github.com/vschoener/auth.vincelivemix/src/http"
)

// AppModule structure
type AppModule struct {
	config   config.Config
	router   router.Router
	database *database.Database
}

// ProvideAppModule an app module
func ProvideAppModule(config config.Config, router router.Router, database *database.Database) AppModule {
	return AppModule{
		config:   config,
		router:   router,
		database: database,
	}
}

// Bootstrap the app module
func (a AppModule) Bootstrap() error {
	// Defer the database close connected during the Provider
	defer a.database.Connection.Close()

	if err := a.router.Start(); err != nil {
		return err
	}

	return nil
}

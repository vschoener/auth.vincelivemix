package main

import (
	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/database"
	router "github.com/vschoener/auth.vincelivemix/src/http"
)

type AppModule struct {
	config   config.Config
	router   router.Router
	database database.Database
}

func ProvideAppModule(config config.Config, router router.Router, database database.Database) AppModule {
	return AppModule{
		config:   config,
		router:   router,
		database: database,
	}
}

func (a AppModule) Bootstrap() error {
	if err := a.database.Connect(); err != nil {
		return err
	}

	if err := a.router.Start(); err != nil {
		return err
	}

	return nil
}

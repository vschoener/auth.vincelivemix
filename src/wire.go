//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/controllers"
	"github.com/vschoener/auth.vincelivemix/src/database"
	router "github.com/vschoener/auth.vincelivemix/src/http"
	"github.com/vschoener/auth.vincelivemix/src/services"
)

func appModuleInjector() (AppModule, error) {
	panic(wire.Build(
		ProvideAppModule,
		config.ProvideConfig,
		config.ProvideWebConfig,
		config.ProvideDatabaseConfig,
		database.ProvideDatabase,
		services.ProvideAuthService,
		controllers.ProvideAuthController,
		controllers.ProvideControllers,
		router.ProvideRouter,
	))
}

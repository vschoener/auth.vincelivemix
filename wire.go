//+build wireinject
package main

import (
	"github.com/google/wire"
	app "github.com/vschoener/auth.vincelivemix/src"
	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/controllers"
	"github.com/vschoener/auth.vincelivemix/src/database"
	router "github.com/vschoener/auth.vincelivemix/src/http"
	"github.com/vschoener/auth.vincelivemix/src/repositories"
	"github.com/vschoener/auth.vincelivemix/src/services"
)

func appModuleInjector() (app.AppModule, error) {
	panic(wire.Build(
		app.ProvideAppModule,
		config.ProvideConfig,
		config.ProvideWebConfig,
		config.ProvideDatabaseConfig,
		config.ProvideSecurityConfig,
		database.ProvideDatabaseConnection,
		repositories.ProvideUserRepository,
		services.ProvideUserService,
		services.ProvideAuthService,
		controllers.ProvideAuthController,
		controllers.ProvideControllers,
		router.ProvideRouter,
	))
}

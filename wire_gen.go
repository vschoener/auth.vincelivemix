// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/vschoener/auth.vincelivemix/src"
	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/controllers"
	"github.com/vschoener/auth.vincelivemix/src/database"
	"github.com/vschoener/auth.vincelivemix/src/http"
	"github.com/vschoener/auth.vincelivemix/src/repositories"
	"github.com/vschoener/auth.vincelivemix/src/services"
)

import (
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func appModuleInjector() (app.AppModule, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return app.AppModule{}, err
	}
	webConfig := config.ProvideWebConfig(configConfig)
	databaseConfig := config.ProvideDatabaseConfig(configConfig)
	databaseDatabase := database.ProvideDatabaseConnection(databaseConfig)
	userRepository := repositories.ProvideUserRepository(databaseDatabase)
	userService := services.ProvideUserService(userRepository)
	authService := services.ProvideAuthService(userService)
	authController := controllers.ProvideAuthController(authService)
	controllersControllers := controllers.ProvideControllers(authController)
	routerRouter := router.ProvideRouter(webConfig, controllersControllers)
	appModule := app.ProvideAppModule(configConfig, routerRouter, databaseDatabase)
	return appModule, nil
}
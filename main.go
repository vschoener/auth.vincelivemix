package main

import (
	"fmt"

	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/database"
	"github.com/vschoener/auth.vincelivemix/src/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config, err := config.New()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	database, err := database.Connect(config.DatabaseConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer database.Connection.Close()

	router := router.New()

	router.Run("localhost:" + config.WebConfig.Port)
}

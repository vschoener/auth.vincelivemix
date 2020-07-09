package main

import (
	"fmt"

	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/controllers"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config, err := config.New()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", controllers.Login)

	r.Run("localhost:" + config.WebConfig.Port)
}

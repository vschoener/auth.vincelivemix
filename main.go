package main

import (
	"github.com/vschoener/auth.vincelivemix/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", controllers.Login)

	r.Run("localhost:8080")
}

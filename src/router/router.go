package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vschoener/auth.vincelivemix/src/controllers"
)

// New create a new Gin Engine
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", controllers.Login)

	return r
}

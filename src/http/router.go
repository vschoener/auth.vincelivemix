package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/controllers"
)

type Router struct {
	engine      *gin.Engine
	config      config.WebConfig
	controllers controllers.Controllers
}

// New create a new Gin Engine
func ProvideRouter(config config.WebConfig, controllers controllers.Controllers) Router {
	r := gin.Default()

	router := Router{
		engine: r,
		config: config,
	}

	router.setHandler()

	return router
}

func (r *Router) setHandler() {
	r.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.engine.POST("/login", r.controllers.AuthController.Login)
}

func (r *Router) Start() error {
	return r.engine.Run("localhost:" + r.config.Port)
}

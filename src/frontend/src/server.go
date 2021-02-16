package main

import (
	"frontend/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run(":8080")
}

func setupServer() *gin.Engine {
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*")

	frontendEngin := engine.Group("/frontend")
	{
		v1Engin := frontendEngin.Group("/v1")
		{
			v1Engin.GET("/homepage", controllers.HandlerHomePage)
		}
	}

	return engine
}

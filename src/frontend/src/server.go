package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*")

	frontendEngin := engine.Group("/frontend")
	{
		v1Engin := frontendEngin.Group("/v1")
		{
			v1Engin.GET("/", controllers.handlerHomePage)
		}
	}

	engine.Run(":8080")
}

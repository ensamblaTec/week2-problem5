package server

import (
	"github.com/ensamblaTec/learning-path/week2/p5/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("public/templates/**/*")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	pages := controllers.PagesController{}

	v1 := router.Group("/v1")
	{
		v1.GET("/", pages.GetIndex)
	}

	return router
}

package controllers

import "github.com/gin-gonic/gin"

type PagesController struct{}

func (pc PagesController) GetIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":   "Index Page",
		"message": "This is the index Page",
	})
}

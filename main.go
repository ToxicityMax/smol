package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ToxicityMax/smol/controller"
	"github.com/ToxicityMax/smol/models"
)

func main() {
	r := gin.Default()
	models.ConnectDb()

	//Routes
	r.GET("/urls", controller.GetallUrls)
	r.POST("/url", controller.GenShortUrl)
	r.POST("/:slug", controller.Redirect)

	err := r.Run()
	if err != nil {
		print(err)
		return
	}
}

package main

import (
	"github.com/ToxicityMax/smol/config"
	"github.com/ToxicityMax/smol/controller"
	"github.com/ToxicityMax/smol/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gin.SetMode(config.C.DEBUG)

	// Connect to db
	models.ConnectDb()

	// Cors setup
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	r.Use(cors.New(defaultConfig))

	//Routes
	// r.GET("/urls", controller.GetallUrls)
	r.POST("/url", controller.GenShortUrl)
	r.GET("/:slug", controller.Redirect)
	r.POST("/:slug", controller.PasswordVerify)

	err := r.Run()
	if err != nil {
		print(err)
		return
	}
}

package main

import (
	"tools.cyberkrypts.dev/controllers"
	"tools.cyberkrypts.dev/db"
	"tools.cyberkrypts.dev/env"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.Init()

	router.LoadHTMLGlob("templates/**/**")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/youtube", controllers.YoutubeController{}.Index)
	router.GET("/youtube/video", controllers.YoutubeController{}.GetVideoInfo)

	router.GET("/shortener", controllers.ShortenerController{}.Index)
	router.POST("/shortener/generate", controllers.ShortenerController{}.Generate)
	router.GET("/r/:short_url_code", controllers.ShortenerController{}.Redirect)

	port := env.GetEnv().Port
	router.Run(":" + port)
}

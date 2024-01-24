package main

import (
	"os"

	"youtube-downloader-web/controllers"
	"youtube-downloader-web/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.Init()

	router.LoadHTMLGlob("templates/**/**")

	router.GET("/youtube", controllers.YoutubeController{}.Index)
	router.GET("/youtube/video", controllers.YoutubeController{}.GetVideoInfo)

	router.GET("/shortener", controllers.ShortenerController{}.Index)
	router.POST("/shortener/generate", controllers.ShortenerController{}.Generate)
	router.GET("/r/:short_url_code", controllers.ShortenerController{}.Redirect)

	if os.Getenv("PORT") != "" {
		router.Run(":" + os.Getenv("PORT"))
		return
	}

	router.Run(":8080")
}

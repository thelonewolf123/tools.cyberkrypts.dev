package main

import (
	"fmt"
	"os"

	"youtube-downloader-web/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/**")

	router.GET("/youtube", controllers.YoutubeController{}.Index)
	router.GET("/youtube/video", controllers.YoutubeController{}.GetVideoInfo)

	router.GET("/shortener", func(ctx *gin.Context) {
		ctx.HTML(200, "shortener.html", gin.H{})
	})

	router.POST("/shortener/generate", func(ctx *gin.Context) {
		longURL := ctx.PostForm("long_url")
		fmt.Println(longURL)

		ctx.HTML(200, "short-url.html", gin.H{})
	})

	if os.Getenv("PORT") != "" {
		router.Run(":" + os.Getenv("PORT"))
		return
	}

	router.Run(":8080")
}

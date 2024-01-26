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

	router.GET("/", controllers.HomeController{}.Index)

	router.GET("/youtube", controllers.YoutubeController{}.Index)
	router.GET("/youtube/video", controllers.YoutubeController{}.GetVideoInfo)

	router.GET("/shortener", controllers.ShortenerController{}.Index)
	router.POST("/shortener/generate", controllers.ShortenerController{}.Generate)
	router.GET("/r/:short_url_code", controllers.ShortenerController{}.Redirect)

	router.GET("/talk-swipe", controllers.TalkSwipeController{}.Index)
	router.GET("/talk-swipe/new-chat", controllers.TalkSwipeController{}.NewChat)

	port := env.GetEnv().Port
	router.Run(":" + port)
}

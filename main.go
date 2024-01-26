package main

import (
	"tools.cyberkrypts.dev/controllers"
	"tools.cyberkrypts.dev/db"
	"tools.cyberkrypts.dev/env"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.Init()

	router.LoadHTMLGlob("templates/**/**")

	router.GET("/", func(ctx *gin.Context) {
		utils.RenderTemplate(200, ctx, pages.Home())
	})

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

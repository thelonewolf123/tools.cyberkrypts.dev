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

	router.Static("/static", "./static")

	router.GET("/", controllers.HomeController{}.Index)

	router.GET("/youtube", controllers.YoutubeController{}.Index)
	router.GET("/youtube/video", controllers.YoutubeController{}.GetVideoInfo)

	router.GET("/shortener", controllers.ShortenerController{}.Index)
	router.POST("/shortener/generate", controllers.ShortenerController{}.Generate)
	router.GET("/r/:short_url_code", controllers.ShortenerController{}.Redirect)

	router.GET("/talk-swipe", controllers.TalkSwipeController{}.Index)
	router.GET("/talk-swipe/new-chat", controllers.TalkSwipeController{}.NewChat)

	router.GET("/send-files", controllers.SendFilesController{}.Index)
	router.GET("/send-files/ws", controllers.SendFilesController{}.SendFilesWs)
	router.POST("/send-files/meta-data", controllers.SendFilesController{}.MetaData)
	router.GET("/f/:file_id", controllers.SendFilesController{}.DownloadFile)

	router.GET("/json-formatter", controllers.JsonFormatterController{}.Index)

	port := env.GetEnv().Port
	router.Run(":" + port)
}

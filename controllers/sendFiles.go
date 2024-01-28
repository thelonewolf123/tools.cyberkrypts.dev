package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"tools.cyberkrypts.dev/db"
	"tools.cyberkrypts.dev/env"
	"tools.cyberkrypts.dev/templates/components"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"
)

type SendFilesController struct{}

var fsSessions map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)

func (c SendFilesController) Index(ctx *gin.Context) {
	utils.RenderTemplate(200, ctx, pages.SendFilesIndex())
}

func (c SendFilesController) SendFilesWs(ctx *gin.Context) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer conn.Close()

	fsSessions[conn] = true

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}

		for client := range fsSessions {
			if client != conn {
				client.WriteMessage(websocket.TextMessage, message)
			}
		}
	}

}

func (c SendFilesController) MetaData(ctx *gin.Context) {
	fileName := ctx.PostForm("file_name")
	fileSize := ctx.PostForm("file_size")
	fileId := utils.GetRandomCode(5)

	database, err := db.GetDb()
	ApplicationBaseUrl := env.GetEnv().ApplicationBaseUrl
	fileDownloadUrl := ApplicationBaseUrl + "/f/" + fileId

	if err != nil {
		utils.RenderTemplate(200, ctx, components.SendFilesResult("", err.Error()))
		return
	}
	_, err = database.Exec(`INSERT INTO send_files (file_id, file_name, file_size) VALUES ($1, $2, $3)`, fileId, fileName, fileSize)

	if err != nil {
		utils.RenderTemplate(200, ctx, components.SendFilesResult("", err.Error()))
		return
	}

	utils.RenderTemplate(200, ctx, components.SendFilesResult(fileDownloadUrl, ""))
}

func (c SendFilesController) DownloadFile(ctx *gin.Context) {
	file_id := ctx.Param("file_id")
	fmt.Println(file_id)
	database, err := db.GetDb()
	if err != nil {
		utils.RenderTemplate(200, ctx, pages.DownloadFilesIndex("", "", "Sever error! Please try again later"))
		return
	}

	var file_name string
	var file_size string

	err = database.QueryRow(`SELECT file_name, file_size FROM send_files WHERE file_id = $1`, file_id).Scan(&file_name, &file_size)

	if err != nil {
		utils.RenderTemplate(200, ctx, pages.DownloadFilesIndex("", "", "File not found"))
		return
	}

	utils.RenderTemplate(200, ctx, pages.DownloadFilesIndex(file_name, file_size, ""))
}

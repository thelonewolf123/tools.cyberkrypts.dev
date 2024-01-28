package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	// file_name := ctx.PostForm("file_name")
	// file_size := ctx.PostForm("file_size")
	// file_id := ctx.PostForm("file_id")

	utils.RenderTemplate(200, ctx, components.SendFilesResult("new file.mp4", ""))
}

func (c SendFilesController) DownloadFile(ctx *gin.Context) {
	file_id := ctx.Param("file_id")
	fmt.Println(file_id)
	utils.RenderTemplate(200, ctx, pages.DownloadFilesIndex("new file.mp4", "1.4 MB"))
}

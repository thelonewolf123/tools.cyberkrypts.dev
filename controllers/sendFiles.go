package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/inhies/go-bytesize"
	"tools.cyberkrypts.dev/db"
	"tools.cyberkrypts.dev/env"
	"tools.cyberkrypts.dev/templates/components"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"
)

type SendFilesController struct{}

type SendFilesMessage struct {
	Type         string `json:"type"`
	Client       string `json:"client"`
	FileId       string `json:"file_id"`
	WebRtcOffer  string `json:"web_rtc_offer"`
	WebRtcAnswer string `json:"web_rtc_answer"`
}

type PeerConnection struct {
	Sender       *websocket.Conn
	Receiver     *websocket.Conn
	WebRtcOffer  string
	WebRtcAnswer string
}

var sendFilesClients map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)
var peerConnection map[string]*PeerConnection = make(map[string]*PeerConnection)

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

	sendFilesClients[conn] = true

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			delete(sendFilesClients, conn)
		}

		var sendFilesMessage SendFilesMessage
		err = json.Unmarshal(message, &sendFilesMessage)

		if err != nil {
			fmt.Println(err)
			continue

		}

		if peerConnection[sendFilesMessage.FileId] == nil {
			peerConnection[sendFilesMessage.FileId] = &PeerConnection{}
		}

		if sendFilesMessage.Client == "sender" {
			peerConnection[sendFilesMessage.FileId].Sender = conn
		}

		if sendFilesMessage.Client == "receiver" {
			peerConnection[sendFilesMessage.FileId].Receiver = conn
		}

		if sendFilesMessage.Type == "web_rtc_offer" {
			peerConnection[sendFilesMessage.FileId].WebRtcOffer = sendFilesMessage.WebRtcOffer
		}

		if sendFilesMessage.Type == "web_rtc_answer" {
			peerConnection[sendFilesMessage.FileId].WebRtcAnswer = sendFilesMessage.WebRtcAnswer
			peerConnection[sendFilesMessage.FileId].Sender.WriteMessage(websocket.TextMessage, message)
		}

		if sendFilesMessage.Type == "get_web_rtc_offer" {
			jsonMessage, _ := json.Marshal(&SendFilesMessage{
				Type:        "web_rtc_offer",
				WebRtcOffer: peerConnection[sendFilesMessage.FileId].WebRtcOffer,
			})
			conn.WriteMessage(websocket.TextMessage, jsonMessage)
		}

		if sendFilesMessage.Type == "start_download" {
			jsonMessage, _ := json.Marshal(sendFilesMessage)
			peerConnection[sendFilesMessage.FileId].Sender.WriteMessage(websocket.TextMessage, jsonMessage)
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

	bytesize.Format = "%.2f"
	file_size_int, _ := strconv.Atoi(file_size)
	file_size = bytesize.New(float64(file_size_int)).String()

	utils.RenderTemplate(200, ctx, pages.DownloadFilesIndex(file_name, file_size, ""))
}

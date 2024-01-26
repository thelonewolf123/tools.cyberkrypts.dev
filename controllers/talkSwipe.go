package controllers

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"tools.cyberkrypts.dev/templates/components"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"
)

type TalkSwipeController struct{}

type Message struct {
	ChatMessage string `json:"chat_message"`
}

var sessions = make(map[*websocket.Conn]bool)

func (ts TalkSwipeController) Index(ctx *gin.Context) {
	utils.RenderTemplate(200, ctx, pages.TalkSwipeIndex())
}

func (ts TalkSwipeController) NewChat(ctx *gin.Context) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	sessions[conn] = true
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			delete(sessions, conn) // Remove the client if there's an error
			return
		}

		jsonMsg := new(Message)
		err = json.Unmarshal(message, &jsonMsg)

		if err != nil {
			delete(sessions, conn) // Remove the client if there's an error
			return
		}

		outgoingMessage := &bytes.Buffer{}
		incomingMessage := &bytes.Buffer{}

		components.ChatMessage(jsonMsg.ChatMessage, true).Render(context.Background(), outgoingMessage)
		components.ChatMessage(jsonMsg.ChatMessage, false).Render(context.Background(), incomingMessage)

		for client := range sessions {
			var err error
			if client != conn {
				err = client.WriteMessage(websocket.TextMessage, outgoingMessage.Bytes())
			} else {
				err = client.WriteMessage(websocket.TextMessage, incomingMessage.Bytes())
			}

			if err != nil {
				delete(sessions, conn) // Remove the client if there's an error
				return
			}
		}
	}
}

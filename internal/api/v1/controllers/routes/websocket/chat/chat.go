package chat

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/websocket/chat"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Hub         *chat.Hub
	Client      *chat.Client
}

func NewChatRoutes(rg *gin.RouterGroup, hub *chat.Hub, client *chat.Client) {
	api := Routes{
		RouterGroup: rg,
		Hub:         hub,
		Client:      client,
	}
	api.setupChatRoutes()
}

func (r *Routes) setupChatRoutes() {
	r.RouterGroup.GET("/ws", func(ctx *gin.Context) {
		r.Client.ServeWs(r.Hub, ctx)
	})
}

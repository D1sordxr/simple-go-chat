package websocket

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/websocket"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *websocket.Handler
}

func NewWebSocketRoutes(rg *gin.RouterGroup, h *websocket.Handler) {
	api := Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	api.setupWebSocketRoutes()
}

func (r *Routes) setupWebSocketRoutes() {
	r.RouterGroup.GET("/ws", r.Handler.HandleWebSocket)
}

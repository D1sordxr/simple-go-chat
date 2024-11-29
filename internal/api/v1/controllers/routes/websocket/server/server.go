package server

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/websocket/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Server      *server.Server
}

func NewWebSocket(rg *gin.RouterGroup, s *server.Server) {
	route := Routes{
		RouterGroup: rg,
		Server:      s,
	}
	route.setupWSRoute()
}

func (r *Routes) setupWSRoute() {
	r.RouterGroup.GET("/ws", func(ctx *gin.Context) {
		websocket.Handler(r.Server.HandleWebSocket).ServeHTTP(ctx.Writer, ctx.Request)
	})
}

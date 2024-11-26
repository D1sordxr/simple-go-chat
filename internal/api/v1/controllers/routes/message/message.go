package message

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/message"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *message.Handler
}

func NewMessageRoutes(rg *gin.RouterGroup, h *message.Handler) {
	api := Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	api.setupMessageRoutes()
}

func (r *Routes) setupMessageRoutes() {
	api := r.RouterGroup.Group("/messages")
	{
		api.POST("/message", r.Handler.WriteMessage)
	}
}

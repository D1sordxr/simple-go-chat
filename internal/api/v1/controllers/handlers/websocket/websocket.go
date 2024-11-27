package websocket

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/responses"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handler struct {
	Upgrader websocket.Upgrader
}

func NewWebSocketHandler() *Handler {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return &Handler{
		Upgrader: upgrader,
	}
}

func (h *Handler) HandleWebSocket(c *gin.Context) {
	conn, err := h.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.CommonResponse{
			Message: "WebSocket upgrade failed",
			Data:    err.Error(),
		})
		return
	}

	defer conn.Close()
	log.Println("WebSocket connection established")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		log.Printf("Received message: %s", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}

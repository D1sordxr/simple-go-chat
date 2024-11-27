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

type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewWebSocketHandler() *Handler {
	upgrader := websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
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
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		log.Printf("Received message: %s", data)

		if err = conn.WriteMessage(messageType, data); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}

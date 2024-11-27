package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocket struct {
	Connection *websocket.Conn
}

func NewWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return &WebSocket{Connection: conn}, nil
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"io"
)

type Server struct {
	Connections map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{Connections: make(map[*websocket.Conn]bool)}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.Connections[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.Connections {
		go func(ws *websocket.Conn) {
			_, err := ws.Write(b)
			if err != nil {
				fmt.Println("error: ", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()
	router := gin.Default()
	router.GET("/ws", func(ctx *gin.Context) {
		websocket.Handler(server.handleWS).ServeHTTP(ctx.Writer, ctx.Request)
	})
	router.Run(":3333")
}

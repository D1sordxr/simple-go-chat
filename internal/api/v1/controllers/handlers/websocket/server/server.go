package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

type Server struct {
	Connections map[*websocket.Conn]bool
	Buf         []byte
}

func NewServer() *Server {
	return &Server{
		Connections: make(map[*websocket.Conn]bool),
		Buf:         make([]byte, 1024),
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	for {
		n, err := ws.Read(s.Buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		message := s.Buf[:n]

		s.broadcast(message)
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

func (s *Server) HandleWebSocket(ws *websocket.Conn) {
	log.Printf("new incoming connection from client: %v", ws.RemoteAddr())

	s.Connections[ws] = true
	s.readLoop(ws)
}

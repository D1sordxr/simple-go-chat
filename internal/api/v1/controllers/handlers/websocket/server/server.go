package server

import (
	"encoding/json"
	"fmt"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
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
	defer func() {
		delete(s.Connections, ws)
		if err := ws.Close(); err != nil {
			log.Printf("closing connection error: %s", err.Error())
		}
	}()

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

func (s *Server) Broadcast(message dto.Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for ws := range s.Connections {
		go func(ws *websocket.Conn) {
			_, err = ws.Write(b)
			if err != nil {
				log.Printf("broadcast error: %v", err)
				delete(s.Connections, ws)
				_ = ws.Close()
			}
		}(ws)
	}

	return nil
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

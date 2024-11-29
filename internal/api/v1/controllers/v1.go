package controllers

import (
	messageHandler "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/message"
	userHandler "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/user"
	setupWSServer "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/websocket/server"
	messageRoutes "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/routes/message"
	userRoutes "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/routes/user"
	setupWSRoutes "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/routes/websocket/server"
	"github.com/D1sordxr/simple-go-chat/internal/application"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	RouterGroup *gin.RouterGroup
	UseCases    *application.UseCases
}

func NewRoutesV1(rg *gin.RouterGroup, uc *application.UseCases) {
	routes := &RoutesV1{
		RouterGroup: rg,
		UseCases:    uc,
	}
	routes.setupRoutesV1()
}

func (r *RoutesV1) setupRoutesV1() {
	// Main path
	v1 := r.RouterGroup.Group("/v1")

	// Users path
	userHandlers := userHandler.NewUserHandler(r.UseCases.UserUseCase)
	userRoutes.NewUserRoutes(v1, userHandlers)

	// Messages path
	messageHandlers := messageHandler.NewMessageHandler(r.UseCases.MessageUseCase)
	messageRoutes.NewMessageRoutes(v1, messageHandlers)

	// Websocket setup
	wsServer := setupWSServer.NewServer()
	setupWSRoutes.NewWebSocket(v1, wsServer)
}

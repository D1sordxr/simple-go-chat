package controllers

import (
	userHandler "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/user"
	userRoutes "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/routes/user"
	"github.com/D1sordxr/simple-go-chat/internal/application/user"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	RouterGroup *gin.RouterGroup
	UseCase     *user.UseCase
}

func NewRoutesV1(rg *gin.RouterGroup, uc *user.UseCase) {
	routes := &RoutesV1{
		RouterGroup: rg,
		UseCase:     uc,
	}
	routes.setupRoutesV1()
}

func (r *RoutesV1) setupRoutesV1() {
	// Main path
	v1 := r.RouterGroup.Group("/v1")

	// Users path
	userHandlers := userHandler.NewUserHandler(r.UseCase)
	userRoutes.NewUserRoutes(v1, userHandlers)
}

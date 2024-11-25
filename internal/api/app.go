package api

import (
	routesV1 "github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers"
	"github.com/D1sordxr/simple-go-chat/internal/config/config"
	"github.com/D1sordxr/simple-go-chat/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Config  *config.Config
	Storage *storage.Storage
	Router  *gin.Engine
}

func NewApp(config *config.Config,
	storage *storage.Storage,
	router *gin.Engine) *App {
	return &App{
		Config:  config,
		Storage: storage,
		Router:  router,
	}
}

func (a *App) Run() {
	port := ":" + a.Config.API.Port
	if err := a.Router.Run(port); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func (a *App) registerRoutes() {
	// Main path
	api := a.Router.Group("/api")

	// V1 path
	routesV1.NewRoutesV1(api, a.UseCase)
}

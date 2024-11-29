package main

import (
	loadApp "github.com/D1sordxr/simple-go-chat/internal/api"
	loadEngine "github.com/D1sordxr/simple-go-chat/internal/api/engine"
	loadUseCases "github.com/D1sordxr/simple-go-chat/internal/application"
	loadMessageUseCase "github.com/D1sordxr/simple-go-chat/internal/application/message"
	loadUserUseCase "github.com/D1sordxr/simple-go-chat/internal/application/user"
	loadConfig "github.com/D1sordxr/simple-go-chat/internal/config"
	loadStorage "github.com/D1sordxr/simple-go-chat/internal/storage"
	loadMessageDAO "github.com/D1sordxr/simple-go-chat/internal/storage/dao/message"
	loadUserDAO "github.com/D1sordxr/simple-go-chat/internal/storage/dao/user"
	loadPostgres "github.com/D1sordxr/simple-go-chat/internal/storage/postgres"
	"log"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	database, err := loadPostgres.NewDatabase(&cfg.Storage)
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}

	userDAO := loadUserDAO.NewUserDAO(database.Connection)
	messageDAO := loadMessageDAO.NewMessageDAO(database.Connection)
	storage, err := loadStorage.NewStorage(userDAO, messageDAO)
	if err != nil {
		log.Fatalf("error initializing storage: %v", err)
	}

	userUseCase := loadUserUseCase.NewUserUseCase(userDAO)
	messageUseCase := loadMessageUseCase.NewMessageUseCase(messageDAO)
	useCases := loadUseCases.NewUseCases(userUseCase, messageUseCase)

	router := loadEngine.NewEngine().Engine

	app := loadApp.NewApp(cfg, storage, useCases, router)
	app.Run()
}

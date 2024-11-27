package main

import (
	loadApp "github.com/D1sordxr/simple-go-chat/internal/api"
	loadEngine "github.com/D1sordxr/simple-go-chat/internal/api/engine"
	loadUseCases "github.com/D1sordxr/simple-go-chat/internal/application"
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
		log.Fatalf("error loading config: %v", err)
	}

	userDAO := loadUserDAO.NewUserDAO(database.Connection)
	messageDAO := loadMessageDAO.NewMessageDAO(database.Connection)

	storage, err := loadStorage.NewStorage(userDAO, messageDAO)
	if err != nil {
		log.Fatalf("error connecting storage: %v", err)
	}

	useCases := loadUseCases.NewUseCases(storage.UserDAO, storage.MessageDAO)

	router := loadEngine.NewEngine().Engine

	app := loadApp.NewApp(cfg, storage, useCases, router)
	app.Run()
}

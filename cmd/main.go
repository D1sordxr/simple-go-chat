package main

import (
	loadApp "github.com/D1sordxr/simple-go-chat/internal/api"
	loadEngine "github.com/D1sordxr/simple-go-chat/internal/api/engine"
	loadConfig "github.com/D1sordxr/simple-go-chat/internal/config"
	loadStorage "github.com/D1sordxr/simple-go-chat/internal/storage"
	"log"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	storage, err := loadStorage.NewStorage(&cfg.Storage)
	if err != nil {
		log.Fatalf("error connecting storage: %v", err)
	}

	router := loadEngine.NewEngine().Engine

	app := loadApp.NewApp(cfg, storage, router)
	app.Run()
}

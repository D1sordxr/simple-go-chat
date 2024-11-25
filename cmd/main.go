package main

import (
	"fmt"
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
	fmt.Println(storage, cfg)

	// TODO: init api

	// TODO: run server
}

package main

import (
	"fmt"
	loadConfig "github.com/D1sordxr/simple-go-chat/internal/config"
	"log"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	fmt.Println(cfg)

	// TODO: init storage

	// TODO: init api

	// TODO: run server
}

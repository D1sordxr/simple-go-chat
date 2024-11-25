package config

import (
	"github.com/D1sordxr/simple-go-chat/internal/config/config"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

const BasicConfigPath = "./configs/app/local.yaml"

func NewConfig() (*config.Config, error) {
	var cfg config.Config

	if err := cleanenv.ReadConfig(BasicConfigPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err.Error())
	}

	return &cfg, nil
}

package config

import (
	api "github.com/D1sordxr/simple-go-chat/internal/api/config"
	storage "github.com/D1sordxr/simple-go-chat/internal/storage/config"
)

type Config struct {
	App     AppConfig             `yaml:"app"`
	Storage storage.StorageConfig `yaml:"db"`
	API     api.APIConfig         `yaml:"api"`
}

type AppConfig struct {
	Mode string `yaml:"mode" env-default:"local"`
}

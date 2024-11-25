package config

import "fmt"

type StorageConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Database  string `yaml:"database"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Migration bool   `yaml:"migration"`
}

func (sc *StorageConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		sc.Host, sc.User, sc.Password, sc.Database, sc.Port,
	)
}

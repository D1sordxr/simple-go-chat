package storage

import (
	"context"
	storage "github.com/D1sordxr/simple-go-chat/internal/storage/config"
	"github.com/D1sordxr/simple-go-chat/internal/storage/migrations"
	"github.com/jackc/pgx/v5"
	"log"
)

type Storage struct {
	Connection *pgx.Conn
}

func NewStorage(config *storage.StorageConfig) (*Storage, error) {
	connectionString := config.ConnectionString()

	storageConnection, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("error connecting database %v", err)
	}

	if config.Migration {
		migrations.Migrate(storageConnection)
	}

	return &Storage{Connection: storageConnection}, nil
}

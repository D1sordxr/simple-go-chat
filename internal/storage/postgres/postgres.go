package postgres

import (
	"context"
	"github.com/D1sordxr/simple-go-chat/internal/storage/config"
	"github.com/D1sordxr/simple-go-chat/internal/storage/migrations"
	"github.com/jackc/pgx/v5"
	"log"
)

type Database struct {
	Connection *pgx.Conn
}

func NewDatabase(config *config.StorageConfig) (*Database, error) {
	connectionString := config.ConnectionString()

	storageConnection, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("error connecting database %v", err)
	}

	if config.Migration {
		migrations.Migrate(storageConnection)
	}

	return &Database{
		Connection: storageConnection,
	}, nil
}

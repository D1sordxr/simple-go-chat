package storage

import (
	"context"
	storage "github.com/D1sordxr/simple-go-chat/internal/storage/config"
	"github.com/D1sordxr/simple-go-chat/internal/storage/dao/message"
	"github.com/D1sordxr/simple-go-chat/internal/storage/dao/user"
	"github.com/D1sordxr/simple-go-chat/internal/storage/migrations"
	"github.com/jackc/pgx/v5"
	"log"
)

type Storage struct {
	Connection *pgx.Conn
	UserDAO    *user.DAOImpl
	MessageDAO *message.DAOImpl
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

	userDAO := user.NewUserDAO(storageConnection)
	messageDAO := message.NewMessageDAO(storageConnection)

	return &Storage{Connection: storageConnection, UserDAO: userDAO, MessageDAO: messageDAO}, nil
}

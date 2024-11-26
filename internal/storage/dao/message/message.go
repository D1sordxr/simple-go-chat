package message

import (
	"context"
	"errors"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/interfaces/dao"
	"github.com/jackc/pgx/v5"
)

type DAOImpl struct {
	Storage *pgx.Conn
	DAO     dao.MessageDAO
}

func NewMessageDAO(conn *pgx.Conn) *DAOImpl {
	return &DAOImpl{Storage: conn}
}

func (dao *DAOImpl) Create(message dto.Message) (dto.Message, error) {
	var err error
	if len(message.Content) == 0 {
		err = errors.New("message content can't be empty")
		return dto.Message{}, err
	}

	err = dao.Storage.QueryRow(context.Background(), `
		INSERT INTO messages (created_at, updated_at, content, user_id) 
		VALUES (NOW(), NOW(), $1, $2) 
		RETURNING created_at, updated_at, content, user_id, id
	`, message.Content, message.UserID).Scan(
		&message.CreatedAt,
		&message.UpdatedAt,
		&message.Content,
		&message.UserID,
		&message.ID,
	)
	if err != nil {
		return dto.Message{}, err
	}

	return message, nil
}

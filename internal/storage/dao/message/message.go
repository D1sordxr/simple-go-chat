package message

import (
	"context"
	"errors"
	"fmt"
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

func (dao *DAOImpl) Create(message dto.Message, ctx context.Context) (dto.Message, error) {
	var err error
	if len(message.Content) == 0 {
		err = errors.New("message content can't be empty")
		return dto.Message{}, err
	}

	err = dao.Storage.QueryRow(ctx, `
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

func (dao *DAOImpl) GetAll(ctx context.Context) (dto.Messages, error) {
	var messages dto.Messages

	rows, err := dao.Storage.Query(ctx, `
		SELECT id, user_id, content, is_edited, created_at, updated_at FROM messages
	`)
	if err != nil {
		return dto.Messages{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var message dto.Message
		err = rows.Scan(
			&message.ID,
			&message.UserID,
			&message.Content,
			&message.IsEdited,
			&message.CreatedAt,
			&message.UpdatedAt,
		)
		if err != nil {
			return dto.Messages{}, err
		}
		messages.Messages = append(messages.Messages, message)
	}
	if err = rows.Err(); err != nil {
		return dto.Messages{}, err
	}

	return messages, nil
}

func (dao *DAOImpl) Delete(id string, ctx context.Context) (dto.Message, error) {
	var message dto.Message

	err := dao.Storage.QueryRow(ctx, `
		DELETE FROM messages WHERE id = $1
		RETURNING created_at, updated_at, content, user_id, id
	`, id).Scan(
		&message.CreatedAt,
		&message.UpdatedAt,
		&message.Content,
		&message.UserID,
		&message.ID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.Message{}, fmt.Errorf("message with id %s not found", id)
		}
		return dto.Message{}, err
	}

	return message, nil
}

func (dao *DAOImpl) Update(message dto.Message, ctx context.Context) (dto.Message, error) {
	if len(message.Content) == 0 {
		return dto.Message{}, errors.New("message content cannot be empty")
	}

	var updatedMessage dto.Message

	err := dao.Storage.QueryRow(ctx, `
		UPDATE messages 
		SET content = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING created_at, updated_at, content, user_id, id
	`, message.Content, message.ID).Scan(
		&updatedMessage.CreatedAt,
		&updatedMessage.UpdatedAt,
		&updatedMessage.Content,
		&updatedMessage.UserID,
		&updatedMessage.ID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.Message{}, fmt.Errorf("message with id %v not found", message.ID)
		}
		return dto.Message{}, err
	}

	return updatedMessage, nil
}

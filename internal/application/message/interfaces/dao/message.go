package dao

import (
	"context"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
)

type MessageDAO interface {
	Create(message dto.Message, ctx context.Context) (dto.Message, error)
	GetAll(ctx context.Context) (dto.Messages, error)
	Delete(id string, ctx context.Context) (dto.Message, error)
}

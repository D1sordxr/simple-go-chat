package message

import (
	"context"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
)

type UseCase interface {
	Create(message dto.Message) (dto.Message, error)
	GetAll() (dto.Messages, error)
	Delete(id string, ctx context.Context) (dto.Message, error)
}

type Broadcaster interface {
	Broadcast(message dto.Message) error
}

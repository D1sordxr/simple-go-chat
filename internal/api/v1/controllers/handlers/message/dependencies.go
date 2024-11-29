package message

import "github.com/D1sordxr/simple-go-chat/internal/application/message/dto"

type UseCase interface {
	Create(message dto.Message) (dto.Message, error)
	GetAll() (dto.Messages, error)
}

type Broadcaster interface {
	Broadcast(message dto.Message) error
}

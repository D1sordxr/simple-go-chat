package dao

import "github.com/D1sordxr/simple-go-chat/internal/application/message/dto"

type MessageDAO interface {
	Create(message dto.Message) (dto.Message, error)
}

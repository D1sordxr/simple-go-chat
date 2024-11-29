package interfaces

import "github.com/D1sordxr/simple-go-chat/internal/application/message/dto"

type Broadcaster interface {
	Broadcast(message dto.Message) error
}

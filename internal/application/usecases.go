package application

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/user"
	"github.com/D1sordxr/simple-go-chat/internal/application/message"
	mDAO "github.com/D1sordxr/simple-go-chat/internal/application/message/interfaces/dao"
	loadUserUC "github.com/D1sordxr/simple-go-chat/internal/application/user"
	"github.com/D1sordxr/simple-go-chat/internal/application/user/interfaces/dao"
)

type UseCases struct {
	UserUseCase    user.UseCase
	MessageUseCase message.UseCase
}

func NewUseCases(userDAO dao.UserDAO, messageDAO mDAO.MessageDAO) *UseCases {
	userUseCase := loadUserUC.NewUserUseCase(userDAO)
	messageUseCase := message.NewMessageUseCase(messageDAO)
	return &UseCases{
		UserUseCase:    userUseCase,
		MessageUseCase: messageUseCase,
	}
}

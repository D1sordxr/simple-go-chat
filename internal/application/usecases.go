package application

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/message"
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/user"
	mDAO "github.com/D1sordxr/simple-go-chat/internal/application/message/interfaces/dao"
	uDAO "github.com/D1sordxr/simple-go-chat/internal/application/user/interfaces/dao"
)

type UseCases struct {
	UserUseCase    user.UseCase
	MessageUseCase message.UseCase
}

func NewUseCases(userDAO uDAO.UserDAO, messageDAO mDAO.MessageDAO) *UseCases {
	return &UseCases{
		UserUseCase:    userDAO,
		MessageUseCase: messageDAO,
	}
}

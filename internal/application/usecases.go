package application

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/user"
	loadUserUC "github.com/D1sordxr/simple-go-chat/internal/application/user"
	"github.com/D1sordxr/simple-go-chat/internal/application/user/interfaces/dao"
)

type UseCases struct {
	UserUseCase user.UseCase
}

func NewUseCases(userDAO dao.UserDAO) *UseCases {
	userUseCase := loadUserUC.NewUserUseCase(userDAO)
	return &UseCases{
		UserUseCase: userUseCase,
	}
}

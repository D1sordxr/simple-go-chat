package user

import (
	"github.com/D1sordxr/simple-go-chat/internal/application/user/dto"
	"github.com/D1sordxr/simple-go-chat/internal/application/user/interfaces/dao"
)

type UseCase struct {
	dao.UserDAO
}

func NewUserUseCase(dao dao.UserDAO) *UseCase {
	return &UseCase{dao}
}

func (uc *UseCase) Create(user dto.User) (dto.User, error) {
	return uc.UserDAO.Create(user)
}

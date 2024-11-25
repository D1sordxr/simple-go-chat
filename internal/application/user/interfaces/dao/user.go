package dao

import (
	"github.com/D1sordxr/simple-go-chat/internal/application/user/dto"
)

type UserDAO interface {
	Create(user dto.User) (dto.User, error)
}

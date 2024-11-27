package mock

import (
	"github.com/D1sordxr/simple-go-chat/internal/application/user/dto"
	"log"
)

type UserMock struct {
	Users []dto.User
}

func NewUserMock() *UserMock {
	return &UserMock{Users: make([]dto.User, 0, 2)}
}

func (um *UserMock) Create(user dto.User) (dto.User, error) {
	um.Users = append(um.Users, user)
	log.Printf("%v", user)
	return user, nil
}

func (um *UserMock) GetAll() ([]dto.User, error) {
	users := um.Users
	log.Printf("%v", users)
	return users, nil
}

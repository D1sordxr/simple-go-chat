package storage

import (
	mDAO "github.com/D1sordxr/simple-go-chat/internal/application/message/interfaces/dao"
	uDAO "github.com/D1sordxr/simple-go-chat/internal/application/user/interfaces/dao"
)

type Storage struct {
	uDAO.UserDAO
	mDAO.MessageDAO
}

func NewStorage(userDAO uDAO.UserDAO, msgDAO mDAO.MessageDAO) (*Storage, error) {
	return &Storage{
		UserDAO:    userDAO,
		MessageDAO: msgDAO,
	}, nil
}

package message

import (
	"context"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/interfaces/dao"
)

type UseCase struct {
	dao.MessageDAO
}

func NewMessageUseCase(dao dao.MessageDAO) *UseCase {
	return &UseCase{dao}
}

func (uc *UseCase) Create(message dto.Message) (dto.Message, error) {
	return uc.MessageDAO.Create(message)
}

func (uc *UseCase) GetAll() (dto.Messages, error) {
	return uc.MessageDAO.GetAll()
}

func (uc *UseCase) Delete(id string, ctx context.Context) (dto.Message, error) {
	return uc.MessageDAO.Delete(id, ctx)
}

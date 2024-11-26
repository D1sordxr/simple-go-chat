package message

import (
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/interfaces/dao"
	"github.com/jackc/pgx/v5"
)

type DAOImpl struct {
	Storage *pgx.Conn
	DAO     dao.MessageDAO
}

func NewMessageDAO(conn *pgx.Conn) *DAOImpl {
	return &DAOImpl{Storage: conn}
}

func (dao *DAOImpl) Create(message dto.Message) (dto.Message, error) {
	return message, nil
}

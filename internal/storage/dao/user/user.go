package user

import (
	"github.com/D1sordxr/simple-go-chat/internal/application/user/dto"
	"github.com/D1sordxr/simple-go-chat/internal/application/user/interfaces/dao"
	"github.com/jackc/pgx/v5"
)

type DAOImpl struct {
	Storage *pgx.Conn
	DAO     dao.UserDAO
}

func NewUserDAO(conn *pgx.Conn) *DAOImpl {
	return &DAOImpl{Storage: conn}
}
func (dao *DAOImpl) Create(user dto.User) (dto.User, error) {
	return dto.User{}, nil
}

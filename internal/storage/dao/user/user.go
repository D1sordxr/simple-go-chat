package user

import (
	"context"
	"errors"
	"github.com/D1sordxr/simple-go-chat/internal/application/user/dto"
	"github.com/D1sordxr/simple-go-chat/internal/application/user/interfaces/dao"
	"github.com/google/uuid"
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
	newUserID := uuid.New()
	if len(user.Username) <= 1 {
		err := errors.New("username has to be 2 or more characters")
		return dto.User{}, err
	}

	ctx := context.Background()
	query := `
		INSERT INTO users(user_id, username, created_at, updated_at) 
		VALUES ($1, $2, NOW(), NOW()) 
		RETURNING id, user_id, username, created_at, updated_at
	`

	err := dao.Storage.QueryRow(ctx, query, newUserID, user.Username).Scan(
		&user.ID,
		&user.UserID,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return dto.User{}, err
	}

	return user, nil
}

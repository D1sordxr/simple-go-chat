package user

import (
	"github.com/D1sordxr/simple-go-chat/internal/application/user/dto"
	"github.com/D1sordxr/simple-go-chat/internal/storage/dao/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestCreateAccrual(t *testing.T) {
	userMock := mock.NewUserMock()

	user := dto.User{
		ID:        1,
		UserID:    uuid.UUID{},
		Username:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	expected := []dto.User{user, user}

	useCase := NewUserUseCase(userMock)

	data, err := useCase.FirstTest(user)
	if err != nil {
		log.Printf("suka")
	}

	assert.Equal(t, expected, data)
}

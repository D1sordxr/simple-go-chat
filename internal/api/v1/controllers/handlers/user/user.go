package user

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/responses"
	"github.com/D1sordxr/simple-go-chat/internal/application/user/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	UseCase
}

type UseCase interface {
	Create(user dto.User) (dto.User, error)
}

func NewUserHandler(uc UseCase) *Handler {
	return &Handler{uc}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var newUser dto.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.CommonResponse{Message: "error", Data: err.Error()})
		return
	}

	newUser, err = h.UseCase.Create(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.CommonResponse{Message: "error", Data: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.CommonResponse{
		Message: "Successfully created!",
		Data:    newUser,
	})
}

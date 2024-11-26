package message

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/responses"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	UseCase
}

type UseCase interface {
	Create(message dto.Message) (dto.Message, error)
}

func NewMessageHandler(uc UseCase) *Handler {
	return &Handler{uc}
}

func (h *Handler) WriteMessage(c *gin.Context) {
	var message dto.Message
	err := c.BindJSON(&message)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.CommonResponse{
			Message: "Error",
			Data:    err.Error(),
		})
		return
	}

	message, err = h.UseCase.Create(message)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.CommonResponse{
			Message: "Error",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.CommonResponse{
		Message: "Successfully created!",
		Data:    message,
	})
}

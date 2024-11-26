package message

import (
	"errors"
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
	c.JSON(http.StatusNotImplemented, responses.CommonResponse{
		Message: "error",
		Data:    errors.New("not implemented"),
	})
}

package message

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/responses"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Server Broadcaster
	UseCase
}

func NewMessageHandler(uc UseCase, server Broadcaster) *Handler {
	return &Handler{
		Server:  server,
		UseCase: uc,
	}
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

	err = h.Server.Broadcast(message)
	if err != nil {
		c.JSON(http.StatusCreated, responses.CommonResponse{
			Message: "Successfully created but broadcast failed. Reload might be required.",
			Data:    message,
		})
		return
	}

	c.JSON(http.StatusCreated, responses.CommonResponse{
		Message: "Successfully created!",
		Data:    message,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	messages, err := h.UseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusConflict, responses.CommonResponse{
			Message: "Error",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, responses.CommonResponse{
		Data: messages,
	})
}

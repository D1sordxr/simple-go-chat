package message

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/responses"
	"github.com/D1sordxr/simple-go-chat/internal/application/message/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	Server  Broadcaster
	UseCase UseCase
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
		responses.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	message, err = h.UseCase.Create(message, ctx)
	if err != nil {
		responses.RespondWithError(c, http.StatusBadRequest, err.Error())
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
	ctx := c.Request.Context()
	messages, err := h.UseCase.GetAll(ctx)
	if err != nil {
		responses.RespondWithError(c, http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, responses.CommonResponse{
		Message: "Fetched all messages",
		Data:    messages,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	deletedMessage, err := h.UseCase.Delete(id, ctx)
	if err != nil {
		responses.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, responses.CommonResponse{
		Message: "Successfully deleted!",
		Data:    deletedMessage,
	})
}

func (h *Handler) Update(c *gin.Context) {
	var message dto.Message
	id := c.Param("id")

	err := c.BindJSON(&message)
	if err != nil {
		responses.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		responses.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	message.ID = intID

	ctx := c.Request.Context()
	updatedMessage, err := h.UseCase.Update(message, ctx)
	if err != nil {
		responses.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, responses.CommonResponse{
		Message: "Successfully updated!",
		Data:    updatedMessage,
	})
}

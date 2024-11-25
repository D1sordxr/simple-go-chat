package user

import (
	"github.com/D1sordxr/simple-go-chat/internal/api/v1/controllers/handlers/user"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *user.Handler
}

func NewUserRoutes(rg *gin.RouterGroup, h *user.Handler) {
	api := Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	api.setupUserRoutes()
}

func (r *Routes) setupUserRoutes() {
	api := r.RouterGroup.Group("/users")
	{
		api.POST("/user", r.Handler.CreateUser)
	}
}

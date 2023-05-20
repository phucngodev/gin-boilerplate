package handler

import (
	"apiserver/internal/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) *UserHandler {
	return &UserHandler{
		userSrv: userSrv,
	}
}

func (h *UserHandler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid int64 = 1
		user, err := h.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

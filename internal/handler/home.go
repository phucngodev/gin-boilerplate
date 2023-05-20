package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HomeHandler struct {
	logger *zap.Logger
}

func NewHomeHandler(logger *zap.Logger) *HomeHandler {
	return &HomeHandler{
		logger: logger,
	}
}

func (h *HomeHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
}

package middleware

import (
	"apiserver/internal/handler"

	"github.com/gin-gonic/gin"
)

type middleware struct {
}

func NewMidddleware() *middleware {
	return &middleware{}
}

func (m *middleware) RegisterRouter(g *gin.Engine) {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.GET("/ping", handler.Ping())
}

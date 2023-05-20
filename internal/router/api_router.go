package router

import (
	"apiserver/internal/handler"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	homeHandler *handler.HomeHandler
}

func NewApiRouter(homeHandler *handler.HomeHandler) *ApiRouter {
	return &ApiRouter{
		homeHandler: homeHandler,
	}
}

func (ar *ApiRouter) RegisterRouter(g *gin.Engine) {
	g.GET("/", ar.homeHandler.Home)
}

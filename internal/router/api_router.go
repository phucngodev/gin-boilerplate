package router

import (
	"apiserver/internal/handler"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	homeHandler *handler.HomeHandler
	userHandler *handler.UserHandler
}

func NewApiRouter(
	homeHandler *handler.HomeHandler,
	userHandler *handler.UserHandler,
) *ApiRouter {
	return &ApiRouter{
		homeHandler: homeHandler,
		userHandler: userHandler,
	}
}

func (ar *ApiRouter) RegisterRouter(g *gin.Engine) {
	g.GET("/", ar.homeHandler.Home)
	g.GET("/users", ar.userHandler.GetUserInfo())
}

package api

import (
	"try-bank/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	env    config.Environment
	router *gin.Engine
}

func NewServer(env config.Environment) *Server {
	server := &Server{
		env:    env,
		router: gin.Default(),
	}
	server.router.Delims("{[{", "}]}")
	server.router.LoadHTMLGlob("public/views")
	return server
}

func (s Server) RouteGroup(group string) *gin.RouterGroup {
	return s.router.Group(group)
}

func (s Server) Run(address ...string) {
	s.router.Run(address...)
}

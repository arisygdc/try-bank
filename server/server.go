package server

import (
	"try-bank/config"
	"try-bank/controller"

	"github.com/gin-gonic/gin"
)

type Server struct {
	env        config.Environment
	engine     *gin.Engine
	controller controller.BaseController
}

func NewServer(env config.Environment, ctr controller.BaseController) *Server {
	gin.SetMode(env.Environment)
	engine := gin.Default()
	server := &Server{
		env:        env,
		engine:     engine,
		controller: ctr,
	}

	return server
}

func (s Server) Run() {
	s.engine.Run(s.env.ServerAddress)
}

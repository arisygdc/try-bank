package server

import (
	"try-bank/config"
	"try-bank/controller"

	"github.com/gin-gonic/gin"
)

type Server struct {
	env        config.Environment
	engine     *gin.Engine
	controller controller.Controller
}

func NewServer(env config.Environment) *Server {
	engine := gin.Default()
	gin.SetMode(env.Environment)
	server := &Server{
		env:    env,
		engine: engine,
	}

	return server
}

func (s *Server) WebRouteCustomConfig() {
	s.engine.Delims("{{", "}}")
	s.engine.LoadHTMLGlob("public/views")
}

func (s Server) Run() {
	s.engine.Run(s.env.ServerAddress)
}

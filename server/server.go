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

func NewServer(env config.Environment, ctr controller.Controller) *Server {
	engine := gin.Default()
	gin.SetMode(env.Environment)
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

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

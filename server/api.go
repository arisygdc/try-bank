package server

func (s *Server) ApiRoute() {
	r := s.engine.Group("api/v1")
	r.POST("/user", s.controller.CreateUser)
}

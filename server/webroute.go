package server

func (s *Server) WebRoute() {
	s.engine.Static("/assets", "./templates/assets")
	s.engine.LoadHTMLGlob("templates/html/*")
	router := s.engine.Group("")
	router.GET("/login", s.controller.ViewLogin)
	router.POST("/login", s.controller.Login)
	router.GET("/payment/va", s.controller.ViewVA)
	router.POST("/payment/va", s.controller.WebVA)
	router.GET("/logout", s.controller.WebLogout)
}

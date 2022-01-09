package server

func (s *Server) ApiRoute() {
	r := s.engine.Group("api/v1")
	r.POST("/user", s.controller.CreateUser)
	r.POST("/level", s.controller.CreateLevel)
	r.POST("/company", s.controller.CreateCompany)
	r.PUT("/company/va", s.controller.ActivateVA)
	r.POST("/payment/va", s.controller.PaymentVA)
	r.POST("/transfer", s.controller.Transfer)
	r.POST("/login", s.controller.Login)
}

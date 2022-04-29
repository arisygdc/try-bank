package server

// TODO

func (s *Server) ApiRoute() {
	r := s.engine.Group("api/v1")
	// customer route
	r.POST("/customer", s.controller.Account.Register)
	r.POST("/transfer", s.DController.Transfer)

	// company route
	r.POST("/company", s.controller.Company.RegisterCompany)

	// virtual account route
	r.POST("/virtual-account", s.DController.ActivateVA)
	r.POST("/virtual-account/payment", s.DController.PaymentVA)
}

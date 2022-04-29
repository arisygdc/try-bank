package server

const apiV1Path = "api/v1"

func (s *Server) ApiV1Route() {
	r := s.engine.Group(apiV1Path)

	// customer route
	r.POST("/customer", s.controller.Account.Register)
	r.POST("/transfer", s.DController.Transfer)

	// company route
	r.POST("/company", s.controller.Company.RegisterCompany)

	// virtual account route
	// TODO
	// refund, get payment status
	r.POST("/virtual-account", s.controller.VirtualAccount.Register)
	r.POST("/virtual-account/payment", s.controller.VirtualAccount.VirtualAccount_pay)
}

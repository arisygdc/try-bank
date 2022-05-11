package server

import "try-bank/server/middleware"

const apiV1Path = "api/v1"

func (s *Server) ApiV1Route(m middleware.Middleware) {
	r := s.engine.Group(apiV1Path)
	authenticated := r.Use(m.AuthBearer())
	// customer route
	// TODO
	r.POST("/customer", s.controller.Account.Register)
	authenticated.POST("/transfer", s.controller.Account.Transfer)

	// company route
	r.POST("/company", s.controller.Company.RegisterCompany)

	// virtual account route
	// TODO
	// refund, get payment status
	authenticated.POST("/virtual-account", s.controller.VirtualAccount.Register)
	authenticated.POST("/virtual-account/payment", s.controller.VirtualAccount.VirtualAccount_pay)
}

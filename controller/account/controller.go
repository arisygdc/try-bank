package account

import appservice "try-bank/app_service"

type ICtrAccount interface {
}

type AccountController struct {
	service appservice.BaseService
}

func New(service appservice.BaseService) ICtrAccount {
	return AccountController{
		service: service,
	}
}

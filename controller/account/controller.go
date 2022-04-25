package account

import appservice "try-bank/app_service"

type ICtrAccount interface {
}

type AccountController struct {
	service appservice.ISBase
}

func New(service appservice.ISBase) ICtrAccount {
	return AccountController{
		service: service,
	}
}

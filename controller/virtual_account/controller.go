package virtualaccount

import (
	appservice "try-bank/app_service"
)

type ICtrVirtualAccount interface {
}

type VirtualAccountController struct {
	service appservice.ISBase
}

func New(service appservice.ISBase) ICtrVirtualAccount {
	return VirtualAccountController{
		service: service,
	}
}

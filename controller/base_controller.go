package controller

import (
	appservice "try-bank/app_service"
	"try-bank/controller/account"
	"try-bank/controller/company"
	virtualaccount "try-bank/controller/virtual_account"
)

type BaseController struct {
	Account        account.ICtrAccount
	Company        company.ICtrCompany
	VirtualAccount virtualaccount.ICtrVirtualAccount
}

func NewController(service appservice.BaseService) BaseController {
	return BaseController{
		Account:        account.New(service),
		Company:        company.New(service),
		VirtualAccount: virtualaccount.New(service),
	}
}

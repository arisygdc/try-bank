package controller

import (
	appservice "try-bank/app_service"
	"try-bank/controller/account"
	"try-bank/controller/company"
	virtualaccount "try-bank/controller/virtual_account"
)

type Controller struct {
	Account        account.AccountController
	Company        company.CompanyController
	VirtualAccount virtualaccount.VirtualAccountController
}

func NewController(service appservice.BaseService) {

}

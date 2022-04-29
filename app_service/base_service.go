package appservice

import (
	"try-bank/app_service/account"
	"try-bank/app_service/company"
	virtualaccount "try-bank/app_service/virtual_account"
	"try-bank/database"
)

type BaseService struct {
	Account        account.ISAccount
	Company        company.ISCompany
	VirtualAccount virtualaccount.ISVirtualAccount
}

func NewService(repository database.IRepository) BaseService {
	r := BaseService{
		Account:        account.New(repository),
		Company:        company.New(repository),
		VirtualAccount: virtualaccount.New(repository),
	}
	return r
}

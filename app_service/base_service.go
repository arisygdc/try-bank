package appservice

import (
	"try-bank/app_service/account"
	"try-bank/app_service/company"
	virtualaccount "try-bank/app_service/virtual_account"
	"try-bank/database"
)

type ISBase interface {
}

type BaseService struct {
	Account        account.Service
	Company        company.Service
	VirtualAccount virtualaccount.Service
}

func NewService(repository database.IRepository) ISBase {
	return BaseService{
		Account:        account.New(repository),
		Company:        company.New(repository),
		VirtualAccount: virtualaccount.New(repository),
	}
}

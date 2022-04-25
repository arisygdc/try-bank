package company

import appservice "try-bank/app_service"

type ICtrCompany interface {
}

type CompanyController struct {
	service appservice.ISBase
}

func New(service appservice.ISBase) ICtrCompany {
	return CompanyController{
		service: service,
	}
}

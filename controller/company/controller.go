package company

import appservice "try-bank/app_service"

type ICtrCompany interface {
}

type CompanyController struct {
	service appservice.BaseService
}

func New(service appservice.BaseService) ICtrCompany {
	return CompanyController{
		service: service,
	}
}

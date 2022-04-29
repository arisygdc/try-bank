package company

import (
	appservice "try-bank/app_service"

	"github.com/gin-gonic/gin"
)

type ICtrCompany interface {
	RegisterCompany(ctx *gin.Context)
}

type CompanyController struct {
	service appservice.BaseService
}

func New(service appservice.BaseService) ICtrCompany {
	return CompanyController{
		service: service,
	}
}

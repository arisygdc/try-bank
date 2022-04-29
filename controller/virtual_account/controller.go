package virtualaccount

import (
	appservice "try-bank/app_service"

	"github.com/gin-gonic/gin"
)

type ICtrVirtualAccount interface {
	Register(ctx *gin.Context)
	VirtualAccount_pay(ctx *gin.Context)
}

type VirtualAccountController struct {
	service appservice.BaseService
}

func New(service appservice.BaseService) ICtrVirtualAccount {
	return VirtualAccountController{
		service: service,
	}
}

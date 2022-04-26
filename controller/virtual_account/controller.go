package virtualaccount

import (
	appservice "try-bank/app_service"

	"github.com/gin-gonic/gin"
)

type ICtrVirtualAccount interface {
	Register(ctx *gin.Context)
}

type VirtualAccountController struct {
	service appservice.ISBase
}

func New(service appservice.ISBase) ICtrVirtualAccount {
	return VirtualAccountController{
		service: service,
	}
}

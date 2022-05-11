package account

import (
	appservice "try-bank/app_service"

	"github.com/gin-gonic/gin"
)

type ICtrAccount interface {
	Register(ctx *gin.Context)
}

type AccountController struct {
	service appservice.BaseService
}

func New(service appservice.BaseService) ICtrAccount {
	return AccountController{
		service: service,
	}
}

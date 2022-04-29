package virtualaccount

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterVirtualAccount struct {
	RegisteredNumber int32  `json:"registered_number" binding:"required"`
	CallbackURL      string `json:"callback_url" binding:"required"`
}

func (ctr VirtualAccountController) Register(ctx *gin.Context) {
	var req RegisterVirtualAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ca, err := ctr.service.Company.CompanyAccount(ctx, req.RegisteredNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	registered, err := ctr.service.VirtualAccount.Register(ctx, ca.CompanyID, req.CallbackURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"va_identity":         registered.Identity,
		"va_authorizationKey": registered.AuthorizationKey,
		"callback_url":        registered.Callback_url,
	})
}

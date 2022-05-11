package virtualaccount

import (
	"net/http"
	"try-bank/server/middleware"
	"try-bank/token"

	"github.com/gin-gonic/gin"
)

type RegisterVirtualAccount struct {
	CallbackURL string `json:"callback_url" binding:"required"`
}

func (ctr VirtualAccountController) Register(ctx *gin.Context) {
	var req RegisterVirtualAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	getPayload, exists := ctx.Get(middleware.PayloadKey)
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}

	payload, ok := getPayload.(*token.Payload)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}

	ca, err := ctr.service.Company.CompanyAccount(ctx, payload.Registered_number)
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
